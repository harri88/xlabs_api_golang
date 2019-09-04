package controllers

import (
	"net/http"

	"github.com/harri88/xlabs_api_golang/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
