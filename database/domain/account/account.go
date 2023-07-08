package account

type Account struct {
	ID                string `gorm:"size:255" json:"id"`
	SID               string `gorm:"size:255" json:"sid"`
	UserID            string `gorm:"size:255" json:"userId"`
	ProviderID        string `gorm:"size:255" json:"providerId"`
	ProviderType      string `gorm:"size:255" json:"providerType"`
	AccountID         string `gorm:"size:255" json:"providerAccountId"`
	RefreshToken      string `gorm:"size:255" json:"refreshToken"`
	AccessToken       string `gorm:"size:255" json:"accessToken"`
	AcessTokenExpires string `gorm:"size:255" json:"accessTokenExpires"`
}
