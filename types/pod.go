package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
)

//
// A Pod is a pod
//
type Pod struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key" scope:"pod.id"`
	Name      string     `json:"name" scope:"pod.name"`
	Sku       string     `json:"sku" scope:"pod.sku"`
	Slug      string     `json:"slug" scope:"pod.slug"`
	Color     string     `json:"color" scope:"pod.color"`
	Cells     uint       `json:"cells" scope:"pod.cells"`
	LabelTall string     `json:"label_tall" scope:"pod.label_tall"`
	LabelWide string     `json:"label_wide" scope:"pod.label_wide"`
	Meta      JSON       `json:"meta" gorm:"type:json" scope:"pod.meta"`
	CreatedAt time.Time  `json:"created_at" scope:"pod.created_at"`
	UpdatedAt time.Time  `json:"updated_at" scope:"pod.updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (pod *Pod) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if pod.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Pods is a slice of Pod
//
type Pods []Pod

//
// Get gets all pods
//
func (pods *Pods) Get(db *gorm.DB) error {
	if err := db.Find(&pods).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets Pod by ID
//
func (pod *Pod) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&pod).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets Pod by Query
//
func (pods *Pods) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).Find(&pods).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets Pod by Query
//
func (pod *Pod) GetOneByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).First(&pod).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one pod by ID
//
func (pod *Pod) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&pod).Error; err != nil {
		return err
	}
	return nil
}

//
// Create creates a Pod
//
func (pod *Pod) Create(db *gorm.DB) error {
	if err := db.Create(&pod).Scan(&pod).Error; err != nil {
		return err
	}
	return nil
}

//
// Update updates a Pod
//
func (pod *Pod) Update(db *gorm.DB) error {
	if err := db.Save(&pod).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete a Pod
//
func (pod *Pod) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&pod).Error; err != nil {
		return err
	}
	return nil
}
