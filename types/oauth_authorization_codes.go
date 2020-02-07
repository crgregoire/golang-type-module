package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthAuthorizationCode is an outside party signed up to use
// our data
//
type OauthAuthorizationCode struct {
	ID            uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	UserID        uuid.UUID  `json:"user_id" gorm:"type:char(36)"`
	OauthClientID string     `json:"client_id" gorm:"type:char(36)"`
	Code          string     `json:"code"`
	GrantType     string     `json:"grant_type"`
	RedirectURI   string     `json:"redirect_uri"`
	ExpiresAt     time.Time  `json:"expires_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthAuthorizationCode *OauthAuthorizationCode) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthAuthorizationCode.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthAuthorizationCodes is a slice of OauthAuthorizationCode
//
type OauthAuthorizationCodes []OauthAuthorizationCode

//
// Get gets all oauthAuthorizationCodes
//
func (oauthAuthorizationCodes *OauthAuthorizationCodes) Get(db *gorm.DB) error {
	if err := db.Find(&oauthAuthorizationCodes).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthAuthorizationCode by id
//
func (oauthAuthorizationCode *OauthAuthorizationCode) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one oauthAuthorizationCode by query
//
func (oauthAuthorizationCode *OauthAuthorizationCode) GetOneByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).Find(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthAuthorizationCode by id
//
func (oauthAuthorizationCode *OauthAuthorizationCode) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthAuthorizationCode
//
func (oauthAuthorizationCode *OauthAuthorizationCode) Create(db *gorm.DB) error {
	if err := db.Create(&oauthAuthorizationCode).Scan(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthAuthorizationCode
//
func (oauthAuthorizationCode *OauthAuthorizationCode) Update(db *gorm.DB) error {
	if err := db.Save(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthAuthorizationCode
//
func (oauthAuthorizationCode *OauthAuthorizationCode) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthAuthorizationCode).Error; err != nil {
		return err
	}
	return nil
}
