package gql

import (
	"context"
	"errors"

	// "errors"

	"github.com/amar-jay/go-api-boilerplate/controllers/gql/gen"
)

// Users returns all users
func (r *queryResolver) Users(ctx context.Context) ([]*gen.User, error) {
	ser, err := r.UserService.GetUsers()
	if err != nil {
		return nil, err
	}

	users := []*gen.User{}
	for _, user := range ser {
		users = append(users, &gen.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
		})
	}
	return users, nil
}

// User returns a user by id
func (r *queryResolver) User(ctx context.Context, id int) (*gen.User, error) {
	if id < 0 {
		return nil, errors.New("invalid id")
	}
	uid := uint(id)
	_user, err := r.UserService.GetUserByID(uid)

	if err != nil {
		return nil, errors.New("invalid id")
	}

	return &gen.User{
		Email:     _user.Email,
		FirstName: _user.FirstName,
		LastName:  _user.LastName,
		Role:      _user.Role,
		Active:    _user.Active,
		ID:        int(_user.ID),
	}, nil
}
