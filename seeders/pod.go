package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var pods = types.Pods{
	{
		ID:        uuid.FromStringOrNil("16ea71db-2adb-45fe-a3fe-9e9ad6dcabd3"),
		Name:      "Men's Complete",
		Slug:      "mens-complete",
		Color:     "#2665A5",
		Cells:     31,
		LabelTall: "https://gettepso.com/menslabeltall",
		LabelWide: "https://gettepso.com/menslabelwide",
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("938b0ff3-a272-49cc-b1c8-b09ccfd07792"),
		Name:      "Women's Complete",
		Slug:      "womens-complete",
		Color:     "#C9287E",
		Cells:     28,
		LabelTall: "https://gettepso.com/womenslabeltall",
		LabelWide: "https://gettepso.com/womenslabelwide",
		Meta:      nil,
	},
}

func seedPods(db *gorm.DB) error {
	if !db.HasTable(&types.Pod{}) {
		if err := db.AutoMigrate(&types.Pod{}).Error; err != nil {
			return err
		}
	}
	for _, pod := range pods {
		if err := pod.Create(db); err != nil {
			return err
		}
	}
	return nil
}
