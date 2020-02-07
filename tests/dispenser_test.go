package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var dispenser = types.Dispenser{

	ID:      uuid.FromStringOrNil("818eb989-5045-4f3f-b898-80559876f22d"),
	Serial:  "dispenser-11218NEWTC",
	Name:    "Tespo Connect",
	Network: "UNIMATRIX2",
	Meta:    nil,
}

func TestDispenserTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Dispenser{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestDispenserCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = dispenser.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestDispenserGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var d = types.Dispenser{}
	err = d.GetByID(db, dispenser.ID)
	if err != nil {
		t.Error(err)
	}
	if d.Name != dispenser.Name {
		t.Error("Could not find the right dispenser")
	}
}

func TestDispenserDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Dispenser{}
	err = a.Delete(db, dispenser.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedDispenser = types.Dispenser{}
	err = softDeletedDispenser.GetByID(db, dispenser.ID)
	if err == nil || softDeletedDispenser.Name != "" {
		t.Error("Record not soft deleted")
	}
	err = softDeletedDispenser.UnscopedGetByID(db, dispenser.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
