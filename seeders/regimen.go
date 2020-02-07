package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var userID0 = uuid.FromStringOrNil("ef837bfd-aae4-4495-8aec-5d70f3aa0ed3")
var userID1 = uuid.FromStringOrNil("8c8aa229-3959-4a40-bbe6-67c2eeace5cb")
var userID2 = uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd32d")
var podID0 = uuid.FromStringOrNil("938b0ff3-a272-49cc-b1c8-b09ccfd07792")
var podID1 = uuid.FromStringOrNil("16ea71db-2adb-45fe-a3fe-9e9ad6dcabd3")
var podID2 = uuid.FromStringOrNil("938b0ff3-a272-49cc-b1c8-b09ccfd07792")
var regimens = types.Regimens{
	{
		ID:        uuid.FromStringOrNil("b3d8bf48-3a8f-4ca5-b1e5-452a61be493a"),
		AccountID: uuid.FromStringOrNil("22b5123d-9cee-4701-b15b-8c9078142666"),
		UserID:    &userID0,
		PodID:     &podID0,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("395a37a3-a63b-49c4-94b4-04c29d47c64c"),
		AccountID: uuid.FromStringOrNil("d8e4c5dc-9767-41bd-b802-060e80d83867"),
		UserID:    &userID1,
		PodID:     &podID1,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("8ba3049b-17a1-4eae-b2bc-db7d18596d28"),
		AccountID: uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		UserID:    &userID2,
		PodID:     &podID2,
		Meta:      nil,
	},
}

func seedRegimens(db *gorm.DB) error {
	if !db.HasTable(&types.Regimen{}) {
		if err := db.AutoMigrate(&types.Regimen{}).Error; err != nil {
			return err
		}
	}
	for _, regimen := range regimens {
		if err := regimen.Create(db); err != nil {
			return err
		}
	}
	return nil
}
