package authservice

import (
	"time"

	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// resolving circular import was done by repetition of Claim
// Claims struct represents the claims in a JWT
type Claim struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}

// for auth service - ( issueing and parsing tokens)
type AuthService interface {
	IssueToken(user user.User) (string, error)
	ParseToken(token string) (*Claim, error)
}

type authService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) AuthService {
	return &authService{
		jwtSecret: jwtSecret,
	}
}

// Generate Token for auth
func (auth *authService) IssueToken(u user.User) (string, error) {
	currTime := time.Now()
	expireTime := currTime.Add(24 * time.Hour) // after 24 hours

	claims := Claim{
		u.Email,
		int(u.ID),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Undefined Issuer",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(auth.jwtSecret))
}

// parse token
func (auth *authService) ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.jwtSecret), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claim)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
