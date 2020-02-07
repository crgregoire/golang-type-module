package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var accounts = types.Accounts{
	{
		ID:         uuid.FromStringOrNil("b27c658a-f885-44a9-ab9d-d288f9d53138"),
		ExternalID: "1",
		Name:       "Tespo",
		Meta:       nil,
	},
	{
		ID:         uuid.FromStringOrNil("22b5123d-9cee-4701-b15b-8c9078142666"),
		ExternalID: "2",
		Name:       "Billy Bob",
		Meta:       nil,
	},
	{
		ID:         uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		ExternalID: "3",
		Name:       "Taco Sauce",
		Meta:       nil,
	},
	{
		ID:         uuid.FromStringOrNil("d8e4c5dc-9767-41bd-b802-060e80d83867"),
		ExternalID: "4",
		Name:       "AYAYRON",
	},
}

func seedAccounts(db *gorm.DB) error {
	if !db.HasTable(&types.Account{}) {
		if err := db.AutoMigrate(&types.Account{}).Error; err != nil {
			return err
		}
	}
	for _, account := range accounts {
		if err := account.Create(db); err != nil {
			return err
		}
	}
	return nil
}
