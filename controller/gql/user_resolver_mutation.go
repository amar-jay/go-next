package gql

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	// "errors"
	"github.com/amar-jay/go-api-boilerplate/controller/gql/gen"
	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
)

// TODO: implement update user
func (r *mutationResolver) UpdateUser(ctx context.Context, input gen.UpdateUser) (*gen.Message, error) {
	return &gen.Message{}, nil
}

// TODO: implement forgot password
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*gen.Message, error) {
	return &gen.Message{
		Text: "foo",
	}, nil
}

// register user and send welcome email
func (r *mutationResolver) Register(ctx context.Context, login gen.RegisterInput) (*gen.RegisterLoginOutput, error) {
	user, err := user.ToUser(&login)
	if err != nil {
		return nil, err
	}
	// register user in db
	if err = r.UserService.Register(user); err != nil {
		return nil, err
	}

	// send welcome email
	if err := r.EmailService.Welcome(user.Email); err != nil {
		return nil, err
	}

	// generate token
	token, err := r.AuthService.IssueToken(*user)
	if err != nil {
		return nil, err
	}

	return &gen.RegisterLoginOutput{
		Token: token,
		User:  user.ToGql(),
	}, nil
}

// To login a user and return a token
func (r *mutationResolver) Login(ctx context.Context, login gen.LoginInput) (*gen.RegisterLoginOutput, error) {
	token, ok := ctx.Value("token").(string)
	if !ok {
		return nil, errors.New("token not found")
	}
	return &gen.RegisterLoginOutput{
		Token: token,
		User:  &gen.User{},
	}, nil
}

// To login a user and return a token
// resetPassword(token: String!, password: String!): RegisterLoginOutput!
func (r *mutationResolver) ResetPassword(ctx context.Context, token string, password string) (*gen.RegisterLoginOutput, error) {
	return &gen.RegisterLoginOutput{
		Token: "foo",
		User:  &gen.User{},
	}, nil
}

func (r *mutationResolver) CreateSession(ctx context.Context, session gen.SessionInput) (*gen.Session, error) {
	return &gen.Session{
		UserID: session.UserID,
		Token:  session.Token,
		// Expires: session.Expires,
	}, nil
}
func (r *mutationResolver) DeleteSession(ctx context.Context, token string) (*gen.Message, error) {
	return &gen.Message{
		Code: fmt.Sprintln(http.StatusNoContent),
		Text: "Session deleted",
	}, nil

}

func (r *mutationResolver) GetSession(ctx context.Context, userId string) (*gen.Session, error) {
	return &gen.Session{
		UserID: userId,
		Token:  "",
		// Expires: session.Expires,
	}, nil

}

func (r *mutationResolver) UpdateSession(ctx context.Context, session gen.SessionInput) (*gen.Session, error) {
	return &gen.Session{
		UserID: session.UserID,
		Token:  session.Token,
		// Expires: session.Expires,
	}, nil
}
