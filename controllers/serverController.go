package controllers

import (
	"api-server/models"
	u "api-server/utils"
	"fmt"

	"github.com/valyala/fasthttp"
)

var GetServerByDomain = func(ctx *fasthttp.RequestCtx) {
	domain := ctx.UserValue("domain").(string)
	apiResp := u.GetApiServerByHostname(domain)
	fmt.Println(apiResp)

	if idSever, exist := models.GetSeverIDAndCheckIfExistByHostname(domain); exist {
		server := models.GetServerByDomain(domain)
		fmt.Println(server)
		if server != nil && apiResp != nil {
			isChange := server.ServerIsChange(apiResp)
			if isChange["status"] == true {
				serverUpdated := models.EncodingApiToServer(apiResp)
				server.UpdateServerByHostname(domain, serverUpdated)
				serversChanges := &models.ServersChanges{}
				serverUpdated.ServerChange = serversChanges
				if isChange["grade"] != nil || isChange["previousGrade"] != nil {
					serverUpdated.ServerChange.IsChanged = true
					serverUpdated.ServerChange.Grade = isChange["grade"].(string)
					serverUpdated.ServerChange.PreviousGrade = isChange["previousGrade"].(string)
					serverUpdated.ServerChange.CreateServersChangesByID(idSever)
				} else {
					serverUpdated.ServerChange = nil
				}
				resp := u.Message(true, "success")
				resp["data"] = serverUpdated
				u.Respond(ctx, resp)
			} else {
				resp := u.Message(true, "success")
				resp["data"] = server
				u.Respond(ctx, resp)
			}
		} else {
			u.Respond(ctx, u.Message(false, "Cannot found the server or server is down"))
		}
	} else {
		if apiResp != nil {

			server := models.EncodingApiToServer(apiResp)
			server.CreateServerByHostname(domain)

			data := models.GetServerByDomain(domain)
			resp := u.Message(true, "success")
			resp["data"] = data
			u.Respond(ctx, resp)
		} else {
			u.Respond(ctx, u.Message(false, "Server is down"))
		}
	}
}

var GetServers = func(ctx *fasthttp.RequestCtx) {
	data := models.GetServersHostnames()
	resp := u.Message(true, "success")
	resp["items"] = data
	u.Respond(ctx, resp)
}
