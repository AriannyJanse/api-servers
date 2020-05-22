package models

import (
	u "api-server/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
)

type Server struct {
	Servers      []*ServerEndpoint
	ServerChange *ServersChanges `json:",omitempty"`
	Logo         string          `json:"logo"`
	Title        string          `json:"title"`
	IsDown       bool            `json:"is_down"`
}

func (server *Server) validate() (map[string]interface{}, bool) {

	if server.Logo == "" {
		return u.Message(false, "Logo should be on the payload"), false
	}

	if server.Title == "" {
		return u.Message(false, "Title should be on the payload"), false
	}

	return u.Message(true, "The server data is valid"), true
}

func (server *Server) CreateServerByHostname(hostname string) map[string]interface{} {
	if resp, ok := server.validate(); !ok {
		return resp
	}
	if _, exist := GetSeverIDAndCheckIfExistByHostname(hostname); exist {
		return u.Message(false, "The server alredy exist")
	} else {
		query := fmt.Sprintf(`INSERT INTO servers (hostname, logo, title, is_down)
						VALUES ('%s', '%s', '%s', %t)`, hostname, server.Logo, server.Title, server.IsDown)
		fmt.Println(query)
		_, err := GetDB().Exec(query)
		if err != nil {
			fmt.Println(err)
			return u.Message(false, "Error trying to exec the server query")
		}
		for _, serverEndpoint := range server.Servers {
			idServer, _ := GetSeverIDAndCheckIfExistByHostname(hostname)
			resp := serverEndpoint.CreateServerEndpointByServerID(idServer)
			if ok := resp["status"]; ok != true {
				fmt.Println(resp)
				return u.Message(false, "Error trying to create serverEndpoint")
			}
		}

		return u.Message(true, "Server was created")
	}
}

func GetServersHostnames() []string {
	hostnames := make([]string, 0)
	rows, err := GetDB().Query(`SELECT hostname FROM servers;`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var hostname string
		if err := rows.Scan(&hostname); err != nil {
			fmt.Println(err)
		}
		hostnames = append(hostnames, hostname)
	}
	return hostnames
}

func GetServerByDomain(domain string) *Server {
	server := &Server{}
	query := fmt.Sprintf(`SELECT hostname, logo, title, is_down FROM servers WHERE hostname = '%s';`, domain)
	fmt.Println(query)
	row := GetDB().QueryRow(query)

	var hostname string
	err := row.Scan(&hostname, &server.Logo, &server.Title, &server.IsDown)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Cannot found server with domain: ", domain)
		return nil
	case nil:
		server.Servers = GetServerEndpointsByHostname(hostname)
	default:
		fmt.Println(err)
	}
	return server
}

func (server *Server) UpdateServerByHostname(hostname string, serverUpdated *Server) map[string]interface{} {
	query := fmt.Sprintf(`UPDATE servers SET logo = '%s', title = '%s', is_down = %t, updatedat = NOW() WHERE hostname = '%s'`, serverUpdated.Logo, serverUpdated.Title, serverUpdated.IsDown, hostname)
	fmt.Println(query)
	_, err := GetDB().Exec(query)
	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Error trying to exec the update server query")
	}

	for i, serverEndpoint := range server.Servers {
		idServer := serverEndpoint.GetIDByAddress()
		resp := serverUpdated.Servers[i].UpdateServerEndpointByID(idServer)
		if ok := resp["status"]; ok != true {
			fmt.Println(resp)
			return u.Message(false, "Error trying to update serverEndpoint")
		}
	}
	return u.Message(true, "The server was updated")
}

func GetSeverIDAndCheckIfExistByHostname(hostname string) (uint, bool) {
	var id uint

	query := fmt.Sprintf(`SELECT id_server FROM servers WHERE hostname = '%s';`, hostname)
	fmt.Println(query)
	row := GetDB().QueryRow(query)

	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Cannont found ID with hostname: ", hostname)
			return 0, false
		}
		fmt.Println(err)
	}
	return id, true
}

func (server *Server) ServerIsChange(serverAPI *u.ApiServer) map[string]interface{} {
	var isChange map[string]interface{}
	if server.Logo != serverAPI.Logo {
		//Do something
	}
	if server.Title != serverAPI.Title {
		//Do something
	}
	if server.IsDown != serverAPI.IsDown {
		//Do something
	}
	if len(server.Servers) != len(serverAPI.Servers) {
		isChange = u.Message(true, "there was a change on servers len")
	} else {
		gradesforMin := []string{}
		gradesforPrev := []string{}
		for i, serverEndpoint := range server.Servers {
			if serverEndpoint.Address != serverAPI.Servers[i].Address {
				//Do something
			}
			if serverEndpoint.Grade != serverAPI.Servers[i].Grade {
				isChange = u.Message(true, "grade changed")
			}
			if serverEndpoint.Country != serverAPI.Servers[i].Country {
				//Do something
			}
			if serverEndpoint.Owner != serverAPI.Servers[i].Owner {
				//Do something
			}
			gradesforMin = append(gradesforMin, serverAPI.Servers[i].Grade)
			gradesforPrev = append(gradesforPrev, serverEndpoint.Grade)
			if i == (len(server.Servers) - 1) {
				isChange["grade"] = getGrade(gradesforMin)
				isChange["previousGrade"] = getGrade(gradesforPrev)
			}
		}
	}
	if isChange == nil {
		isChange = u.Message(false, "Servers doesnt change")
	}
	return isChange
}

func getGrade(grades []string) string {
	var minGrade string
	sort.Slice(grades, func(i, j int) bool {
		return grades[i] < grades[j]
	})
	minGrade = grades[len(grades)-1]
	return minGrade
}

func EncodingApiToServer(apiServer *u.ApiServer) *Server {
	jsonAPI, err := json.Marshal(apiServer)
	if err != nil {
		fmt.Println("Error while encoding")
	}
	var server *Server

	err = json.Unmarshal(jsonAPI, &server)
	if err != nil {
		fmt.Println("Error while decoding")
	}
	return server
}
