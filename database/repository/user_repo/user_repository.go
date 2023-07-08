package user_repo

import (
	"errors"

	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"gorm.io/gorm"
)

// Repository interface
type Repo interface {
	GetUserByID(id string) (*user.User, error) // get useer by user id; id refers to **user_id** in db
	GetUsers() ([]*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	// GetUserByAccount(provider_type string, acc_id string) (*user.User, error)
	CreateUser(user *user.User) error
	Update(user *user.User) error
	DeleteUser(id string) error
}

type userRepo struct {
	db *gorm.DB
}

// New user repo instance
func NewUserRepo(db *gorm.DB) Repo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) GetUsers() ([]*user.User, error) {
	var users []*user.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// get first user by id
func (repo *userRepo) GetUserByID(id string) (*user.User, error) {
	var user user.User
	if err := repo.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get first user by Email
func (repo *userRepo) GetUserByEmail(email string) (*user.User, error) {
	var user user.User

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// create user in db
func (repo *userRepo) CreateUser(user *user.User) error {
	return repo.db.Create(user).Error
}

// change user info
func (repo *userRepo) Update(input_user *user.User) error {
	return repo.db.Where("email = ?", input_user.Email).Updates(input_user).Error
}

// delete user from db by user_id
func (repo *userRepo) DeleteUser(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	// soft delete user
	return repo.db.Where("user_id = ?", id).Update("active", false).Error
}

// func (repo *userRepo) GetUserByAccount(provider_type string, acc_id string) (*user.User, error) {
// 	var user user.User
// 	var acc account.Account
// 	if err := repo.db.Where("provider_type = ? AND acc_id = ?", provider_type, acc_id).First(&acc).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := repo.db.Where("user_id = ?", acc.UserID).First(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
