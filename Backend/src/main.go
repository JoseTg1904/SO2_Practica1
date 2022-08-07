package main

import (
	"api/src/concurrence"
	"api/src/controllers"
	"api/src/globals"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const PORT int = 5000;

func main() {
	globals.PID = -1
	router := mux.NewRouter().StrictSlash(true)



	// * -------------------- CONCURRENCIA -------------------------
	concurrence.GetDataRama() 



	// * -------------------- ROUTES -------------------------
	router.HandleFunc("/", controllers.IndexRoute).Methods("GET")

	router.HandleFunc("/ram", controllers.RamRoute).Methods("GET")

	router.HandleFunc("/procesos/numeros", controllers.ProcessQuantitiesRoute).Methods("GET")
	router.HandleFunc("/procesos/lista", controllers.ProcessListRoute).Methods("GET")

	router.HandleFunc("/strace/{pid}", controllers.GetStrace).Methods("GET")
	router.HandleFunc("/strace/histograma/{pid}", controllers.GetStraceHistograma).Methods("GET")
	router.HandleFunc("/strace/syscalls/{pid}", controllers.PostStraceSysCalls).Methods("POST")
	router.HandleFunc("/kill/{pid}", controllers.GetKillProcess).Methods("GET")




	// *  --------------------  SETTINGS Server API  -------------------- 

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  40 * time.Second,
		Handler:      router,
	}

	// *  -------------------- START SERVER -------------------- 
	log.Printf("Server started at %d", PORT)
	log.Println(server.ListenAndServe())
}