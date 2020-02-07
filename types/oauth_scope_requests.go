package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthScopeRequest is a record of when a 3rd party
// application is requesting new scopes
//
type OauthScopeRequest struct {
	ID            uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	Scopes        JSON       `json:"scopes" gorm:"type:json"`
	OauthClientID uuid.UUID  `json:"oauth_client_id" gorm:"unique;not null"`
	Meta          JSON       `json:"meta" gorm:"type:json"`
	RequestedAt   time.Time  `json:"requested_at"`
	ApprovedAt    time.Time  `json:"approved_at"`
	DeniedAt      time.Time  `json:"denied_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthScopeRequest *OauthScopeRequest) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthScopeRequest.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthScopeRequests is a slice of OauthScopeRequest
//
type OauthScopeRequests []OauthScopeRequest

//
// Get gets all oauthScopeRequests
//
func (oauthScopeRequests *OauthScopeRequests) Get(db *gorm.DB) error {
	if err := db.Find(&oauthScopeRequests).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthScopeRequest by id
//
func (oauthScopeRequest *OauthScopeRequest) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthScopeRequest).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthScopeRequest by id
//
func (oauthScopeRequest *OauthScopeRequest) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthScopeRequest).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthScopeRequest
//
func (oauthScopeRequest *OauthScopeRequest) Create(db *gorm.DB) error {
	if err := db.Create(&oauthScopeRequest).Scan(&oauthScopeRequest).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthScopeRequest
//
func (oauthScopeRequest *OauthScopeRequest) Update(db *gorm.DB) error {
	if err := db.Save(&oauthScopeRequest).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthScopeRequest
//
func (oauthScopeRequest *OauthScopeRequest) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthScopeRequest).Error; err != nil {
		return err
	}
	return nil
}
