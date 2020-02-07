package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// Invitation is used as our schema
// record of the invitations table
//
type Invitation struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key" scope:"invitation.id"`
	AccountID uuid.UUID  `json:"account_id" scope:"invitation.account_id"`
	Code      string     `json:"code" scope:"invitation.code"`
	Email     string     `json:"email" scope:"invitation.email"`
	Meta      JSON       `json:"meta" scope:"invitation.meta"`
	Account   Account    `json:"-"`
	ExpiresAt time.Time  `json:"expires_at" scope:"invitation.expires_at"`
	CreatedAt time.Time  `json:"created_at" scope:"invitation.created_at"`
	UpdatedAt time.Time  `json:"updated_at" scope:"invitation.updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (invitation *Invitation) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if invitation.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Invitations is a slice of Invitation
//
type Invitations []Invitation

//
// Get gets all invitations
//
func (invitations *Invitations) Get(db *gorm.DB) error {
	if err := db.Find(&invitations).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one invitation by id
//
func (invitation *Invitation) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&invitation).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one invitation by id
//
func (invitation *Invitation) GetOneByQuery(db *gorm.DB, queryS string, value ...interface{}) error {
	if err := db.Where(queryS, value...).First(&invitation).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets one invitation by id
//
func (invitations *Invitations) GetByQuery(db *gorm.DB, queryS string, value ...interface{}) error {
	if err := db.Where(queryS, value...).Find(&invitations).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one invitation by id
//
func (invitation *Invitation) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&invitation).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Invitation
//
func (invitation *Invitation) Create(db *gorm.DB) error {
	if err := db.Create(&invitation).Scan(&invitation).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Invitation
//
func (invitation *Invitation) Update(db *gorm.DB) error {
	if err := db.Save(&invitation).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an invitation
//
func (invitation *Invitation) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&invitation).Error; err != nil {
		return err
	}
	return nil
}
