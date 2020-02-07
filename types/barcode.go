package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// Barcode describes a barcode on a Pod
//
type Barcode struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key" scope:"barcode.id"`
	PodID     uuid.UUID  `json:"pod_id" gorm:"type:char(36);foreign_key" scope:"barcode.pod_id"`
	Sku       string     `json:"sku" scope:"barcode.sku"`
	Code      string     `json:"code" scope:"barcode.code"`
	LabelTall string     `json:"label_tall" scope:"barcode.label_tall"`
	LabelWide string     `json:"label_wide" scope:"barcode.label_wide"`
	CreatedAt time.Time  `json:"created_at" scope:"barcode.created_at"`
	UpdatedAt time.Time  `json:"updated_at" scope:"barcode.updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (barcode *Barcode) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if barcode.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Barcodes is a slice of Barcode
//
type Barcodes []Barcode

//
// Get gets all barcodes
//
func (barcodes *Barcodes) Get(db *gorm.DB) error {
	if err := db.Find(&barcodes).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one barcode by id
//
func (barcode *Barcode) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one barcode by id
//
func (barcode *Barcode) GetOneByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).First(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets one barcode by id
//
func (barcodes *Barcodes) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).Find(&barcodes).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one barcode by id
//
func (barcode *Barcode) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Barcode
//
func (barcode *Barcode) Create(db *gorm.DB) error {
	if err := db.Create(&barcode).Scan(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Barcode
//
func (barcode *Barcode) Update(db *gorm.DB) error {
	if err := db.Save(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an barcode
//
func (barcode *Barcode) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&barcode).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (barcode *Barcode) Scope(scopes []string) {
	*barcode = scoping.FilterByScopes(scopes, *barcode).(Barcode)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (barcodes *Barcodes) Scope(scopes []string) {
	barcodeSlice := *barcodes
	for i, barcode := range barcodeSlice {
		barcodeSlice[i] = scoping.FilterByScopes(scopes, barcode).(Barcode)
	}
	*barcodes = barcodeSlice
}
