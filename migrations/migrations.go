package migrations

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

//
// Migrate runs all migrations
//
func Migrate() {
	db, err := db.Open()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	// Tables
	// create primary tables such as Account, Dispensers, et cetera
	migratePrimaryTables(db)
	// create the tables that use foreign keys
	migrateForeignTables(db)
	// add the foreign keys to foreign tables
	addForeignKeys(db)
}

//
// Nuke will drop all tables
//
func Nuke() {
	db, err := db.Open()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	// remove all the foreign key constraints before dropping
	removeForeignKeys(db)
	// drop it like it's hot
	dropAllTables(db)
}

func dropAllTables(db *gorm.DB) {
	// Drop tables

	if err := db.DropTable(&types.Usage{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Reminder{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Regimen{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Connection{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Insertion{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable("oauth_client_role").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable("role_user").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable("permission_role").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable("oauth_client_grant").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Role{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Barcode{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Permission{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.User{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Pod{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Dispenser{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Account{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthAuthorizationCode{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthClientUserRefreshToken{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthScopeRequest{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthAccessToken{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthGrant{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.OauthClient{}).Error; err != nil {
		fmt.Print(err)
	}
}

func removeForeignKeys(db *gorm.DB) {
	// Drop Constraints
	if err := db.Model(&types.Barcode{}).RemoveForeignKey("pod_id", "pods(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Reminders{}).RemoveForeignKey("user_id", "users(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Reminders{}).RemoveForeignKey("regimen_id", "regimens(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("permission_role").RemoveForeignKey("permission_id", "permissions(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("permission_role").RemoveForeignKey("role_id", "roles(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Connections{}).RemoveForeignKey("account_id", "accounts(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Connections{}).RemoveForeignKey("dispenser_id", "dispensers(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Insertions{}).RemoveForeignKey("dispenser_id", "dispensers(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Insertions{}).RemoveForeignKey("regimen_id", "regimens(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Insertions{}).RemoveForeignKey("barcode_id", "barcodes(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).RemoveForeignKey("account_id", "accounts(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).RemoveForeignKey("user_id", "users(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).RemoveForeignKey("pod_id", "pods(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).RemoveForeignKey("dispenser_id", "dispensers(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).RemoveForeignKey("regimen_id", "regimens(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("role_user").RemoveForeignKey("user_id", "users(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("role_user").RemoveForeignKey("role_id", "roles(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.User{}).RemoveForeignKey("account_id", "accounts(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("oauth_client_grant").RemoveForeignKey("oauth_client_id", "oauth_clients(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("oauth_client_grant").RemoveForeignKey("oauth_grant_id", "oauth_grants(id)").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.DropTable(&types.Invitation{}).Error; err != nil {
		fmt.Print(err)
	}
}

func addForeignKeys(db *gorm.DB) {
	// Foreign Keys
	if err := db.Model(&types.Users{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Barcodes{}).AddForeignKey("pod_id", "pods(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Connections{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Connections{}).AddForeignKey("dispenser_id", "dispensers(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).AddForeignKey("dispenser_id", "dispensers(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).AddForeignKey("regimen_id", "regimens(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Usages{}).AddForeignKey("barcode_id", "barcodes(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Regimens{}).AddForeignKey("pod_id", "pods(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Reminders{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Reminders{}).AddForeignKey("regimen_id", "regimens(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Insertions{}).AddForeignKey("dispenser_id", "dispensers(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Model(&types.Insertions{}).AddForeignKey("regimen_id", "regimens(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("permission_role").AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("permission_role").AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("role_user").AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("role_user").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("oauth_client_grant").AddForeignKey("oauth_client_id", "oauth_clients(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
	if err := db.Table("oauth_client_grant").AddForeignKey("oauth_grant_id", "oauth_grants(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		fmt.Print(err)
	}
}

func migrateForeignTables(db *gorm.DB) {
	if err := db.AutoMigrate(&types.Barcode{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Regimen{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Connection{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Insertion{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Usage{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Reminder{}).Error; err != nil {
		fmt.Print(err)
	}
}

func migratePrimaryTables(db *gorm.DB) {
	if err := db.AutoMigrate(&types.Invitation{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Account{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Role{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Permission{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.User{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Dispenser{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.Pod{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthClient{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthAccessToken{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthAuthorizationCode{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthClientUserRefreshToken{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthGrant{}).Error; err != nil {
		fmt.Print(err)
	}
	if err := db.AutoMigrate(&types.OauthScopeRequest{}).Error; err != nil {
		fmt.Print(err)
	}
}
