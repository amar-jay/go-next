package session_repo

import (
	models "github.com/amar-jay/go-api-boilerplate/database/domain/session"
	"gorm.io/gorm"
)

// Repository interface
type Repo interface {
	CreateSession(s *models.Session) error
	GetSession(token string) (*models.Session, error)
	DeleteSessionByToken(token string) error
	DeleteSession(s *models.Session) error
	Update(s *models.Session) (*models.Session, error)
}

type sessRepo struct {
	db *gorm.DB
}

// New user repo instance
func NewSessRepo(db *gorm.DB) Repo {
	return &sessRepo{
		db: db,
	}
}

func (repo *sessRepo) CreateSession(s *models.Session) error {

	return repo.db.Create(s).Error
}

// get first user by id
func (repo *sessRepo) GetSession(tok string) (*models.Session, error) {
	var s models.Session
	if err := repo.db.Where("token=?", tok).First(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

// delete all sessions for a user id
func (repo *sessRepo) DeleteSession(s *models.Session) error {
	if err := repo.db.Where("user_id = ?", s.UserID).Delete(&s).Error; err != nil {
		return err
	}

	return nil
}

// Delete all session for a token
func (repo *sessRepo) DeleteSessionByToken(tok string) error {
	if err := repo.db.Where("token = ?", tok).Delete(&models.Session{}).Error; err != nil {
		return err
	}

	return nil
}

// update user session in db
func (repo *sessRepo) Update(input_sess *models.Session) (*models.Session, error) {
	if err := repo.db.Where("user_id = ?", input_sess.UserID).Updates(input_sess).Error; err != nil {
		return nil, err
	}

	return input_sess, nil
}
