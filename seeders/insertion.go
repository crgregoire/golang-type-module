package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	types "github.com/tespo/satya/v2/types"
)

var barcodeID = uuid.FromStringOrNil("62012ee1-f3c2-4a4b-a0ab-8686e3d173e4")
var insertions = types.Insertions{
	{
		ID:          uuid.FromStringOrNil("79cf3dce-2af0-4e17-afa8-7fd7deed15c6"),
		DispenserID: uuid.FromStringOrNil("142201c2-0c5f-4650-8c99-fc233412e030"),
		RegimenID:   uuid.FromStringOrNil("8ba3049b-17a1-4eae-b2bc-db7d18596d28"),
		BarcodeID:   &barcodeID,
		Flags:       5,
		Servings:    31,
		LabelTall:   "https://gettespo.com/menscompletetall",
		LabelWide:   "https://gettespo.com/menscompletewide",
		Meta:        nil,
	},
}

func seedInsertions(db *gorm.DB) error {
	if !db.HasTable(&types.Insertion{}) {
		if err := db.AutoMigrate(&types.Insertion{}).Error; err != nil {
			return err
		}
	}
	for _, insertion := range insertions {
		if err := insertion.Create(db); err != nil {
			return err
		}
	}
	return nil
}
