package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/amar-jay/go-api-boilerplate/controller/gql"
	controllers "github.com/amar-jay/go-api-boilerplate/controller/rest"
	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"github.com/amar-jay/go-api-boilerplate/database/repository/password_reset"
	"github.com/amar-jay/go-api-boilerplate/database/repository/user_repo"
	"github.com/amar-jay/go-api-boilerplate/infra/redis"
	"github.com/amar-jay/go-api-boilerplate/middleware"
	"github.com/amar-jay/go-api-boilerplate/pkg/config"
	hmachash "github.com/amar-jay/go-api-boilerplate/pkg/hash"
	"github.com/amar-jay/go-api-boilerplate/service/authservice"
	"github.com/amar-jay/go-api-boilerplate/service/emailservice"
	"github.com/amar-jay/go-api-boilerplate/service/userservice"
)

var (
	router = gin.Default()
)

func redis_test() {

	// TODO: remove this, meant for testing
	conn := redis.Init()
	err := conn.Dial()
	if err != nil {
		panic(err)
	}

	err = conn.Set("test", "test")
	if err != nil {
		panic(err)
	}

	var res string
	err = conn.Get("test", &res)
	if err != nil {
		panic(err)
	}
	log.Printf("Redis test: %s", res)
}

func other_routers(config config.Config, userController controllers.UserController) {
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
}

func main() {
	var port string;
	log.Println("Starting server...")
	err := router.SetTrustedProxies([]string{"192.168.1.2", "::1"})

	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.GinContextToMiddleWare())


	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env files. Default $PORT must be set")
	}

	config := config.GetConfig();
	//port = strconv.Itoa(config.Port);
	if len(os.Args) > 1 {
		port = os.Args[1];
	} else {
		port = strconv.Itoa(config.Port);
	}

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
	log.Println("Database migrated successfully")

	redis_test();
	router.GET("/", func(c *gin.Context) {
		log.Printf("ClientIP: %s\n", c.ClientIP())
		c.Redirect(http.StatusMovedPermanently, "/graphql")
//		c.JSON(http.StatusOK, gin.H{"Amar": "Jay", "clientIP": c.ClientIP()})
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
	hash := hmachash.NewHMAC(config.HashKey)
	userService := userservice.NewUserService(userrepo, pswdrepo, hash, config.Pepper)
	authService := authservice.NewAuthService(config.HashKey)
	emailService := emailservice.NewEmailService()

	/**
	 ----- Controllers -----
	 */

	userController := controllers.NewUserController(userService, authService, emailService)

	/**
	  ----- Routing -----
	 */
	router.GET("/graphql", gql.PlaygroundHandler("/query"))
	router.POST("/query", gql.GraphQLHandler(userService, authService, emailService),
		func(c *gin.Context) {
			middleware.SetUserContext(config.JWTSecret)
		})

	other_routers(config, userController);

	// Run server
	log.Printf("Running on http://localhost:%s/\n", port)
	err = router.Run(":" + port)
	log.Fatal(err)
}
