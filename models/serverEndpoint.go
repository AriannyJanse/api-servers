package models

import (
	u "api-server/utils"
	"database/sql"
	"fmt"
)

type ServerEndpoint struct {
	Address string `json:"address"`
	Grade   string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner   string `json:"owner"`
}

func (serverEndpoint *ServerEndpoint) validate() (map[string]interface{}, bool) {

	if serverEndpoint.Address == "" {
		return u.Message(false, "Address should be on the payload"), false
	}

	if serverEndpoint.Country == "" {
		return u.Message(false, "Country should be on the payload"), false
	}

	if serverEndpoint.Owner == "" {
		return u.Message(false, "Owner should be on the payload"), false
	}

	return u.Message(true, "The serverEndpoint data is valid"), true
}

func (serverEndpoint *ServerEndpoint) CreateServerEndpointByServerID(id uint) map[string]interface{} {
	if resp, ok := serverEndpoint.validate(); !ok {
		return resp
	}
	query := fmt.Sprintf(`INSERT INTO server_endpoints (endpoint_id_server, endpoint_address, 
						endpoint_ssl_grade, endpoint_country, endpoint_owner)
						VALUES (%d, '%s', '%s', '%s', '%s')`, id, serverEndpoint.Address, serverEndpoint.Grade,
		serverEndpoint.Country, serverEndpoint.Owner)
	fmt.Println(query)
	if _, err := GetDB().Exec(query); err != nil {
		fmt.Println(err)
		return u.Message(false, "Error trying to exec serverEndpoint query")
	}
	return u.Message(true, "ServerEndpoint was created")
}

func GetServerEndpointsByHostname(hostname string) []*ServerEndpoint {
	serverEndpoints := make([]*ServerEndpoint, 0)
	query := fmt.Sprintf(`SELECT endpoint_address, endpoint_ssl_grade, endpoint_country, endpoint_owner
						FROM server_endpoints INNER JOIN servers 
						ON server_endpoints.endpoint_id_server = servers.id_server
						WHERE servers.hostname = '%s';`, hostname)
	fmt.Println(query)
	rows, err := GetDB().Query(query)
	if err != nil {
		fmt.Println("Cannot found serverEndpoint with hostname: ", hostname, "Error: ", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var serverEndpoint ServerEndpoint
		if err := rows.Scan(&serverEndpoint.Address, &serverEndpoint.Grade, &serverEndpoint.Country, &serverEndpoint.Owner); err != nil {
			fmt.Println(err)
		}
		serverEndpoints = append(serverEndpoints, &serverEndpoint)
	}
	return serverEndpoints
}

func (serverEndpoint *ServerEndpoint) UpdateServerEndpointByID(id uint) map[string]interface{} {
	query := fmt.Sprintf(`UPDATE server_endpoints SET endpoint_address = '%s', endpoint_ssl_grade = '%s', 
						endpoint_country = '%s', endpoint_owner = '%s', updatedat = NOW() WHERE id_server_endpoint = '%d'`,
		serverEndpoint.Address, serverEndpoint.Grade, serverEndpoint.Country, serverEndpoint.Owner, id)
	fmt.Println(query)
	if _, err := GetDB().Exec(query); err != nil {
		fmt.Println(err)
		return u.Message(false, "Error trying to exec serverEndpoint update query")
	}
	return u.Message(true, "ServerEndpoint was updated")
}

func (serverEndpoint *ServerEndpoint) GetIDByAddress() uint {
	var id uint

	query := fmt.Sprintf(`SELECT id_server_endpoint FROM server_endpoints WHERE endpoint_address = '%s';`, serverEndpoint.Address)
	fmt.Println(query)
	row := GetDB().QueryRow(query)

	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Cannot found ID with address: ", serverEndpoint.Address)
			return 0
		}
		fmt.Println(err)
	}
	return id
}
