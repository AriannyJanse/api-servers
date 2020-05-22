package utils

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type ApiServer struct {
	Servers []*ApiEndpoint
	Logo    string `json:"logo"`
	Title   string `json:"title"`
	IsDown  bool   `json:"is_down"`
}

func GetApiServerByHostname(domain string) *ApiServer {
	servers, status := GetApiEndpoints(domain)

	if !status {
		apiServer := &ApiServer{
			Servers: servers,
			Logo:    getLogo(domain),
			Title:   getTitle(domain),
			IsDown:  status,
		}
		return apiServer
	} else {
		fmt.Println("Cannot get server, status: ", status)
		return nil
	}
}

func getTitle(domain string) string {
	url := "https://" + domain
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	title := "Not Found"

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return title
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenType = tokenizer.Next()

				if tokenType == html.TextToken {
					resul := tokenizer.Token().Data
					t := &title
					*t = resul
					break
				}
			}
		}
	}
	return title
}

func getLogo(domain string) string {
	url := "https://" + domain
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	logo := "Not Found"

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return logo
		}

		if tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()

			if token.Data == "link" {
				okRel, relVal := searchKey("rel", token.Attr)
				okRef, refVal := searchKey("href", token.Attr)
				fmt.Println("rel:", relVal, "value:", refVal)
				if okRel && okRef && relVal == "shortcut icon" {
					logo = refVal
					break
				}
			}
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "link" {
				tokenType = tokenizer.Next()
				okRel, relVal := searchKey("rel", token.Attr)
				okRef, refVal := searchKey("href", token.Attr)
				fmt.Println("rel:", relVal, "value:", refVal)
				if okRel && okRef && relVal == "shortcut icon" {
					logo = refVal
					break
				}
			}
		}
	}
	return logo
}

func searchKey(key string, lista []html.Attribute) (bool, string) {
	for _, attribute := range lista {
		if attribute.Key == key {
			return true, attribute.Val
		}
	}
	return false, ""
}
