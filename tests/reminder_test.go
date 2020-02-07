package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var reminder = types.Reminder{
	ID:        uuid.FromStringOrNil("10608c7e-5e63-49f9-913b-662271ed5fa6"),
	UserID:    uuid.FromStringOrNil("dc59e7aa-7fd6-4e90-8ee4-a3af780fd32d"),
	RegimenID: uuid.FromStringOrNil("395a37a3-a63b-49c4-94b4-04c29d47c64c"),
	Minute:    195,
}

func TestReminderTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.Reminder{}) {
		t.Error("Database table does not exist. Have you run migrations?")
	}
}

func TestReminderCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = reminder.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestReminderGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Reminder{}
	err = u.GetByID(db, reminder.ID)
	if err != nil {
		t.Error(err)
	}
	if u.ID != reminder.ID {
		t.Error("Could not find the right dispenser")
	}
}

func TestReminderGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Reminders{}
	err = u.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(u) == 0 {
		t.Error("Could not find assignments")
	}
}

func TestReminderDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.Reminder{}
	err = u.Delete(db, reminder.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedReminder = types.Reminder{}
	err = softDeletedReminder.GetByID(db, reminder.ID)
	if err == nil || softDeletedReminder.DeletedAt != nil {
		t.Error("Record not soft deleted")
	}
	err = softDeletedReminder.UnscopedGetByID(db, reminder.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}
