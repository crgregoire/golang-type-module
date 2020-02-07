package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var reminders = types.Reminders{
	{
		ID:        uuid.FromStringOrNil("0b703c2b-72c6-43ef-a089-4ce85e06519a"),
		UserID:    uuid.FromStringOrNil("8c8aa229-3959-4a40-bbe6-67c2eeace5cb"),
		RegimenID: uuid.FromStringOrNil("8ba3049b-17a1-4eae-b2bc-db7d18596d28"),
		Minute:    750,
	},
}

func seedReminders(db *gorm.DB) error {
	if !db.HasTable(&types.Reminder{}) {
		if err := db.AutoMigrate(&types.Reminder{}).Error; err != nil {
			return err
		}
	}
	for _, reminder := range reminders {
		if err := reminder.Create(db); err != nil {
			return err
		}
	}
	return nil
}
