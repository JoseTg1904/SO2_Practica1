package helpers

import (
	"api/src/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	respondWithJSON = envia respuesta en formato json

	Parametros:
		* res =
		* codeStatus = codigo de status de petición
		* strucRes = estructura a enviar en formato JSON

*/
func RespondWithJSON(res http.ResponseWriter, codeStatus int, strucRes interface{}) error {
    response, err := json.Marshal(strucRes)
    if err != nil {
        return err
    }
    res.Header().Set("Content-Type", "application/json")
    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.WriteHeader(codeStatus)
    res.Write(response)
    return nil
}


/*
	ReadFile = Lee un Archivo

	Parametros:
		* path = ruta del archivo
*/
func ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("Error al leer el archivo: %s , Error: %s \n", path, err.Error())
		return content, fmt.Errorf("error al leer el archivo del modulo")
	}

	return content, nil
}



/*
	ByteToRamStruct = Convierte de []byte a struct.RAM{}
*/
func ByteToRamStruct(content []byte) (structs.RAM, error) {
	var data structs.RAM

	err := json.Unmarshal(content, &data)

	if err != nil {
		log.Printf("Error convert JSON = %s \n",err.Error())
		return data, fmt.Errorf("error al leer el archivo ram json")
	}

	return data, nil
}



/*
	ByteToResponseProcessQuantity = Convierte de []byte a struct.ResponseProcessQuantity{}
*/
func ByteToResponseProcessQuantity(content []byte) (structs.ResponseProcessQuantity, error) {
	var data structs.ResponseProcessQuantity

	err := json.Unmarshal(content, &data)

	if err != nil {
		log.Printf("Error convert JSON = %s \n",err.Error())
		return data, fmt.Errorf("error al leer el archivo procesos numero json")
	}

	return data, nil
}



/*
	ByteToProcessList = Convierte de []byte a struct.ProcessList{}
*/
func ByteToProcessList(content []byte) ([]structs.ProcessList, error) {
	var data []structs.ProcessList

	err := json.Unmarshal(content, &data)

	if err != nil {
		log.Printf("Error convert JSON = %s \n",err.Error())
		return data, fmt.Errorf("error al leer el archivo lista de procesos json")
	}

	return data, nil
}


