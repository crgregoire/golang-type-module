package types

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//
// OauthClient is an outside party signed up to use
// our data
//
type OauthClient struct {
	ID               uuid.UUID    `json:"id" gorm:"type:char(36);primary_key"`
	Name             string       `json:"name" gorm:"unique;not null"`
	ClientID         string       `json:"client_id" gorm:"unique;not null"`
	ClientSecret     string       `json:"client_secret" gorm:"unique;not null"`
	RedirectURI      string       `json:"redirect_uri"`
	ScopePermissions JSON         `json:"scope_permissions" gorm:"type:json"`
	ScopedFields     JSON         `json:"scoped_fields" gorm:"type:json"`
	User             []Users      `json:"-"`
	OauthGrants      []OauthGrant `json:"-" gorm:"many2many:oauth_client_grant;"`
	Roles            Roles        `gorm:"many2many:oauth_client_role;"`
	Meta             JSON         `json:"meta" gorm:"type:json"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        *time.Time   `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (oauthClient *OauthClient) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	if oauthClient.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// OauthClients is a slice of OauthClient
//
type OauthClients []OauthClient

//
// Get gets all oauthClients
//
func (oauthClients *OauthClients) Get(db *gorm.DB) error {
	if err := db.Find(&oauthClients).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one oauthClient by id
//
func (oauthClient *OauthClient) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByIDWithRoles gets one oauthClient by id with roles
//
func (oauthClient *OauthClient) GetByIDWithRoles(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Find(&oauthClient).Error; err != nil {
		return err
	}
	if err := db.Model(&oauthClient).Association("Roles").Find(&oauthClient.Roles).Error; err != nil {
		return err
	}
	return nil
}

//
// GetPartial gets one oauthClient by id
//
func (oauthClient *OauthClient) GetPartial(db *gorm.DB) error {
	if err := db.Find(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscoppedGetByID gets one oauthClient by id
//
func (oauthClient *OauthClient) UnscoppedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an OauthClient
//
func (oauthClient *OauthClient) Create(db *gorm.DB) error {
	if err := db.Create(&oauthClient).Scan(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an OauthClient
//
func (oauthClient *OauthClient) Update(db *gorm.DB) error {
	if err := db.Save(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an oauthClient
//
func (oauthClient *OauthClient) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery will retreive the oauthClient with the givene
// query
//
func (oauthClient *OauthClient) GetOneByQuery(db *gorm.DB, query string, where ...interface{}) error {
	if err := db.Where(query, where...).First(&oauthClient).Error; err != nil {
		return err
	}
	return nil
}

//
// GetWithGrants will return the oauthClient with oauthGrants associated to the client
//
func (oauthClient *OauthClient) GetWithGrants(db *gorm.DB) error {
	if err := db.Find(&oauthClient).Association("OauthGrants").Find(&oauthClient.OauthGrants).Error; err != nil {
		return err
	}
	return nil
}
