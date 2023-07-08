package acc_repo

import (
	models "github.com/amar-jay/go-api-boilerplate/database/domain/account"
	"gorm.io/gorm"
)

// Repository interface
type Repo interface {
	CreateAccount(s *models.Account) error
	DeleteAccount(
		provider_type string,
		providerID string,
	) error
	GetAccountByProvider(type_ string, id string) (*models.Account, error)
}

type accRepo struct {
	db *gorm.DB
}

// New user repo instance
func NewAccountRepo(db *gorm.DB) Repo {
	return &accRepo{
		db: db,
	}
}

func (repo *accRepo) CreateAccount(s *models.Account) error {

	return repo.db.Create(s).Error
}

// Delete all session for a token
func (repo *accRepo) DeleteAccount(
	provider_type string,
	providerID string,
) error {
	if err := repo.db.Where("provider_type = ? AND provider_id = ?", provider_type, providerID).Delete(&models.Account{}).Error; err != nil {
		return err
	}

	return nil
}

// get first account by provider type and id
func (repo *accRepo) GetAccountByProvider(type_ string, id string) (*models.Account, error) {
	var s models.Account
	// TODO: rewrite this query in a better way
	if err := repo.db.Where("provider_type = ? AND provider_id = ?", type_, id).First(&s).Error; err != nil {
		return nil, err
	}
	// if err := repo.db.Where("provider_type = ? AND provider_id = ?", type_, id).First(&s).Error; err != nil {
	// 	return nil, err
	// }

	return &s, nil
}
