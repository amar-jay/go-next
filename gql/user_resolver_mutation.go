package gql

import (
	"context"
	// "errors"
	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/amar-jay/go-api-boilerplate/gql/gen"
)

// // foo
// TODO:
func (r *mutationResolver) UpdateUser(ctx context.Context, input gen.UpdateUser) (*gen.Message, error) {
  return &gen.Message{}, nil
}

// // foo
// TODO: 
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*gen.Message, error) {
  return &gen.Message{
    Text: "foo",
  }, nil
}

// register user and send welcome email
func (r *mutationResolver) Register(ctx context.Context, login gen.RegisterInput) (*gen.RegisterLoginOutput, error) {
  u := r.UserService.Register(user.ToUser(login))
  return &gen.RegisterLoginOutput{
    Token: "not implemented",
    User: &gen.User{},
  }, nil
}

// To login a user and return a token
func (r *mutationResolver) Login(ctx context.Context, login gen.LoginInput) (*gen.RegisterLoginOutput, error) {
  return &gen.RegisterLoginOutput{
    Token: "foo",
    User: &gen.User{},
  }, nil
}

// To login a user and return a token
//resetPassword(token: String!, password: String!): RegisterLoginOutput!
func (r *mutationResolver) ResetPassword(ctx context.Context, token string, password string) (*gen.RegisterLoginOutput, error) {
  return &gen.RegisterLoginOutput{
    Token: "foo",
    User: &gen.User{},
  }, nil
}
