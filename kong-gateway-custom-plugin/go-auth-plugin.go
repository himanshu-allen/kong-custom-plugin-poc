// auth_plugin.go

package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type AuthPlugin struct{}

func New() interface{} {
	return &AuthPlugin{}
}

func (p *AuthPlugin) Access(kong *pdk.PDK) {
	// Retrieve the request object
	authorization, err := kong.Request.GetHeader("Authorization")
	if err != nil {
		kong.Log.Err(err.Error())
		kong.Response.Exit(500, "Internal Server Error", nil)
		return
	}
	kong.Response.SetHeader("X-Custom-Auth-Check", "Checked")
	kong.Log.Info("Authorization Header found : ", authorization)
	kong.Log.Info("Sending Req to Auth Service for auth....")

	//// Make a call to the authentication service
	//// Replace the following with your authentication service endpoint
	//authServiceURL := "http://auth-service.example.com/authenticate"
	//authReq, err := http.NewRequest("GET", authServiceURL, nil)
	//if err != nil {
	//	kong.Log.Err(err.Error())
	//	kong.Response.Exit(500, "Internal Server Error", nil)
	//	return
	//}
	//authReq.Header.Add("Authorization", authorization)
	//
	//authResp, err := http.DefaultClient.Do(authReq)
	//if err != nil {
	//	kong.Log.Err(err.Error())
	//	kong.Response.Exit(500, "Internal Server Error", nil)
	//	return
	//}
	//defer authResp.Body.Close()
	//
	//// Check the authentication response and proceed accordingly
	//if authResp.StatusCode != http.StatusOK {
	//	kong.Response.Exit(401, "Unauthorized", nil)
	//	return
	//}
	// Continue with the request flow if the authentication is successful
	kong.Log.Info("Authentication successful")
}

func main() {
	Version := "1.1"
	Priority := 1400
	_ = server.StartServer(New, Version, Priority)
}
