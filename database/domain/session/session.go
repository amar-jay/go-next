package session

import (
	"time"

	"github.com/amar-jay/go-api-boilerplate/controller/gql/gen"
	"github.com/jinzhu/gorm"
)

type Session struct {
	UserID  string    `gorm:"size:255" json:"userId"`
	Token   string    `gorm:"size:255" json:"sessionToken"`
	Expires time.Time `gorm:"NOT NULL; DEFAULT: now()" json:"expires"`
	gorm.Model
}

func (u *Session) ToGql() *gen.Session {
	return &gen.Session{
		UserID:  u.UserID,
		Token:   u.Token,
		Expires: u.Expires.String(),
	}
}

func ToSession(g *gen.SessionInput) (*Session, error) {
	exp, err := time.Parse(time.RFC3339, g.Expires)
	if err != nil {
		return nil, err
	}

	return &Session{
		UserID:  g.UserID,
		Token:   g.Token,
		Expires: exp,
	}, nil
}
