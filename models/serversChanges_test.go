package models

import (
	u "api-server/utils"
	"reflect"
	"testing"
)

func TestValidateWithIDServer(t *testing.T) {
	server := &ServersChanges{}

	resp, respOk := server.validateWithIDServer(1)
	expected := u.Message(true, "The serverChanges data is valid")
	expectedOk := true

	result := reflect.DeepEqual(resp, expected)
	if !result || (respOk != expectedOk) {
		t.Errorf("Unexpected result: GOT %v, %v WANT %v, %v", resp, respOk, expected, expectedOk)
	}
}

func TestCreateServersChangesByID(t *testing.T) {
	//Initializing db
	Init()

	serverChanges := &ServersChanges{
		Grade:         "B",
		PreviousGrade: "C",
		IsChanged:     true,
	}

	resp := serverChanges.CreateServersChangesByID(559430992791568385)
	expected := u.Message(true, "SeverChanges was created")
	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}
