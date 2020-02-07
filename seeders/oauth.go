package seeders

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/types"
)

var oauthClients = []types.OauthClient{
	{
		ID:               uuid.FromStringOrNil("a88c5c53-3209-42f0-b29e-c6d4bfc686ac"),
		Name:             "user scoped oauth client",
		ClientID:         "e4f30140a8d647168796e4e0ee6b61ca",
		ClientSecret:     "4f58fd37769c439595c2263ef94d8af6",
		RedirectURI:      "http://localhost:6969/redirect",
		ScopePermissions: []byte("[\"user.info\",\"user.update\",\"user.regimens\",\"user.regimen\",\"user.reminder\",\"user.create.reminder\",\"user.update.reminder\",\"user.delete.reminder\"]"),
		ScopedFields:     []byte("[\"user.*\"]"),
	},
	{
		ID:               uuid.FromStringOrNil("3708a993-a42a-4f72-b48e-fef0a1edb7a6"),
		Name:             "all scopes oauth client",
		ClientID:         "24b251eb165e409885c635344a3c6bc6",
		ClientSecret:     "5bc4f24beac5494d8f4f56a98c8f46b4",
		RedirectURI:      "http://localhost:6969/redirect",
		ScopePermissions: []byte("[\"*\"]"),
		ScopedFields:     []byte("[\"*\"]"),
	},
}

var oauthGrants = []types.OauthGrant{
	{
		ID:   uuid.NewV4(),
		Type: "authorization_code",
	},
	{
		ID:   uuid.NewV4(),
		Type: "password",
	},
	{
		ID:   uuid.NewV4(),
		Type: "client_credentials",
	},
	{
		ID:   uuid.NewV4(),
		Type: "refresh_token",
	},
	{
		ID:   uuid.NewV4(),
		Type: "device_code",
	},
}

var oauthClientWithRole = map[uuid.UUID][]uuid.UUID{
	uuid.FromStringOrNil("3708a993-a42a-4f72-b48e-fef0a1edb7a6"): {
		uuid.FromStringOrNil("b22510b1-b501-4c68-802a-e0ebce2b8307"),
	},
}

//
// SeedOauthClients seeds oauth clients
//
func seedOauth(db *gorm.DB) error {
	if !db.HasTable(&types.OauthClient{}) {
		if err := db.AutoMigrate(&types.OauthClient{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&types.OauthAccessToken{}) {
		if err := db.AutoMigrate(&types.OauthAccessToken{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&types.OauthAuthorizationCode{}) {
		if err := db.AutoMigrate(&types.OauthAuthorizationCode{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&types.OauthClientUserRefreshToken{}) {
		if err := db.AutoMigrate(&types.OauthClientUserRefreshToken{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&types.OauthGrant{}) {
		if err := db.AutoMigrate(&types.OauthGrant{}).Error; err != nil {
			return err
		}
	}
	if !db.HasTable(&types.OauthScopeRequest{}) {
		if err := db.AutoMigrate(&types.OauthScopeRequest{}).Error; err != nil {
			return err
		}
	}
	for _, oauthGrant := range oauthGrants {
		if err := oauthGrant.Create(db); err != nil {
			return err
		}
	}
	for _, oauthClient := range oauthClients {
		if err := oauthClient.Create(db); err != nil {
			return err
		}
		if err := db.Model(oauthClient).Association("OauthGrants").Append(oauthGrants).Error; err != nil {
			return err
		}
	}

	for clientID, roles := range oauthClientWithRole {
		for _, roleID := range roles {
			if err := db.Table("oauth_client_role").Create(struct {
				OauthClientID uuid.UUID
				RoleID        uuid.UUID
			}{
				OauthClientID: clientID,
				RoleID:        roleID,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
