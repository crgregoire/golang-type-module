package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var role = types.Role{
	ID:   uuid.FromStringOrNil(""),
	Name: "",
	Meta: []byte("{\"field\":\"value\"}"),
}

func TestRoleTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Role{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestRoleCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = role.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestRoleGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Role{}
	err = a.GetByID(db, role.ID)
	if err != nil {
		t.Error(err)
	}
	if a.Name != role.Name {
		t.Error("Could not find the right role")
	}
}

func TestRoleGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Roles{}
	err = a.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(a) == 0 {
		t.Error("Could not find roles")
	}
}

func TestRoleDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Role{}
	err = a.Delete(db, role.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedRole = types.Role{}
	err = softDeletedRole.GetByID(db, role.ID)
	if err == nil || softDeletedRole.Name != "" {
		t.Error("Record not soft deleted")
	}
	err = softDeletedRole.UnscopedGetByID(db, role.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
