package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// Account is an account
//
type Account struct {
	ID          uuid.UUID    `json:"id" gorm:"type:char(36);primary_key"`
	ExternalID	string		 `json:"external_id"`
	Name        string       `json:"name" scope:"account.name" scope:"name"`
	Meta        JSON         `json:"meta" gorm:"type:json" scope:"account.meta" scope:"meta"`
	Users       []User       `json:"users" scope:"users"`
	Connections []Connection `json:"connections" scope:"connections"`
	Regimens    []Regimen    `json:"regimens"`
	CreatedAt   time.Time    `json:"created_at" scope:"account.created_at"`
	UpdatedAt   time.Time    `json:"updated_at" scope:"account.updated_at"`
	DeletedAt   *time.Time   `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if account.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Accounts is a slice of Account
//
type Accounts []Account

//
// Get gets all accounts
//
func (accounts *Accounts) Get(db *gorm.DB) error {
	if err := db.Find(&accounts).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one account by id
//
func (account *Account) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&account).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one account by id
//
func (account *Account) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&account).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Account
//
func (account *Account) Create(db *gorm.DB) error {
	if err := db.Create(&account).Scan(&account).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Account
//
func (account *Account) Update(db *gorm.DB) error {
	if err := db.Save(&account).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an account
//
func (account *Account) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&account).Error; err != nil {
		return err
	}
	return nil
}

//
// GetConnections gets the Connections associated with an account
//
func (account *Account) GetConnections(db *gorm.DB) error {
	if err := db.Find(&account).Related(&account.Connections).Error; err != nil {
		return err
	}
	return nil
}

//
// GetConnectionByID gets an account Connection by ID
//
func (account *Account) GetConnectionByID(db *gorm.DB, ConnectionID uuid.UUID) error {
	if err := db.Model(&account).Where("id = ? ", ConnectionID).Related(&account.Connections).Error; err != nil {
		return err
	}
	return nil
}

//
// UpdateAccountConnectionByID updates an account Connection by ID
//
func (account *Account) UpdateAccountConnectionByID(db *gorm.DB, Connection Connection) error {
	if err := db.Model(&account).Association("Connections").Append(Connection).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUsers gets the users associated with an account
//
func (account *Account) GetUsers(db *gorm.DB) error {
	if err := db.Find(&account).Related(&account.Users).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserByID gets an account user by ID
//
func (account *Account) GetUserByID(db *gorm.DB, userID uuid.UUID) error {
	if err := db.Model(&account).Where("id = ? ", userID).Related(&account.Users).Error; err != nil {
		return err
	}
	return nil
}

//
// UpdateAccountUserByID updates an account user by ID
//
func (account *Account) UpdateAccountUserByID(db *gorm.DB, user User) error {
	if err := db.Model(&account).Association("Users").Append(user).Error; err != nil {
		return err
	}
	return nil
}

//
// UpdateAccountRegimenByID updates a regimen associated with an account
//
func (account *Account) UpdateAccountRegimenByID(db *gorm.DB, regimen Regimen) error {
	if err := db.Model(&account).Association("Regimens").Append(regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// GetRegimens gets the regimens associated with an account
//
func (account *Account) GetRegimens(db *gorm.DB) error {
	if err := db.Model(&account).Association("Regimens").Find(&account.Regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (account *Account) Scope(scopes []string) {
	*account = scoping.FilterByScopes(scopes, *account).(Account)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (accounts *Accounts) Scope(scopes []string) {
	accountSlice := *accounts
	for i, account := range accountSlice {
		accountSlice[i] = scoping.FilterByScopes(scopes, account).(Account)
	}
	*accounts = accountSlice
}

//
// GetByQuery gets one user by id
//
func (account *Account) GetByQuery(db *gorm.DB, queryS string, value ...interface{}) error {
	if err := db.Where(queryS, value...).Find(&account).Error; err != nil {
		return err
	}
	return nil
}