package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	hand "github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amar-jay/go-api-boilerplate/controllers/gql/gen"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/gin-gonic/gin"
)

// This defines all the Gqlgen graphql server handlers
func GraphQLHandler(us userservice.UserService, as authservice.AuthService, es emailservice.EmailService) gin.HandlerFunc {

  conf := gen.Config{
    Resolvers: &Resolver{
      UserService: us,
      AuthService: as,
      EmailService: es,
    },
  }

  // NewExecutableSchema and Config are in the generated.go file
  // Resolver is in the resolver.go file
  exec := gen.NewExecutableSchema(conf)
  h := handler.NewDefaultServer(exec)
  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }
}


// PlaygroundHandler defined the playground handler to expose
func PlaygroundHandler(path string) gin.HandlerFunc {
  h := playground.Handler("GraphQL", path)
  return func(ctx *gin.Context) {
    h.ServeHTTP(ctx.Writer, ctx.Request)
  }
}

func httpPlayground(path string) http.HandlerFunc {
   h := hand.Playground("GraphQL  Playground", path)
  return h;
}
