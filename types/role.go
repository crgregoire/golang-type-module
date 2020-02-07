package types

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
)

//
// Role for a role
//
type Role struct {
	ID          uuid.UUID     `json:"id" gorm:"type:char(36);primary_key"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Permissions []Permission  `json:"permission" gorm:"many2many:permission_role;"`
	User        []User        `json:"-" gorm:"many2many:role_user;"`
	OauthClient []OauthClient `json:"-" gorm:"many2many:oauth_client_role;"`
	Meta        JSON          `json:"meta" gorm:"type:json"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
	DeletedAt   *time.Time    `json:"-"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (role *Role) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if role.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Roles is a slice of Role
//
type Roles []Role

//
// Get gets all roles
//
func (roles *Roles) Get(db *gorm.DB) error {
	if err := db.Find(&roles).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one role by id
//
func (role *Role) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&role).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one role by id
//
func (role *Role) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&role).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Role
//
func (role *Role) Create(db *gorm.DB) error {
	if err := db.Create(&role).Scan(&role).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Role
//
func (role *Role) Update(db *gorm.DB) error {
	if err := db.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an role
//
func (role *Role) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

//
// GetRolePermissions get's the role with the permissions
// attached
//
func (roles *Roles) GetRolePermissions(db *gorm.DB) (*PermissionsClaim, error) {
	ids := []uuid.UUID{}
	for _, role := range *roles {
		ids = append(ids, role.ID)
	}
	var permissions Permissions
	if err := db.Where("permission_role.role_id IN (?)", ids).Model(&roles).Related(&permissions, "Permissions").Error; err != nil {
		return nil, err
	}
	permissionsClaim := PermissionsClaim{}
	for _, permission := range permissions {
		data, _ := permission.Actions.MarshalJSON()
		actions := []string{}
		if err := json.Unmarshal(data, &actions); err != nil {
			return nil, err
		}
		for _, action := range actions {
			permissionsClaim[action] = append(permissionsClaim[action], permission.Slug)
		}
	}
	return &permissionsClaim, nil
}

//
// AddPermissionToRoleByID associates a role to a permission
//
func (role *Role) AddPermissionToRoleByID(db *gorm.DB, permissionID uuid.UUID) error {
	permission := Permission{}
	if err := db.Where("id = ?", permissionID).Find(&permission).Error; err != nil {
		return err
	}
	if err := db.Model(&role).Association("Permissions").Append(permission).Error; err != nil {
		return err
	}
	return nil
}
