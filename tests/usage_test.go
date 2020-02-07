package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var usage = types.Usage{
	ID:          uuid.FromStringOrNil("10608c7e-5e63-49f9-913b-662271ed5fa6"),
	DispenserID: uuid.FromStringOrNil("142201c2-0c5f-4650-8c99-fc233412e030"),
	RegimenID:   uuid.FromStringOrNil("395a37a3-a63b-49c4-94b4-04c29d47c64c"),
	Meta:        nil,
}

func TestUsageTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Usage{}) {
		t.Error("Database table does not exist. Have you run migrations?")
	}
}

func TestUsageCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = usage.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestUsageGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Usage{}
	err = u.GetByID(db, usage.ID)
	if err != nil {
		t.Error(err)
	}
	if u.ID != usage.ID {
		t.Error("Could not find the right dispenser")
	}
}

func TestUsageGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Usages{}
	err = u.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(u) == 0 {
		t.Error("Could not find assignments")
	}
}

func TestUsageDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Usage{}
	err = u.Delete(db, usage.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedUsage = types.Usage{}
	err = softDeletedUsage.GetByID(db, usage.ID)
	if err == nil || softDeletedUsage.DeletedAt != nil {
		t.Error("Record not soft deleted")
	}
	err = softDeletedUsage.UnscopedGetByID(db, usage.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
