package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthClientUserRefreshToken is a collection of tokens issued to an
// oauth client
//
type OauthClientUserRefreshToken struct {
	ID                 uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	OauthAccessTokenID uuid.UUID  `json:"oauth_access_token_id"`
	Token              string     `json:"token" gorm:"type:text;size:999999"`
	ExpiresAt          time.Time  `json:"expires_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthClientUserRefreshToken.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthClientUserRefreshTokens is a slice of OauthClientUserRefreshToken
//
type OauthClientUserRefreshTokens []OauthClientUserRefreshToken

//
// Get gets all oauthClientUserRefreshTokens
//
func (oauthClientUserRefreshTokens *OauthClientUserRefreshTokens) Get(db *gorm.DB) error {
	if err := db.Find(&oauthClientUserRefreshTokens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthClientUserRefreshToken by id
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one oauthClientUserRefreshToken by query
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) GetOneByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).First(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthClientUserRefreshToken by id
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthClientUserRefreshToken
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) Create(db *gorm.DB) error {
	if err := db.Create(&oauthClientUserRefreshToken).Scan(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthClientUserRefreshToken
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) Update(db *gorm.DB) error {
	if err := db.Save(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthClientUserRefreshToken
//
func (oauthClientUserRefreshToken *OauthClientUserRefreshToken) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthClientUserRefreshToken).Error; err != nil {
		return err
	}
	return nil
}
