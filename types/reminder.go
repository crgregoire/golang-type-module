package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// A Reminder is a reminder for a user to take their vitamins
//
type Reminder struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key;not null" scope:"reminder.id"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:char(36);foreign_key" scope:"reminder.user_id"`
	RegimenID uuid.UUID  `json:"regimen_id" gorm:"type:char(36);foreign_key" scope:"reminder.regimen_id"`
	Minute    uint       `json:"minute" scope:"reminder.minute"`
	Meta      JSON       `json:"meta" scope:"reminder.meta"`
	CreatedAt time.Time  `json:"created_at" scope:"reminder.created_at"`
	UpdatedAt time.Time  `json:"updated_at" scope:"reminder.updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (reminder *Reminder) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if reminder.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Reminders is a slice of Reminder
//
type Reminders []Reminder

//
// Get gets all reminders
//
func (reminders *Reminders) Get(db *gorm.DB) error {
	if err := db.Find(&reminders).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one reminder by id
//
func (reminder *Reminder) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery returns the users reminders
//
func (reminder *Reminder) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).First(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery returns the users reminders
//
func (reminders *Reminders) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).Find(&reminders).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserReminders returns the users reminders
//
func (reminders *Reminders) GetUserReminders(db *gorm.DB, userID uuid.UUID) error {
	if err := db.Find(&reminders, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one reminder by id
//
func (reminder *Reminder) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Reminder
//
func (reminder *Reminder) Create(db *gorm.DB) error {
	if err := db.Create(&reminder).Scan(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Reminder
//
func (reminder *Reminder) Update(db *gorm.DB) error {
	if err := db.Save(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an reminder
//
func (reminder *Reminder) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (reminder *Reminder) Scope(scopes []string) {
	*reminder = scoping.FilterByScopes(scopes, *reminder).(Reminder)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (reminders *Reminders) Scope(scopes []string) {
	reminderSlice := *reminders
	for i, reminder := range reminderSlice {
		reminderSlice[i] = scoping.FilterByScopes(scopes, reminder).(Reminder)
	}
	*reminders = reminderSlice
}
