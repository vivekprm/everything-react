package controllers

import (
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Templates struct {
		New Template
	}
}

func (l AuthRequest) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password:", r.FormValue("password"))
}
