package seeders

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var connections = types.Connections{
	{
		ID:          uuid.FromStringOrNil("5da4d0fc-e27c-4e0c-94fc-8ba3126908c8"),
		AccountID:   uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		DispenserID: uuid.FromStringOrNil("142201c2-0c5f-4650-8c99-fc233412e030"),
		Meta:        nil,
		ConnectedAt: time.Now(),
	},
}

func seedConnections(db *gorm.DB) error {
	if !db.HasTable(&types.Connection{}) {
		if err := db.AutoMigrate(&types.Connection{}).Error; err != nil {
			return err
		}
	}
	for _, connection := range connections {
		if err := connection.Create(db); err != nil {
			return err
		}
	}
	return nil
}
