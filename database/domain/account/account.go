package account

import "gorm.io/gorm"

type Account struct {
	ID                 string `gorm:"size:255" json:"id"`
	SID                string `gorm:"size:255" json:"sid"`
	UserId             string `gorm:"size:255" json:"userId"`
	ProviderId         string `gorm:"size:255" json:"providerId"`
	ProviderType       string `gorm:"size:255" json:"providerType"`
	AccountID          string `gorm:"size:255" json:"providerAccountId"`
	RefreshToken       string `gorm:"size:255" json:"refreshToken"`
	AccessToken        string `gorm:"size:255" json:"accessToken"`
	AccessTokenExpires string `gorm:"size:255" json:"accessTokenExpires"`
	gorm.Model
}
