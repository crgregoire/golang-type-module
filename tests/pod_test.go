package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var pod = types.Pod{

	ID:        uuid.FromStringOrNil("a28a64ce-4f6f-4461-b51e-cd4c1d5522ec"),
	Name:      "Vision Support",
	Slug:      "vision",
	Color:     "#E2C52C",
	Cells:     20,
	LabelTall: "https://gettespo.com/visionsupportlabeltall",
	LabelWide: "https://gettespo.com/visionsupportlabelwide",
	Meta:      nil,
}

func TestPodTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Pod{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestPodCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = pod.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestPodGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var d = types.Pod{}
	err = d.GetByID(db, pod.ID)
	if err != nil {
		t.Error(err)
	}
	if d.Name != pod.Name {
		t.Error("Could not find the right pod")
	}
}

func TestPodDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var p = types.Pod{}
	err = p.Delete(db, pod.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedPod = types.Pod{}
	err = softDeletedPod.GetByID(db, pod.ID)
	if err == nil || softDeletedPod.Name != "" {
		t.Error("Record not soft deleted")
	}
	err = softDeletedPod.UnscopedGetByID(db, pod.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
