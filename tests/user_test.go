package main

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/db"
	"github.com/tespo/satya/v2/types"
)

var user = types.User{
	ID:        uuid.FromStringOrNil(""),
	AccountID: uuid.FromStringOrNil("22b5123d-9cee-4701-b15b-8c9078142666"),
	FirstName: "User",
	LastName:  "Test",
	Email:     "test@test.com",
	Phone:     "123",
	Gender:    "male",
	Height:    189.9,
	Weight:    310,
	Meta:      []byte("{\"field\":\"value\"}"),
}

func TestUserTableExists(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if !db.HasTable(&types.User{}) {
		t.Error("Database table does not exist.  Have you run migrations?")
	}
}

func TestUserCreate(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = user.Create(db)
	if err != nil {
		t.Error(err)
	}
}

func TestUserGetByID(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.User{}
	err = a.GetByID(db, user.ID)
	if err != nil {
		t.Error(err)
	}
	if a.FirstName != user.FirstName {
		t.Error("Could not find the right user")
	}
}

func TestUserGet(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.Users{}
	err = a.Get(db)
	if err != nil {
		t.Error(err)
	}
	if len(a) == 0 {
		t.Error("Could not find users")
	}
}

func TestGetUserByEmailWithRoles(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var u = types.User{}
	if err = u.GetUserByEmailWithRoles(db, "aaron@gettespo.com"); err != nil {
		t.Error(err)
	}
}

func TestUserDelete(t *testing.T) {
	db, err := db.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	var a = types.User{}
	err = a.Delete(db, user.ID)
	if err != nil {
		t.Error(err)
	}
	var softDeletedUser = types.User{}
	err = softDeletedUser.GetByID(db, user.ID)
	if err == nil || softDeletedUser.FirstName != "" {
		t.Error("Record not soft deleted")
	}
	err = softDeletedUser.UnscopedGetByID(db, user.ID)
	if err != nil {
		t.Error("Could not find soft deleted record")
	}
}

type userScopeTest struct {
	Scopes               []string
	ExpectedFields       []string
	UnExpectedFields     []string
	ExpectedShouldPass   bool
	UnexpectedShouldPass bool
}

var ScopeTestTable = []userScopeTest{
	{[]string{"user.first_name", "user.last_name"}, []string{"FirstName", "LastName"}, []string{"Email"}, true, true},
	{[]string{"user.first_name", "user.last_name"}, []string{"FirstName", "LastName"}, []string{"LastName"}, true, false},
}

func TestUserScope(t *testing.T) {
	var testUser = types.User{
		FirstName: "User",
		LastName:  "Test",
		Email:     "test@test.com",
		Phone:     "123",
		Gender:    "male",
		Height:    189.9,
		Weight:    310,
	}
	emptyUser := reflect.ValueOf(types.User{})
	for i, test := range ScopeTestTable {
		unscopedUserValue := reflect.ValueOf(testUser)
		testUser.Scope(test.Scopes)
		scopedUserValue := reflect.ValueOf(testUser)
		for _, expectedField := range test.ExpectedFields {
			if scopedUserValue.FieldByName(expectedField).Interface() != unscopedUserValue.FieldByName(expectedField).Interface() && test.ExpectedShouldPass {
				t.Errorf("Test %v failed with expected field %v returned with value %v", i, expectedField, scopedUserValue.FieldByName(expectedField))
			}
			if scopedUserValue.FieldByName(expectedField).Interface() == unscopedUserValue.FieldByName(expectedField).Interface() && !test.ExpectedShouldPass {
				t.Errorf("False positive test %v failed with expected field %v returned with value %v", i, expectedField, scopedUserValue.FieldByName(expectedField))
			}
		}
		for _, unexpectedField := range test.UnExpectedFields {
			if scopedUserValue.FieldByName(unexpectedField).Interface() != emptyUser.FieldByName(unexpectedField).Interface() && test.UnexpectedShouldPass {
				t.Errorf("Test %v failed with unexpected field %v returned with value %v", i, unexpectedField, scopedUserValue.FieldByName(unexpectedField))
			}

			if scopedUserValue.FieldByName(unexpectedField).Interface() == emptyUser.FieldByName(unexpectedField).Interface() && !test.UnexpectedShouldPass {
				t.Errorf("Test %v failed with unexpected field %v returned with value %v", i, unexpectedField, scopedUserValue.FieldByName(unexpectedField))
			}
		}
	}
}
