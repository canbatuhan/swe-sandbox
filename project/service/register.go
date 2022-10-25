package service

import (
	"net/http"
	structs "project/structs"
)

func RegisterService(w http.ResponseWriter, req *http.Request) {
	var registerReq structs.RegisterRequest
	var registerResp structs.RegisterResponse

	ReadAndDecodeRequest(&registerReq, req)

	/*
		user_username = database.getUser(registerReq.Username)
		user_email = database.getUser(registerReq.Email)

		if user_username != nil
			registerResp.Registered = false
			registerResp.ErrorMessage = "Username is taken"
		else user_email != nil
			registerResp.Registered = false
			registerResp.ErrorMessage = "Email is in use"
	*/
	if registerReq.Password != registerReq.ConfirmPassword {
		registerResp.Registered = false
		registerResp.ErrorMessage = "Passwords does not match"
	} else {
		/*
			database.addUser(User{registerReq.Email,
							registerReq.Username,
							registerReq.Password})
		*/
		registerResp.Registered = true
		registerResp.ErrorMessage = ""
	}

	EncodeAndWriteResponse(registerResp, w)
}
