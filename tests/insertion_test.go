package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var barcodeID = uuid.FromStringOrNil("62012ee1-f3c2-4a4b-a0ab-8686e3d173e4")

var insertion = types.Insertion{
	ID:          uuid.FromStringOrNil("e83b7066-34d4-43f1-a384-4b10d009e9b5"),
	DispenserID: uuid.FromStringOrNil("9f717337-dba7-415c-9daf-c607df526d14"),
	RegimenID:   uuid.FromStringOrNil("b3d8bf48-3a8f-4ca5-b1e5-452a61be493a"),
	BarcodeID:   &barcodeID,
	Flags:       5,
	Servings:    31,
	LabelTall:   "https://gettespo.com/menscompletetall",
	LabelWide:   "https://gettespo.com/menscompletewide",
	Meta:        nil,
}

func TestInsertionTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Insertion{}) {
		t.Error("Database table does not exist. Have you run migrations?")
	}
}

func TestInsertionCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = insertion.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertionGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var i = types.Insertion{}
	err = i.GetByID(db, insertion.ID)
	if err != nil {
		t.Error(err)
	}
	if i.ID != insertion.ID {
		t.Error("Could not find the right insertion")
	}
}

func TestInsertionGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var i = types.Insertions{}
	err = i.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(i) == 0 {
		t.Error("Could not find insertions")
	}
}

func TestInsertionDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var i = types.Insertion{}
	err = i.Delete(db, insertion.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeleteInsertion = types.Insertion{}
	err = softDeleteInsertion.GetByID(db, insertion.ID)
	if err == nil || softDeleteInsertion.DeletedAt != nil {
		t.Error("Record not soft deleted")
	}
	err = softDeleteInsertion.UnscopedGetByID(db, insertion.ID)
	if err != nil {
		t.Error("Count not find soft deleted record")
	}
}
