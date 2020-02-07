package main

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var connection = types.Connection{
	ID:          uuid.FromStringOrNil("157f7fad-ea09-4c68-88e2-e015ea613bce"),
	AccountID:   uuid.FromStringOrNil("d794ce08-0793-440c-ab0f-518d14e11377"),
	DispenserID: uuid.FromStringOrNil("9f717337-dba7-415c-9daf-c607df526d14"),
	Meta:        nil,
	ConnectedAt: time.Now(),
}

func TestConnectionTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Connection{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestConnectionCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = connection.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestConnectionGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var c = types.Connection{}
	err = c.GetByID(db, connection.ID)
	if err != nil {
		t.Error(err)
	}
	if c.ID != connection.ID {
		t.Error("Could not find the right connection")
	}
}

func TestConnectionDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var c = types.Connection{}
	err = c.Delete(db, connection.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedConnection = types.Connection{}
	err = softDeletedConnection.GetByID(db, connection.ID)
	if err == nil || softDeletedConnection.ID != uuid.Nil {
		t.Error("Record not soft deleted")
	}
	err = softDeletedConnection.UnscopedGetByID(db, connection.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
