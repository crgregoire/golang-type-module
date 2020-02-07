package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthGrant is a table of all available grant types
//
type OauthGrant struct {
	ID           uuid.UUID     `json:"id" gorm:"type:char(36);primary_key"`
	Type         string        `json:"type"`
	OauthClients []OauthClient `json:"-" gorm:"many2many:oauth_client_grant;"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthGrant *OauthGrant) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthGrant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthGrants is a slice of OauthGrant
//
type OauthGrants []OauthGrant

//
// Get gets all oauthGrants
//
func (oauthGrants *OauthGrants) Get(db *gorm.DB) error {
	if err := db.Find(&oauthGrants).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthGrant by id
//
func (oauthGrant *OauthGrant) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthGrant).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthGrant by id
//
func (oauthGrant *OauthGrant) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthGrant).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthGrant
//
func (oauthGrant *OauthGrant) Create(db *gorm.DB) error {
	if err := db.Create(&oauthGrant).Scan(&oauthGrant).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthGrant
//
func (oauthGrant *OauthGrant) Update(db *gorm.DB) error {
	if err := db.Save(&oauthGrant).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthGrant
//
func (oauthGrant *OauthGrant) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthGrant).Error; err != nil {
		return err
	}
	return nil
}
