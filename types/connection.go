package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// Connection connects accounts and their dispenser Connection
//
type Connection struct {
	ID             uuid.UUID   `json:"id" gorm:"type:char(36);primary_key"`
	AccountID      uuid.UUID   `json:"account_id" gorm:"type:char(36);foreign_key"`
	DispenserID    uuid.UUID   `json:"dispenser_id" gorm:"type:char(36);foreign_key" scope:"connection.dispenser_id"`
	Dispensers     []Dispenser `json:"dispensers"`
	Meta           JSON        `json:"meta" gorm:"type:json" scope:"connection.meta"`
	ConnectedAt    time.Time   `json:"connected_at"`
	DisconnectedAt *time.Time  `json:"disconnected_at"`
	CreatedAt      time.Time   `json:"created_at" scope:"connection.created_at"`
	UpdatedAt      time.Time   `json:"updated_at" scope:"connection.updated_at"`
	DeletedAt      *time.Time  `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (connection *Connection) BeforeCreate(scope *gorm.Scope) error {
	noTime := time.Time{}
	if connection.DisconnectedAt != nil {
		if connection.DisconnectedAt.String() == noTime.String() {
			scope.SetColumn("DisconnectedAt", nil)
		}
	}
	uuid := uuid.NewV4()
	if connection.ID.String() == "00000000-0000-0000-0000-000000000000" {
		scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Connections is a slice of Connection
//
type Connections []Connection

//
// Get gets all connections
//
func (connections *Connections) Get(db *gorm.DB) error {
	if err := db.Find(&connections).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one connection by id
//
func (connection *Connection) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).First(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets Connection by query
//
func (connection *Connection) GetOneByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Where(where, query...).First(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one connection by id
//
func (connection *Connection) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Unscoped().Where("id = ?", id).Find(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Connection
//
func (connection *Connection) Create(db *gorm.DB) error {
	if err := db.Create(&connection).Scan(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Connection
//
func (connection *Connection) Update(db *gorm.DB) error {
	if err := db.Save(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an connection
//
func (connection *Connection) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&connection).Error; err != nil {
		return err
	}
	return nil
}

//
// GetAccountDispensers gets dispensers associated with an account
//
func (connections *Connections) GetAccountDispensers(db *gorm.DB, accountID uuid.UUID) (Dispensers, error) {
	if err := db.Where("connections.account_id = ?", accountID).Find(&connections).Error; err != nil {
		return nil, err
	}
	dispensers := Dispensers{}
	for _, connection := range *connections {
		if err := db.Model(&connection).Related(&connection.Dispensers).Error; err != nil {
			return nil, err
		}
		dispensers = append(dispensers, connection.Dispensers...)
	}
	return dispensers, nil
}

//
// GetAccountDispenserByID hard deletes a connection
//
func (connection *Connection) GetAccountDispenserByID(db *gorm.DB, accountID uuid.UUID, dispenserID uuid.UUID) error {
	if err := db.Find(&connection, "account_id = ? and dispenser_id = ? ", accountID, dispenserID).Related(&connection.Dispensers).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (connection *Connection) Scope(scopes []string) {
	*connection = scoping.FilterByScopes(scopes, *connection).(Connection)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (connections *Connections) Scope(scopes []string) {
	connectionSlice := *connections
	for i, connection := range connectionSlice {
		connectionSlice[i] = scoping.FilterByScopes(scopes, connection).(Connection)
	}
	*connections = connectionSlice
}
