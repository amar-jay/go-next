package gql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/amar-jay/go-api-boilerplate/controller/gql/gen"
	"github.com/amar-jay/go-api-boilerplate/service/authservice"
	"github.com/amar-jay/go-api-boilerplate/service/emailservice"
	"github.com/amar-jay/go-api-boilerplate/service/userservice"
)

type Resolver struct {
	AuthService  authservice.AuthService
	UserService  userservice.UserService
	EmailService emailservice.EmailService
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
