package userservice

import (
	"errors"
	"regexp"
	"strings"

	"github.com/amar-jay/go-api-boilerplate/domain/user"
	pswd_repo "github.com/amar-jay/go-api-boilerplate/repositories/password_reset"
	"github.com/amar-jay/go-api-boilerplate/repositories/user_repo"
	"github.com/amar-jay/go-api-boilerplate/utils/hash"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
		ComparePassword(inputpswd string, dbpswd string) error
		validate(input *user.User) error
		Register(user *user.User) error
		Update(user *user.User) error
		GetUserByID(id uint) (*user.User, error)
		Login(input *user.User) (*user.User, error)
		GetUsers() ([]*user.User, error)
}

type userService struct {
	pepper string
	Repo user_repo.Repo
	pswd pswd_repo.Repo
	hmac hmachash.HMAC 
}

func NewUserService(repo user_repo.Repo, pswd pswd_repo.Repo, hmac hmachash.HMAC, pepper string ) UserService {

	return &userService{
		Repo: repo,
		pepper: pepper,
		pswd: pswd,
		hmac: hmac,
	}
}

func (us *userService) Register(u *user.User) error {
	hashed, err := us.hashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashed

	if err := us.validate(u); err != nil {
	  return err
	}
	return us.Repo.CreateUser(u)
	//return fmt.Errorf("USER SERVICE ERROR: Register not implemented")
}

/**
* ----- UPDATE METHODS ---
*/


func (us *userService) Update(u *user.User) error {
	return us.Repo.Update(u)
}

/**
* ----- GET METHODS ---
*/

func (us *userService) GetUsers() ([]*user.User, error) {
	users, err := us.Repo.GetUsers()

	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		return nil, errors.New("there is no user") 
	}

	return users, nil
}

func (us *userService) GetUserByID(id uint) (*user.User, error) {
	if id <= 0 {
		return nil, errors.New("id params is required")
	}
	user, err := us.Repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (us *userService) Login(input *user.User) (*user.User, error) {
	if err := validateEmail(input.Email); err != nil {
		return nil, err 
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err 
	}

	user, err := us.Repo.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
/**
* -- Other
*/

// HashPassword hashes the password using bcrypt
func (us *userService) hashPassword(password string) (string, error) {
	pswdAndPepper := password + us.pepper
	hashed, err := bcrypt.GenerateFromPassword([]byte(pswdAndPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}


// ComparePassword compares the password with the hash
func (us *userService) ComparePassword(inputpswd string, dbpswd string) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(dbpswd),
		[]byte(inputpswd+us.pepper),
	)
}

// validateEmail validates the email
func validateEmail(email string) error {
	 emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	  if !emailRegex.MatchString(email) {
			return errors.New("invalid email param entered")
		}

		return nil
}

// validate password 
func validatePassword(password string) error {
  valid := len(password) > 8 && strings.ToUpper(password) != password && strings.ToLower(password) != password;

  if !valid {
	return errors.New("invalid password entered")
  }
  return nil
}
// validate validates the user (password, email, name)
func (us *userService) validate(input *user.User) error {
  // validate email
  if err := validateEmail(input.Email); err != nil {
    return err
  }
  // validate password
  if err := validatePassword(input.Password); err != nil {
    return err
  }
  // if user already exists
  if _, err := us.Login(input); err == nil {
    return errors.New("user already exists")
  }

  return nil
}
