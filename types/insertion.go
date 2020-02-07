package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// An Insertion is the data model for pod insertion
//
type Insertion struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	DispenserID uuid.UUID  `json:"dispenser_id" gorm:"type:char(36);foreign_key"`
	RegimenID   uuid.UUID  `json:"regimen_id" gorm:"type:char(36);foreign_key" scope:"insertion.regimen_id"`
	BarcodeID   *uuid.UUID `json:"barcode_id" gorm:"type:char(36);foreign_key" scope:"insertion.barcode_id"`
	Flags       uint       `json:"flags" scope:"insertion.flags"`
	Servings    uint       `json:"servings" scope:"insertion.servings"`
	LabelTall   string     `json:"label_tall" scope:"insertion.label_tall"`
	LabelWide   string     `json:"label_wide" scope:"insertion.label_wide"`
	Regimen     Regimen    `json:"regimen" gorm:"-"`
	Meta        JSON       `json:"meta" gorm:"type:json" scope:"insertion.meta"`
	CreatedAt   time.Time  `json:"created_at" scope:"insertion.created_at"`
	UpdatedAt   time.Time  `json:"updated_at" scope:"insertion.updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (insertion *Insertion) BeforeCreate(scope *gorm.Scope) error {
	if insertion.BarcodeID != nil {
		if insertion.BarcodeID.String() == "00000000-0000-0000-0000-000000000000" {
			if err := scope.SetColumn("BarcodeID", nil); err != nil {
				return err
			}
		}
	}
	uuid := uuid.NewV4()
	if insertion.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Insertions is a slice of Insertion
//
type Insertions []Insertion

//
// Get gets all insertions
//
func (insertions *Insertions) Get(db *gorm.DB) error {
	if err := db.Find(&insertions).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one insertion by id
//
func (insertion *Insertion) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one insertion by id
//
func (insertion *Insertion) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Insertion
//
func (insertion *Insertion) Create(db *gorm.DB) error {
	if err := db.Create(&insertion).Scan(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Insertion
//
func (insertion *Insertion) Update(db *gorm.DB) error {
	if err := db.Save(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an insertion
//
func (insertion *Insertion) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets the regimens associate withe an account
//
func (insertion *Insertion) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).Find(&insertion).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (insertion *Insertion) Scope(scopes []string) {
	*insertion = scoping.FilterByScopes(scopes, *insertion).(Insertion)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (insertions *Insertions) Scope(scopes []string) {
	insertionSlice := *insertions
	for i, insertion := range insertionSlice {
		insertionSlice[i] = scoping.FilterByScopes(scopes, insertion).(Insertion)
	}
	*insertions = insertionSlice
}
