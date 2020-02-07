package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var users = types.Users{
	{
		ID:         uuid.FromStringOrNil("ef837bfd-aae4-4495-8aec-5d70f3aa0ed3"),
		AccountID:  uuid.FromStringOrNil("22b5123d-9cee-4701-b15b-8c9078142666"),
		CognitoID:  uuid.NewV4(),
		ExternalID: 26000,
		FirstName:  "Bobby",
		LastName:   "Baratheon",
		Email:      "bestking@kingslanding.gov",
		Phone:      "800-555-0666",
		Gender:     "Male",
		Height:     192.9,
		Weight:     300,
		Meta:       nil,
	},
	{
		ID:        uuid.FromStringOrNil("8c8aa229-3959-4a40-bbe6-67c2eeace5cb"),
		AccountID: uuid.FromStringOrNil("d8e4c5dc-9767-41bd-b802-060e80d83867"),
		CognitoID: uuid.NewV4(),
		FirstName: "Ben'",
		LastName:  "Dickman",
		Email:     "ben@gettespo.com",
		Phone:     "800-555-0555",
		Gender:    "Male",
		Height:    167,
		Weight:    250,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd32d"),
		AccountID: uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		CognitoID: uuid.NewV4(),
		FirstName: "AyAyRon",
		LastName:  "Feys",
		Email:     "aaron@gettespo.com",
		Phone:     "800-555-0777",
		Gender:    "Male",
		Height:    200,
		Weight:    280,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd31b"),
		AccountID: uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		CognitoID: uuid.NewV4(),
		FirstName: "Developer",
		LastName:  "",
		Email:     "test2",
		Phone:     "800-555-0777",
		Gender:    "Male",
		Height:    200,
		Weight:    280,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd31a"),
		AccountID: uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
		CognitoID: uuid.NewV4(),
		FirstName: "Master",
		LastName:  "",
		Email:     "dev@gettespo.com",
		Phone:     "800-555-0777",
		Gender:    "Male",
		Height:    200,
		Weight:    280,
		Meta:      nil,
	},
	{
		ID:        uuid.FromStringOrNil("93fbb5ef-f0ee-4597-a51d-4400829f13ca"),
		AccountID: uuid.FromStringOrNil("b27c658a-f885-44a9-ab9d-d288f9d53138"),
		CognitoID: uuid.NewV4(),
		FirstName: "Lambda",
		LastName:  "Services",
		Email:     "dev+lambda@gettespo.com",
		Phone:     "",
		Gender:    "",
		Height:    0,
		Weight:    0,
		Meta:      nil,
	},
}

var userRoles = map[uuid.UUID][]uuid.UUID{
	uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd32d"): {
		uuid.FromStringOrNil("6d814967-ebdb-4ba7-b8e8-2a94ee5fafeb"),
	},
	uuid.FromStringOrNil("ef837bfd-aae4-4495-8aec-5d70f3aa0ed3"): {
		uuid.FromStringOrNil("6146f7d1-7f7c-4519-a243-dfe71ae2284d"),
		uuid.FromStringOrNil("9651b56b-9781-42a6-99b5-4db352474f2a"),
	},
	uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd31b"): {
		uuid.FromStringOrNil("d9be45aa-0d38-412e-8eee-cc0fe9b38239"),
		uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"),
	},
	uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd31a"): {
		uuid.FromStringOrNil("f4e17fdb-2728-4b61-af7f-ba220e89fa86"),
		uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"),
	},
	uuid.FromStringOrNil("93fbb5ef-f0ee-4597-a51d-4400829f13ca"): {
		uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"),
	},
}

func seedUsers(db *gorm.DB) error {
	if !db.HasTable(&types.User{}) {
		if err := db.AutoMigrate(&types.User{}).Error; err != nil {
			return err
		}
	}
	for _, user := range users {
		if err := user.Create(db); err != nil {
			return err
		}
	}
	for userID, roles := range userRoles {
		for _, roleID := range roles {
			if err := db.Table("role_user").Create(struct {
				UserID uuid.UUID
				RoleID uuid.UUID
			}{
				UserID: userID,
				RoleID: roleID,
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
