package controllers

import (
	"api/src/helpers"
	"api/src/structs"
	"net/http"
)

/*
	IndexRoute = respuesta de ruta inicial
	endpoint = /
*/
func IndexRoute(res http.ResponseWriter, req *http.Request) {
    helpers.RespondWithJSON(res, http.StatusOK, structs.ResponseMessage{
        Message: " SOPES 2 - Group 30 - API in GO ",
    })
}