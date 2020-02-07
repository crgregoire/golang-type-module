package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
)

//
// Permission based on permission
//
type Permission struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Actions   JSON       `json:"actions"`
	Meta      JSON       `json:"meta" gorm:"type:json"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (permission *Permission) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if permission.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Permissions is a slice of Permission
//
type Permissions []Permission

//
// Get gets all permissions
//
func (permissions *Permissions) Get(db *gorm.DB) error {
	if err := db.Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one permission by id
//
func (permission *Permission) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&permission).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one permission by id
//
func (permission *Permission) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&permission).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Permission
//
func (permission *Permission) Create(db *gorm.DB) error {
	if err := db.Create(&permission).Scan(&permission).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Permission
//
func (permission *Permission) Update(db *gorm.DB) error {
	if err := db.Save(&permission).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an permission
//
func (permission *Permission) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&permission).Error; err != nil {
		return err
	}
	return nil
}
