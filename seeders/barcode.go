package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var barcodes = types.Barcodes{
	{
		ID:        uuid.FromStringOrNil("82a0757e-9d26-4e46-8f22-4105d185acd0"),
		PodID:     uuid.FromStringOrNil("16ea71db-2adb-45fe-a3fe-9e9ad6dcabd3"),
		Code:      "000120150016",
		LabelTall: "https://gettespo.com/menslabeltall",
		LabelWide: "https://gettespo.com/menslabelwide",
	},
	{
		ID:        uuid.FromStringOrNil("62012ee1-f3c2-4a4b-a0ab-8686e3d173e4"),
		PodID:     uuid.FromStringOrNil("938b0ff3-a272-49cc-b1c8-b09ccfd07792"),
		Code:      "000640280018",
		LabelTall: "https://gettespo.com/womenslabeltall",
		LabelWide: "https://gettespo.com/womenslabelwide",
	},
}

func seedBarcodes(db *gorm.DB) error {
	if !db.HasTable(&types.Barcode{}) {
		if err := db.AutoMigrate(&types.Barcode{}).Error; err != nil {
			return err
		}
	}
	for _, barcode := range barcodes {
		if err := barcode.Create(db); err != nil {
			return err
		}
	}
	return nil
}
