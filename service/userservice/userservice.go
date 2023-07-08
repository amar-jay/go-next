package userservice

import (
	"errors"
	"regexp"
	"strings"

	"github.com/amar-jay/go-api-boilerplate/database/domain/account"
	models "github.com/amar-jay/go-api-boilerplate/database/domain/session"
	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	acc_repo "github.com/amar-jay/go-api-boilerplate/database/repository/account"
	pswd_repo "github.com/amar-jay/go-api-boilerplate/database/repository/password_reset"
	sess_repo "github.com/amar-jay/go-api-boilerplate/database/repository/session"
	"github.com/amar-jay/go-api-boilerplate/database/repository/user_repo"
	hmachash "github.com/amar-jay/go-api-boilerplate/pkg/hash"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	ComparePassword(inputpswd string, dbpswd string) error
	validate(input *user.User) error
	Register(user *user.User) error
	Update(user *user.User) error
	CreateUser(user *user.User) error
	GetUserByID(id string) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	GetUserByAccount(provider_type string, acc_id string) (*user.User, error)
	DeleteUser(id string) error
	Login(input *user.User) (*user.User, error)
	GetUsers() ([]*user.User, error)

	CreateSession(s *models.Session) error
	GetSession(token string) (*models.Session, error)
	DeleteSession(token string) error
	UpdateSession(s *models.Session) (*models.Session, error)

	LinkAccount(s *account.Account) error
	UnlinkAccount(provider_type string, id string) error
}

type userService struct {
	pepper string
	Repo   user_repo.Repo
	sess   sess_repo.Repo
	acc    acc_repo.Repo
	pswd   pswd_repo.Repo
	hmac   hmachash.HMAC
}

func NewUserService(repo user_repo.Repo, pswd pswd_repo.Repo, sess sess_repo.Repo, acc acc_repo.Repo, hmac hmachash.HMAC, pepper string) UserService {

	return &userService{
		Repo:   repo,
		pepper: pepper,
		pswd:   pswd,
		hmac:   hmac,
		acc:    acc,
		sess:   sess,
	}
}

func (us *userService) DeleteUser(id string) error {

	if id == "" {
		return errors.New("id is required")
	}

	return us.Repo.DeleteUser(id)

}

// no password here. this creates a user without a password
func (us *userService) CreateUser(u *user.User) error {
	if err := validateEmail(u.Email); err != nil {
		return err
	}

	// check if user already exists
	_, err := us.Repo.GetUserByEmail(u.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	return us.Repo.CreateUser(u)
	//return fmt.Errorf("USER SERVICE ERROR: Register not implemented")
}
func (us *userService) Register(u *user.User) error {
	if err := us.validate(u); err != nil {
		return err
	}

	// hashing password
	hashed, err := us.hashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashed

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

func (us *userService) GetUserByID(id string) (*user.User, error) {
	if id == "" {
		return nil, errors.New("id params is required")
	}
	user, err := us.Repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (us *userService) GetUserByEmail(email string) (*user.User, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	user, err := us.Repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) GetUserByAccount(provider_type string, acc_id string) (*user.User, error) {
	if provider_type == "" || acc_id == "" {
		return nil, errors.New("provider_type and acc_id params are required")
	}

	acc, err := us.acc.GetAccountByProvider(provider_type, acc_id)

	if err != nil {
		return nil, err
	}

	u, err := us.Repo.GetUserByID(acc.UserId)
	if err != nil {
		return nil, err
	}

	return u, nil
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
	invalid := len(password) < 8 || strings.ToUpper(password) == password || strings.ToLower(password) == password

	if invalid {
		return errors.New("invalid password entered, Must contain at least 8 characters, 1 uppercase, 1 lowercase")
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

func (us *userService) CreateSession(s *models.Session) error {
	return us.sess.CreateSession(s)
}

func (us *userService) GetSession(token string) (*models.Session, error) {
	return us.sess.GetSession(token)
}

func (us *userService) DeleteSession(token string) error {
	return us.sess.DeleteSessionByToken(token)
}

func (us *userService) UpdateSession(s *models.Session) (*models.Session, error) {
	return us.sess.Update(s)
}

func (us *userService) LinkAccount(s *account.Account) error {
	return us.acc.CreateAccount(s)
}

func (us *userService) UnlinkAccount(provider_type string, id string) error {
	return us.acc.DeleteAccount(provider_type, id)
}
