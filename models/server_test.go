package models

import (
	u "api-server/utils"
	"reflect"
	"testing"
)

func TestValidateServer(t *testing.T) {
	server := &Server{
		Logo:  "test logo",
		Title: "test title",
	}

	resp, respOk := server.validate()
	expected := u.Message(true, "The server data is valid")
	expectedOk := true

	result := reflect.DeepEqual(resp, expected)
	if !result || (respOk != expectedOk) {
		t.Errorf("Unexpected result: GOT %v, %v WANT %v, %v", resp, respOk, expected, expectedOk)
	}
}

func TestCreateServerByHostname(t *testing.T) {
	//Initializing db
	Init()

	domain := "testingdomain.com"

	server := &Server{
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	resp := server.CreateServerByHostname(domain)
	expected := u.Message(true, "Server was created")

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}

}

func TestCreateServerByHostnameAlredyExist(t *testing.T) {
	//Initializing db
	Init()

	server := &Server{
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	resp := server.CreateServerByHostname("testingdomain.com")
	expected := u.Message(false, "The server alredy exist")

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}

func TestGetServersHostnames(t *testing.T) {
	//Initializing db
	Init()

	resp := GetServersHostnames()
	expected := []string{
		"facebook.com",
		"github.com",
		"google.com",
		"testingdomain.com",
		"testingendpoints.com",
		"truora.com",
		"twitter.com",
		"youtube.com",
	}

	for i, value := range resp {
		if value != expected[i] {
			t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
		}
	}
}

func TestGetServerByDomain(t *testing.T) {
	//Initializing db
	Init()

	resp := GetServerByDomain("testingdomain.com")
	expected := &Server{
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	if (resp.Logo != expected.Logo) || (resp.Title != expected.Title) || (resp.IsDown != expected.IsDown) {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}

}

func TestUpdateServerByHostname(t *testing.T) {
	//Initializing db
	Init()

	serverUpdated := &Server{
		Logo:   "test logo updated",
		Title:  "test title updated",
		IsDown: false,
	}

	resp := serverUpdated.UpdateServerByHostname("testingdomain.com", serverUpdated)
	expected := u.Message(true, "The server was updated")

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}

func TestGetSeverIDAndCheckIfExistByHostname(t *testing.T) {
	//Initializing db
	Init()

	respID, respBool := GetSeverIDAndCheckIfExistByHostname("google.com")
	expectedID := uint(558322320757424129)
	expectedBool := true

	if (respID != expectedID) || (respBool != expectedBool) {
		t.Errorf("Unexpected result: GOT %v, %v WANT %v, %v", respID, respBool, expectedID, expectedBool)
	}

}

func TestServerIsChange(t *testing.T) {
	server := &Server{
		Servers: []*ServerEndpoint{
			{
				Address: "1.1.1.1",
				Grade:   "B",
				Country: "US",
				Owner:   "Example",
			},
			{
				Address: "1.1.1.1",
				Grade:   "B",
				Country: "US",
				Owner:   "Example",
			},
		},
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	serverChanged := &u.ApiServer{
		Servers: []*u.ApiEndpoint{
			{
				Address: "1.1.1.1",
				Grade:   "A",
				Country: "US",
				Owner:   "Example",
			},
			{
				Address: "1.1.1.1",
				Grade:   "B",
				Country: "US",
				Owner:   "Example",
			},
		},
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	resp := server.ServerIsChange(serverChanged)
	expected := map[string]interface{}{"grade": "B", "message": "grade changed", "previousGrade": "B", "status": true}

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}

func TestGetGrade(t *testing.T) {
	grades := []string{
		"Z",
		"X",
		"B",
		"A",
	}

	resp := getGrade(grades)
	expected := "Z"

	if resp != expected {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}

func TestEncodingApiToServer(t *testing.T) {
	serverToEncode := &u.ApiServer{
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	serverEncoded := EncodingApiToServer(serverToEncode)
	expected := &Server{
		Logo:   "test logo",
		Title:  "test title",
		IsDown: false,
	}

	if (serverEncoded.Logo != expected.Logo) || (serverEncoded.Title != expected.Title) || (serverEncoded.IsDown != expected.IsDown) {
		t.Errorf("Unexpected result: GOT %v WANT %v", serverEncoded, expected)
	}
}
