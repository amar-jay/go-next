package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	controllers "github.com/amar-jay/go-api-boilerplate/controllers/rest"
	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"github.com/amar-jay/go-api-boilerplate/controllers/gql"
	"github.com/amar-jay/go-api-boilerplate/controllers/gql/gen"
	"github.com/amar-jay/go-api-boilerplate/middleware"
	"github.com/amar-jay/go-api-boilerplate/database/repositories/password_reset"
	"github.com/amar-jay/go-api-boilerplate/database/repositories/user_repo"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/amar-jay/go-api-boilerplate/utils/config"
	"github.com/amar-jay/go-api-boilerplate/infra/redis"
	hmachash "github.com/amar-jay/go-api-boilerplate/utils/hash"
)


var (
	router = gin.Default()
)

func main() {
  fmt.Println("Starting server...")
  err := router.SetTrustedProxies([]string{"192.168.1.2", "::1"})

  if err != nil {
    log.Fatalf("Error setting trusted proxies: %v", err)
  }



  router.Use(gin.Logger())
  router.Use(gin.Recovery())
  router.Use(middleware.GinContextToMiddleWare())

/*

*/
  // load env file
  if err := godotenv.Load(); err != nil {
	  log.Fatal("Error in loading env files. $PORT must be set")
  }
  config := config.GetConfig()

  if config.IsProduction() {
    gin.SetMode(gin.ReleaseMode)
  }
  // swagger url - http://localhost:8080/swagger/index.html
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  db, err := gorm.Open(
	  config.Postgres.GetConnectionInfo(),
	  config.Postgres.Config(),
  )

  if err != nil {
	  panic(err)
  }

  // Migrate the schema
  err = db.AutoMigrate(&user.User{})
  if err != nil {
    panic(err)
  }
  //	defer db.Close()
  fmt.Println("Database migrated successfully")

  // TODO: remove this, meant for testing
  err = redis.Set("test", "test");
  if err != nil {
    panic(err)
  }
  res, err := redis.Get("test")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Redis test: %s %s", res)
  router.GET("/", func(c *gin.Context) {
	  // If the client is 192.168.1.2, use the X-Forwarded-For
	  // header to deduce the original client IP from the trust-
	  // worthy parts of that header.
	  // Otherwise, simply return the direct client IP
	  fmt.Printf("ClientIP: %s\n", c.ClientIP())
	  c.JSON(http.StatusOK, gin.H{"Amar": "Jay", "clientIP": c.ClientIP()})
  })

  // Testing the database
  router.GET("/ping", func(c *gin.Context) {
	  c.String(http.StatusOK, "pong")
  })

  /**
  *  ----- Services -----
   */
  userrepo := user_repo.NewUserRepo(db)
  pswdrepo := password_reset.CreatePasswordReserRepo(db)
//  randomstr := randomstring.CreateRandomString()
  hash := hmachash.NewHMAC(config.HashKey)
  userService := userservice.NewUserService(userrepo, pswdrepo, hash, config.Pepper)
  authService := authservice.NewAuthService(config.HashKey)
  emailService := emailservice.NewEmailService()

  /**
  * ----- Controllers -----
   */

  userController := controllers.NewUserController(userService, authService, emailService)

  /**
  *  ----- Routing -----
   */

   srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &gql.Resolver{
    UserService: userService,
    AuthService: authService,
    EmailService: emailService,
  }}))
  playground := playground.Handler("GraphQL playground", "/query")
  http.Handle("/", playground)
  http.Handle("/query", srv)
  router.GET("/graphql",  gql.PlaygroundHandler("/query"))
  router.POST("/query", gql.GraphQLHandler(userService, authService, emailService))
    //func(_ *gin.Context) {
	//  middleware.SetUserContext(config.JWTSecret)
  //})
  // http.Handle("/query", srv)

  auth := router.Group("/auth")

  auth.POST("/register", userController.Register)
  auth.POST("/update", userController.Update)
  auth.POST("/login", userController.Login)
  auth.POST("/forgot-password", userController.ForgotPassword)
  auth.POST("/update-password", userController.ResetPassword)

  user := router.Group("/users")

  user.GET("/", userController.GetUsers)
  user.GET("/:id", userController.GetUserByID)

  //  accounts and profiles
  account := router.Group("/account")
  account.Use(middleware.RequireTobeloggedIn(config.JWTSecret))
  {
	  account.GET("/profile", userController.GetProfile)
	  account.PUT("/profile", userController.Update)
  }

  // Run server
  log.Printf("Running on http://localhost:%d/ ", config.Port)
  port := fmt.Sprintf(":%d", config.Port)
 // log.Fatal(router.Run(port))
  log.Fatal(http.ListenAndServe(port, nil))
}
