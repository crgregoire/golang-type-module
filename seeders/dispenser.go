package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var dispensers = types.Dispensers{
	{
		ID:      uuid.FromStringOrNil("142201c2-0c5f-4650-8c99-fc233412e030"),
		Serial:  "dispenser-11218BIGPP",
		Name:    "Tespo Connect",
		Network: "UNIMATRIX0",
		Meta:    nil,
	},
	{
		ID:      uuid.FromStringOrNil("9f717337-dba7-415c-9daf-c607df526d14"),
		Serial:  "dispenser-11218SCOOT",
		Name:    "Tespo Connect",
		Network: "UNIMATRIX1",
		Meta:    nil,
	},
}

func seedDispensers(db *gorm.DB) error {
	if !db.HasTable(&types.Dispenser{}) {
		if err := db.AutoMigrate(&types.Dispenser{}).Error; err != nil {
			return err
		}
	}
	for _, dispenser := range dispensers {
		if err := dispenser.Create(db); err != nil {
			return err
		}
	}
	return nil
}