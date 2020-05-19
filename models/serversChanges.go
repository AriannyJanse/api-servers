package models

import (
	u "api-server/utils"
	"fmt"
)

type ServersChanges struct {
	Grade         string `json:"ssl_grade"`
	PreviousGrade string `json:"previous_ssl_grade"`
	IsChanged     bool   `json:"servers_changed"`
}

func (serverChanges *ServersChanges) validateWithIDServer(idServer uint) (map[string]interface{}, bool) {

	if idServer < 0 {
		return u.Message(false, "ID should be on the payload"), false
	}
	if serverChanges.Grade == "" {
		return u.Message(false, "Grade should be on the payload"), false
	}
	if serverChanges.PreviousGrade == "" {
		return u.Message(false, "Previous Grade should be on the payload"), false
	}

	return u.Message(true, "The serverChanges data is valid"), true
}

func (serverChanges *ServersChanges) CreateServersChangesByID(id uint) map[string]interface{} {
	if resp, ok := serverChanges.validateWithIDServer(id); !ok {
		return resp
	}

	query := fmt.Sprintf(`INSERT INTO servers_changes (changes_id_server, ssl_grade, previous_ssl_grade)
						VALUES ('%d', '%s', '%s')`, id, serverChanges.Grade, serverChanges.PreviousGrade)
	fmt.Println(query)
	_, err := GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Error trying to exec the serverChanges query")
	}
	return u.Message(true, "SeverChanges was created")

}
