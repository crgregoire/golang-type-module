package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var account = types.Account{

	ID:   uuid.FromStringOrNil("b517c1e1-bdfc-4998-8e47-f268fe8eff60"),
	Name: "AYAYRON",
}

func TestAccountTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Account{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestAccountCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = account.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Account{}
	err = a.GetByID(db, account.ID)
	if err != nil {
		t.Error(err)
	}
	if a.Name != account.Name {
		t.Error("Could not find the right account")
	}
}

func TestAccountGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Accounts{}
	err = a.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(a) == 0 {
		t.Error("Could not find accounts")
	}
}

func TestAccountDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Account{}
	err = a.Delete(db, account.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedAccount = types.Account{}
	err = softDeletedAccount.GetByID(db, account.ID)
	if err == nil || softDeletedAccount.Name != "" {
		t.Error("Record not soft deleted")
	}
	err = softDeletedAccount.UnscopedGetByID(db, account.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
