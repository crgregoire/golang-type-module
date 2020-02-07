package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthAccessToken is a collection of tokens issued to an
// oauth client
//
type OauthAccessToken struct {
	ID               uuid.UUID  `json:"id" gorm:"type:char(36)"`
	UserID           uuid.UUID  `json:"user_id" gorm:"type:char(36);primary_key"`
	OauthClientID    string     `json:"oauth_client_id" gorm:"primary_key"`
	Token            string     `json:"token" gorm:"type:text;size:999999"`
	ScopePermissions JSON       `json:"scope_permissions" gorm:"type:json"`
	ScopedFields     JSON       `json:"scoped_fields" gorm:"type:json"`
	Revoked          bool       `json:"revoked"`
	ExpiresAt        time.Time  `json:"expires_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthAccessToken *OauthAccessToken) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthAccessToken.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthAccessTokens is a slice of OauthAccessToken
//
type OauthAccessTokens []OauthAccessToken

//
// Get gets all oauthAccessTokens
//
func (oauthAccessTokens *OauthAccessTokens) Get(db *gorm.DB) error {
	if err := db.Find(&oauthAccessTokens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthAccessToken by id
//
func (oauthAccessToken *OauthAccessToken) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one oauthAccessToken by id
//
func (oauthAccessToken *OauthAccessToken) GetOneByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).Find(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthAccessToken by id
//
func (oauthAccessToken *OauthAccessToken) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthAccessToken
//
func (oauthAccessToken *OauthAccessToken) Create(db *gorm.DB) error {
	if err := db.Create(&oauthAccessToken).Scan(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthAccessToken
//
func (oauthAccessToken *OauthAccessToken) Update(db *gorm.DB) error {
	if err := db.Save(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthAccessToken
//
func (oauthAccessToken *OauthAccessToken) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthAccessToken).Error; err != nil {
		return err
	}
	return nil
}
