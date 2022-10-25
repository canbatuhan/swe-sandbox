package service

import (
	"net/http"
	structs "project/structs"
)

func LoginService(w http.ResponseWriter, req *http.Request) {
	var loginReq structs.LoginRequest
	var loginResp structs.LoginResponse

	ReadAndDecodeRequest(&loginReq, req)
	/*
		user = database.getUser(loginReq.Username)
		if user != nil && user.Password == loginReq.Password
			loginResp.LoggedIn = true
			loginResp.ErrorMessage = ""
		else
			loginResp.LoggedIn = false
			loginResp.ErrorMessage= "Username or password is incorrect"
	*/
	loginResp.LoggedIn = true
	loginResp.ErrorMessage = ""
	EncodeAndWriteResponse(loginResp, w)
}
