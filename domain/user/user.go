package user

import (
	"errors"

	"github.com/amar-jay/go-api-boilerplate/gql/gen"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"size:255"`
	LastName string `gorm:"size:255"`
	Email string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password string `gorm:"NOT NULL"`
	Role string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active bool `gorm:"NOT NULL; DEFAULT: true"`
}

func (u *User) ToGenUser() *gen.User {
  return &gen.User{
    FirstName: u.FirstName,
    LastName: u.LastName,
    Email: u.Email,
    Role: u.Role,
    Active: u.Active,
    ID: int(u.ID),
  }
}

func ToUser(g any) (*User, error) {

  u, ok := g.(*User);
  if !ok {
    return nil, errors.New("invalid type: expected *User")
  }
  return u, nil;
}
