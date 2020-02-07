package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var usages = types.Usages{
	{
		ID:          uuid.FromStringOrNil("e34fe136-c804-4355-9b00-0be4f7093a62"),
		DispenserID: uuid.FromStringOrNil("142201c2-0c5f-4650-8c99-fc233412e030"),
		RegimenID:   uuid.FromStringOrNil("8ba3049b-17a1-4eae-b2bc-db7d18596d28"),
		Meta:        nil,
	},
}

func seedUsages(db *gorm.DB) error {
	if !db.HasTable(&types.Usage{}) {
		if err := db.AutoMigrate(&types.Usage{}).Error; err != nil {
			return err
		}
	}
	for _, usage := range usages {
		if err := usage.Create(db); err != nil {
			return err
		}
	}
	return nil
}