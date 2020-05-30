package models

import (
	u "api-server/utils"
	"reflect"
	"testing"
)

func TestValidateServerEndpoint(t *testing.T) {
	serverEndpoint := &ServerEndpoint{
		Address: "1.1.1.1",
		Country: "US",
		Owner:   "Example",
	}

	resp, respOk := serverEndpoint.validate()
	expected := u.Message(true, "The serverEndpoint data is valid")
	expectedOk := true

	result := reflect.DeepEqual(resp, expected)
	if !result || (respOk != expectedOk) {
		t.Errorf("Unexpected result: GOT %v, %v WANT %v, %v", resp, respOk, expected, expectedOk)
	}
}

func TestCreateServerEndpointByServerID(t *testing.T) {
	//Initializing db
	Init()

	serverEndpoint := &ServerEndpoint{
		Address: "1.1.1.1",
		Grade:   "B",
		Country: "US",
		Owner:   "Example",
	}

	resp := serverEndpoint.CreateServerEndpointByServerID(559430992791568385)
	expected := u.Message(true, "ServerEndpoint was created")

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}

}

func TestGetServerEndpointsByHostname(t *testing.T) {
	//Initializing db
	Init()

	resp := GetServerEndpointsByHostname("testingendpoints.com")
	expected := []*ServerEndpoint{
		{
			Address: "1.2.3.4",
			Grade:   "B",
			Country: "US",
			Owner:   "EXAMPLE",
		},
		{
			Address: "1.1.1.1",
			Grade:   "B",
			Country: "US",
			Owner:   "Example",
		},
	}

	for i, value := range resp {
		if (value.Address != expected[i].Address) || (value.Grade != expected[i].Grade) ||
			(value.Country != expected[i].Country) || (value.Owner != expected[i].Owner) {
			t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
		}
	}
}

func TestUpdateServerEndpointByID(t *testing.T) {
	//Initializing db
	Init()

	serverEndpoint := &ServerEndpoint{
		Address: "1.2.3.4",
		Grade:   "B",
		Country: "US",
		Owner:   "EXAMPLE",
	}

	resp := serverEndpoint.UpdateServerEndpointByID(559434526969266177)
	expected := u.Message(true, "ServerEndpoint was updated")

	result := reflect.DeepEqual(resp, expected)
	if !result {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}

func TestGetIDByAddress(t *testing.T) {
	//Initializing db
	Init()

	serverEndpoint := &ServerEndpoint{
		Address: "1.2.3.4",
		Grade:   "C",
		Country: "US",
		Owner:   "EXAMPLE",
	}

	resp := serverEndpoint.GetIDByAddress()
	expected := uint(559434526969266177)
	if resp != expected {
		t.Errorf("Unexpected result: GOT %v WANT %v", resp, expected)
	}
}
