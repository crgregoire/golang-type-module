package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var barcode = types.Barcode{

	ID:        uuid.FromStringOrNil("71113cbe-f3d9-4708-b53d-4e17ab37613c"),
	PodID:     uuid.FromStringOrNil("16ea71db-2adb-45fe-a3fe-9e9ad6dcabd3"),
	Code:      "104080920472",
	LabelTall: "https://gettespo.com/menslabeltall",
	LabelWide: "https://gettespo.com/menslabelwide",
}

func TestBarcodeTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Barcode{}) {
		t.Error("Database table does not exist. Have you run migrations?")
	}
}

func TestBarcodeCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = barcode.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestBarcodeGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var b = types.Barcode{}
	err = b.GetByID(db, barcode.ID)
	if err != nil {
		t.Error(err)
	}
	if b.ID != barcode.ID {
		t.Error("Could not find the right barcode")
	}
}

func TestBarcodeGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var b = types.Barcodes{}
	err = b.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(b) == 0 {
		t.Error("Could not find barcodes")
	}
}

func TestBarcodeDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var b = types.Barcode{}
	err = b.Delete(db, barcode.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedBarcode = types.Barcode{}
	err = softDeletedBarcode.GetByID(db, barcode.ID)
	if err == nil || softDeletedBarcode.DeletedAt != nil {
		t.Error("Record not soft deleted")
	}
	err = softDeletedBarcode.UnscopedGetByID(db, barcode.ID)
	if err != nil {
		t.Error("Could not find soft deleted barcode")
	}
}
