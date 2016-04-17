package controllers

import (
	"net/http"
)

type AuthController struct{
}

func FacebookCallback(w http.ResponseWriter, r *http.Request){
	addCors(w, r)
}
