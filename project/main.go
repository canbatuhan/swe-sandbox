package main

import (
	"net/http"
	service "project/service"
)

func main() {
	http.HandleFunc("/api/register", service.RegisterService)
	http.HandleFunc("/api/login", service.LoginService)
	http.ListenAndServe(":8090", nil)
}
