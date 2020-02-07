package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// A User is linked to an user
//
type User struct {
	ID                      uuid.UUID                `json:"id" gorm:"type:char(36);primary_key;not null" scope:"user.id"`
	AccountID               uuid.UUID                `json:"account_id" gorm:"type:char(36);foreign_key" scope:"user.account_id"`
	CognitoID               uuid.UUID                `json:"cognito_id" gorm:"type:char(36);" scope:"user.cognito_id"`
	ExternalID              uint                     `json:"external_id" scope:"user.external_id"`
	FirstName               string                   `json:"first_name" scope:"user.first_name"`
	LastName                string                   `json:"last_name" scope:"user.last_name"`
	Email                   string                   `json:"email" scope:"user.email"`
	Phone                   string                   `json:"phone" scope:"user.phone"`
	Gender                  string                   `json:"gender" scope:"user.gender"`
	Height                  float32                  `json:"height" scope:"user.height"`
	Weight                  float32                  `json:"weight" scope:"user.weight"`
	Owner                   bool                     `json:"owner"`
	Meta                    JSON                     `json:"meta" gorm:"type:json" scope:"user.meta"`
	Roles                   Roles                    `gorm:"many2many:role_user;"`
	Regimens                Regimens                 `json:"regimen" scope:"user.regimens"`
	Usages                  Usages                   `json:"usages" scope:"user.usages"`
	OauthAuthorizationCodes []OauthAuthorizationCode `json:"oauth_authorization_code" scope:"user.oauth_authorization_codes"`
	CreatedAt               time.Time                `json:"created_at" scope:"user.created_at"`
	UpdatedAt               time.Time                `json:"updated_at" scope:"user.updated_at"`
	DeletedAt               *time.Time               `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// AfterCreate set up some additional associations
//
func (user *User) AfterCreate(scope *gorm.Scope) error {
	if err := scope.DB().Table("role_user").Create(struct {
		UserID uuid.UUID
		RoleID uuid.UUID
	}{
		UserID: user.ID,
		RoleID: uuid.FromStringOrNil("26841819-b845-4ecb-aa71-ebbbbfc60a6e"),
	}).Error; err != nil {
		return err
	}
	if user.Owner {
		if err := scope.DB().Table("role_user").Create(struct {
			UserID uuid.UUID
			RoleID uuid.UUID
		}{
			UserID: user.ID,
			RoleID: uuid.FromStringOrNil("0109f277-66cd-46f0-9107-7aa4ecb2201d"),
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

//
// Users is a slice of User
//
type Users []User

//
// Get gets all users
//
func (users *Users) Get(db *gorm.DB) error {
	if err := db.Find(&users).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one user by id
//
func (user *User) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets one user by id
//
func (user *User) GetByQuery(db *gorm.DB, queryS string, value ...interface{}) error {
	if err := db.Where(queryS, value...).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQueryWithRoles gets one user by id
//
func (user *User) GetByQueryWithRoles(db *gorm.DB, queryS string, value ...interface{}) error {
	if err := db.Where(queryS, value...).Find(&user).Error; err != nil {
		return err
	}
	if err := db.Find(&user).Association("Roles").Find(&user.Roles).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one user by id
//
func (user *User) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an User
//
func (user *User) Create(db *gorm.DB) error {
	if err := db.Create(&user).Scan(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an User
//
func (user *User) Update(db *gorm.DB) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an user
//
func (user *User) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserByEmailWithRoles returns user roles
//
func (user *User) GetUserByEmailWithRoles(db *gorm.DB, email string) error {
	if err := db.First(&user, "email = ? ", email).Error; err != nil {
		return err
	}
	if err := db.First(&user, "email = ? ", email).Association("Roles").Find(&user.Roles).Error; err != nil {
		return err
	}
	return nil
}

//
// GetRegimens gets the regimens associate withe an account
//
func (user *User) GetRegimens(db *gorm.DB) error {
	if err := db.Model(&user).Association("Regimens").Find(&user.Regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserWithAllData gets all of the data associated with the user
//
func (user *User) GetUserWithAllData(db *gorm.DB) error {
	if err := db.Find(&user).Preload("Usages").Preload("Reminders").Association("Regimens").Find(&user.Regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (user *User) Scope(scopes []string) {
	*user = scoping.FilterByScopes(scopes, *user).(User)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (users *Users) Scope(scopes []string) {
	u := *users
	for i, user := range u {
		u[i] = scoping.FilterByScopes(scopes, user).(User)
	}
	*users = u
}
