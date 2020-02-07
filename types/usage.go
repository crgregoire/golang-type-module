package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// Usage data model for usage tracking
//
type Usage struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primary_key" scope:"usage.id"`
	RegimenID   uuid.UUID  `json:"regimen_id" gorm:"type:char(36);foreign_key" scope:"usage.regimen_id"`
	DispenserID uuid.UUID  `json:"dispenser_id" gorm:"type:char(36);foreign_key" scope:"usage.dispenser_id"`
	UserID      *uuid.UUID `json:"user_id" gorm:"type:char(36);foreign_key" scope:"usage.user_id"`
	BarcodeID   *uuid.UUID `json:"barcode_id" gorm:"type:char(36);foreign_key" scope:"usage.barcode_id"`
	Flags       uint       `json:"flags" scope:"usage.flags"`
	Servings    uint       `json:"servings" scope:"usage.servings"`
	Meta        JSON       `json:"meta" gorm:"type:json" scope:"usage.meta"`
	CreatedAt   time.Time  `json:"created_at" scope:"usage.created_at"`
	UpdatedAt   time.Time  `json:"updated_at" scope:"usage.updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (usage *Usage) BeforeCreate(scope *gorm.Scope) error {

	if usage.UserID != nil {
		if usage.UserID.String() == "00000000-0000-0000-0000-000000000000" {
			if err := scope.SetColumn("UserID", nil); err != nil {
				return err
			}
		}
	}
	if usage.BarcodeID != nil {
		if usage.BarcodeID.String() == "00000000-0000-0000-0000-000000000000" {
			if err := scope.SetColumn("BarcodeID", nil); err != nil {
				return err
			}
		}
	}

	if usage.ID.String() == "00000000-0000-0000-0000-000000000000" {
		if err := scope.SetColumn("ID", uuid.NewV4()); err != nil {
			return err
		}
	}
	return nil
}

//
// Usages is a slice of Usage
//
type Usages []Usage

//
// Get gets all usages
//
func (usages *Usages) Get(db *gorm.DB) error {
	if err := db.Find(&usages).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one usage by id
//
func (usage *Usage) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets all usages by given query
//
func (usages *Usages) GetByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).Find(&usages).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets all usages by given query
//
func (usage *Usage) GetByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).Find(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one usage by id
//
func (usage *Usage) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Usage
//
func (usage *Usage) Create(db *gorm.DB) error {
	if err := db.Create(&usage).Scan(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Usage
//
func (usage *Usage) Update(db *gorm.DB) error {
	if err := db.Save(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an usage
//
func (usage *Usage) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&usage).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (usage *Usage) Scope(scopes []string) {
	*usage = scoping.FilterByScopes(scopes, *usage).(Usage)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (usages *Usages) Scope(scopes []string) {
	usageSlice := *usages
	for i, usage := range usageSlice {
		usageSlice[i] = scoping.FilterByScopes(scopes, usage).(Usage)
	}
	*usages = usageSlice
}
