package middleware

import (
	"context"
	"errors"

	"github.com/amar-jay/go-api-boilerplate/utils/config"
	"github.com/gin-gonic/gin"
)

// moving context to GinContextKey
func GinContextToMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), config.GinContextKey(), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// recover gin.Context from context.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := errors.New("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := errors.New("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
