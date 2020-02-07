package seeders

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/tespo/satya/v2/db"
)

//
// SeedDatabase is used to seed the Database for
// local development
//
func SeedDatabase() {
	db, err := db.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	seedPrimaryTables(db)
	seedForeignTables(db)
}

//
// seedForeignTables seeds the tables which require foreign keys for local development
//
func seedForeignTables(db *gorm.DB) {
	if err := seedRegimens(db); err != nil {
		log.Println(err)
	}
	if err := seedReminders(db); err != nil {
		log.Println(err)
	}
	if err := seedBarcodes(db); err != nil {
		log.Println(err)
	}
	if err := seedConnections(db); err != nil {
		log.Println(err)
	}
	if err := seedInsertions(db); err != nil {
		log.Println(err)
	}
	if err := seedUsages(db); err != nil {
		log.Println(err)
	}
}

//
// seedPrimaryTables seeds the tables from which foreign keys are made for local development
//
func seedPrimaryTables(db *gorm.DB) {
	if err := seedPermissions(db); err != nil {
		log.Println(err)
	}
	if err := seedRoles(db); err != nil {
		log.Println(err)
	}
	if err := seedAccounts(db); err != nil {
		log.Println(err)
	}
	if err := seedDispensers(db); err != nil {
		log.Println(err)
	}
	if err := seedUsers(db); err != nil {
		log.Println(err)
	}
	if err := seedPods(db); err != nil {
		log.Println(err)
	}
	if err := seedOauth(db); err != nil {
		log.Println(err)
	}
}
