package helpers

import (
	"api/src/globals"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
	WriteFileSyscalls = escribe en un archivo
*/
func WriteFileSyscalls(x string) int {
	f, err := os.OpenFile(globals.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	    }
	_, errf := f.WriteString(x + "\n")
	if errf != nil {
		fmt.Println(errf)
	}
	f.Close()
	return 0
}

/*
	EraseFileSyscalls = Elimina el Archivo
*/
func EraseFileSyscalls() int {
    err := os.Remove(globals.Path)
    if err != nil {
        log.Println(err.Error())
    }
	return 0
}

/*
	ReadFileSyscalls = Lee el archivo
*/
func ReadFileSyscalls() string {
    bytes, err := ioutil.ReadFile(globals.Path)
    if err != nil {
        log.Println(err.Error())
		return ""
    }
    data := string(bytes)
	return data
}
