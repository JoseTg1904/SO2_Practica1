package controllers

import (
	"api/src/concurrence"
	"api/src/globals"
	"api/src/helpers"
	"api/src/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

/*
	GetStrace = se encarga de responder el contenido del comando Strace
	endpoint = /strace/:pid
*/
func GetStrace(res http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)

	pid, err := strconv.Atoi(urlParams["pid"])

	if err != nil {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "El PID no era un número entero"})
		return
	}
	
	execStrace(pid)

	if !globals.ChangePID { 
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: fmt.Sprintf("No se Logro Attachar el proceso %d, puede intentar otra vez, si en dado caso sigue recibiendo este error, entonces significa que el Proceso %d no es permitido Attachar", pid, pid)})
		return
	}

	globals.Stringstrace = helpers.ReadFileSyscalls()
	
    helpers.RespondWithJSON(res, http.StatusOK, structs.ResponseStrace{Data: globals.Stringstrace})
}

/*
	GetStraceHistograma = se encarga de responder el contenido dividio por syscalls
	endpoint = /strace/histograma/:pid  
*/
func GetStraceHistograma(res http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)

	pid, err := strconv.Atoi(urlParams["pid"])

	if err != nil {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "El PID no era un número entero"})
		return
	}

	execStrace(pid)


	if !globals.ChangePID { 
    	helpers.RespondWithJSON(res, 500, structs.ResponseMessage{Message: fmt.Sprintf("No se Logro obtener las Llamadas del Sistema del proceso %d, puede intentar otra vez, si en dado caso sigue recibiendo este error, entonces significa que el Proceso %d no es permitido obtener sus llamadas al sistema", pid, pid)})
		return
	}



	if !globals.FlagActiveMap {
		go concurrence.FillMapSyscalls()
		time.Sleep(1*time.Second)
	} 
	

    keys := []string{}
    values := []int{}

	for key, el := range globals.MapStrace {
		keys = append(keys, key)
		values = append(values, el)
    }

    helpers.RespondWithJSON(res, http.StatusOK, structs.ResponseStraceHistograma{Names: keys, Values: values})
}

/*
	execStrace = se encarga de ejecutar el comando strace de un pid diferente al anterior
*/
func execStrace(pid int) {
	if pid != globals.PID {
		helpers.EraseFileSyscalls()
		globals.FlagDieStrace = true;
		time.Sleep(3*time.Second)


		go func() {
			log.Printf("----- Empezando go Rutina - ID = %d ----- \n\n",pid)
			helpers.Attach(pid)
			log.Printf("----- Terminando go Rutina - ID = %d----- \n\n",pid)
		}()	

		time.Sleep(500*time.Millisecond)
		if globals.ChangePID { globals.PID = pid }
	}
}



/*
	PostStraceSysCalls = retorna la cantidad de veces repetidas de las syscalls seleccionadas
	endpoint = /strace/syscalls/:pid
*/
func PostStraceSysCalls(res http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)

	pid, err := strconv.Atoi(urlParams["pid"])
	if err != nil {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "El PID no era un número entero"})
		return
	}

	if pid != globals.PID {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "No se ha inicializado el Proceso anteriormente"})
		return
	} 

	var syscalls structs.RequestStraceSyscalls

	err = json.NewDecoder(req.Body).Decode(&syscalls)
	if err != nil {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "Forma incorrecta del cuerpo POST"})
		return
	}

	

	if !globals.FlagActiveMap {
		go concurrence.FillASyscallsSelected(syscalls.Syscalls)
		time.Sleep(1*time.Second)
	} 
	


    helpers.RespondWithJSON(res, http.StatusOK, structs.ResponseStrace{Data: globals.ArrayStrace2})
}

/*
	GetKillProcess = se encarga de matar el proceso enviado 
	endpoint = /kill/:id
*/
func GetKillProcess(res http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)

	pid, err := strconv.Atoi(urlParams["pid"])

	if err != nil {
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "El PID no era un número entero"})
		return
	}
	
    err = syscall.Kill(pid, syscall.SIGKILL)

    if err != nil {
        log.Println("ERROR KILL = ", err)
    	helpers.RespondWithJSON(res, 400, structs.ResponseMessage{Message: "No se pudo Eliminar el Proceso"})
		return
    }
	
    helpers.RespondWithJSON(res, http.StatusOK, structs.ResponseMessage{Message: "Proceso Eliminado Satisfactoriamente"})
}

