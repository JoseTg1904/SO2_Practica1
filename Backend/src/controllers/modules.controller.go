package controllers

import (
	"api/src/concurrence"
	"api/src/globals"
	"api/src/helpers"
	"api/src/structs"
	"net/http"
)

/*
	RamRoute = se encarga de responde el contenido del modulo de RAM
	endpoint = /ram
*/
func RamRoute(res http.ResponseWriter, req *http.Request) {
    helpers.RespondWithJSON(res, http.StatusOK, globals.DataRAM)
}

/*
	ProcessQuantitiesRoute = se encarga de enviar la cantidad de procesos ejecutandose
	endpoint = /procesos/numeros
*/
func ProcessQuantitiesRoute(res http.ResponseWriter, req *http.Request) {
	
	data, err := helpers.ReadProcessQuantities()

	if err != nil {
    	helpers.RespondWithJSON(res, http.StatusInternalServerError, structs.ResponseMessage{
        Message: err.Error()})
		return
	}

    helpers.RespondWithJSON(res, http.StatusOK, data)
}

/*
	ProcessListRoute = Se encarga de enviar la lista de procesos ejecutandose
	endpoint = /procesos/lista
*/
func ProcessListRoute(res http.ResponseWriter, req *http.Request) {
	
	data, err := helpers.ReadProcessList()

	if err != nil {
    	helpers.RespondWithJSON(res, http.StatusInternalServerError, structs.ResponseMessage{
        Message: err.Error()})
		return
	}

	data2 := concurrence.NameUser(data)
    helpers.RespondWithJSON(res, http.StatusOK, data2)
}