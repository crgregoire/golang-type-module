package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// A Dispenser is our internet of things vitamin dispenser
//
type Dispenser struct {
	ID         uuid.UUID   `json:"id" gorm:"type:char(36);primary_key"`
	Serial     string      `json:"serial" scope:"dispenser.serial"`
	Name       string      `json:"name" scope:"dispenser.name"`
	Network    string      `json:"network" scope:"dispenser.network"`
	Insertions []Insertion `json:"insertions"`
	Meta       JSON        `json:"meta" gorm:"type:json" scope:"dispenser.meta"`
	CreatedAt  time.Time   `json:"created_at" scope:"dispenser.created_at"`
	UpdatedAt  time.Time   `json:"updated_at" scope:"dispenser.updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (dispenser *Dispenser) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if dispenser.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Dispensers is a slice of Dispenser
//
type Dispensers []Dispenser

//
// Get gets all dispensers
//
func (dispensers *Dispensers) Get(db *gorm.DB) error {
	if err := db.Find(&dispensers).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets Dispenser by ID
//
func (dispenser *Dispenser) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets Dispensers by query
//
func (dispensers *Dispensers) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).Find(&dispensers).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets Dispenser by query
//
func (dispenser *Dispenser) GetOneByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).First(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one dispenser by ID
//
func (dispenser *Dispenser) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// Create creates a Dispenser
//
func (dispenser *Dispenser) Create(db *gorm.DB) error {
	if err := db.Create(&dispenser).Scan(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// Update updates a Dispenser
//
func (dispenser *Dispenser) Update(db *gorm.DB) error {
	if err := db.Save(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete a Dispenser
//
func (dispenser *Dispenser) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&dispenser).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (dispenser *Dispenser) Scope(scopes []string) {
	*dispenser = scoping.FilterByScopes(scopes, *dispenser).(Dispenser)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (dispensers *Dispensers) Scope(scopes []string) {
	dispenserSlice := *dispensers
	for i, dispenser := range dispenserSlice {
		dispenserSlice[i] = scoping.FilterByScopes(scopes, dispenser).(Dispenser)
	}
	*dispensers = dispenserSlice
}
