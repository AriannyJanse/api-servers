package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type ApiEndpoint struct {
	Address string `json:"address"`
	Grade   string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner   string `json:"owner"`
}

type Endpoint struct {
	IPAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

type Response struct {
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
}

func GetApiEndpoints(domain string) ([]*ApiEndpoint, bool) {
	fmt.Println("Statirng GetApiEndpoints")
	apiEndpoints := make([]*ApiEndpoint, 0)
	dataEndpoints := getEndpointsAndStatus(domain)

	if dataEndpoints.Status != "DNS" && dataEndpoints.Status != "ERROR" {
		for _, endpoint := range dataEndpoints.Endpoints {
			ip := endpoint.IPAddress
			apiEndpoint := ApiEndpoint{
				Address: ip,
				Grade:   endpoint.Grade,
				Country: whoIs(ip, "Country"),
				Owner:   whoIs(ip, "OrgName"),
			}
			apiEndpoints = append(apiEndpoints, &apiEndpoint)
		}
		fmt.Println("Ending GetApiEndpoints")
		return apiEndpoints, false
	}
	fmt.Println("Ending GetApiEndpoints")
	return nil, true
}

func getEndpointsAndStatus(domain string) *Response {
	requestURL := "https://api.ssllabs.com/api/v3/analyze?host=" + domain
	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseObject := &Response{}
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func whoIs(ip, key string) string {
	cmd := fmt.Sprintf("whois %s | grep -iE ^'%s:'", ip, key)
	response, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	resp := strings.TrimSpace(strings.ReplaceAll(string(response), key+":", ""))
	return resp
}
