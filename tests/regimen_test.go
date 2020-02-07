package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var userID = uuid.FromStringOrNil("ef837bfd-aae4-4495-8aec-5d70f3aa0ed3")
var podID = uuid.FromStringOrNil("938b0ff3-a272-49cc-b1c8-b09ccfd07792")
var regimen = types.Regimen{
	ID:        uuid.FromStringOrNil("afdb81ab-5da2-4cf6-87ca-e626e6853e27"),
	AccountID: uuid.FromStringOrNil("22b5123d-9cee-4701-b15b-8c9078142666"),
	UserID:    &userID,
	PodID:     &podID,
}

func TestRegimenTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Regimen{}) {
		t.Error("Database table does not exist. Have you run migrations?")
	}
}

func TestRegimenCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = regimen.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestRegimenGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var r = types.Regimen{}
	err = r.GetByID(db, regimen.ID)
	if err != nil {
		t.Error(err)
	}
	if r.ID != regimen.ID {
		t.Error("Could not find the right regimen")
	}
}

func TestRegimenGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var r = types.Regimens{}
	err = r.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(r) == 0 {
		t.Error("Could not find regimens")
	}
}

func TestRegimenDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var r = types.Regimen{}
	err = r.Delete(db, regimen.ID)
	if err != nil {
		t.Error(err)
	}

	var softDeletedRegimen = types.Regimen{}
	err = softDeletedRegimen.GetByID(db, regimen.ID)
	if err == nil || softDeletedRegimen.DeletedAt != nil {
		t.Error("Record not soft deleted")
	}
	err = softDeletedRegimen.UnscopedGetByID(db, regimen.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
