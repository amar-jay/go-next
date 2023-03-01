package authservice

import (
	"testing"

	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	t.Run("generating Token", func(t *testing.T) {
		u := &user.User{
			FirstName: "Amar",
			LastName:  "Jay",
			Password:  "password",
			Email:     "me@themanan.me",
			Role:      "admin",
			Active:    true,
		}

		svc := NewAuthService("secret")

		token, err := svc.IssueToken(*u)
		assert.Nil(t, err)
		assert.IsType(t, "string", token)

	})
}
