package user

import (
	"github.com/amar-jay/go-api-boilerplate/controller/gql/gen"
	"github.com/jinzhu/gorm"
)

type User struct {
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	UserID    string `gorm:"size:255;unique"`

	Image    string `gorm:"size:255"`
	Email    string `gorm:"NOT NULL;UNIQUE_INDEX"`
	Password string `gorm:"NOT NULL"`
	Role     string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active   bool   `gorm:"NOT NULL; DEFAULT: true"`
	gorm.Model
	gen.RegisterInput
}

func (u *User) ToGql() *gen.User {
	return &gen.User{
		FirstName: u.FirstName,
		LastName:  &u.LastName,
		Email:     u.Email,
		Role:      u.Role,
		Active:    u.Active,
		Image:     &u.Image,
		// UserID:    u.UserID,
		ID: int(u.ID),
	}
}

func ToUser(g *gen.RegisterInput) (*User, error) {

	return &User{
		FirstName: g.FirstName,
		LastName:  g.LastName,
		Email:     g.Email,
		Password:  g.Password,
		Role:      "standard",
		Active:    true,
	}, nil
}
