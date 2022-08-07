package concurrence

import (
	"api/src/globals"
	"api/src/helpers"
	"fmt"
	"log"
	"strings"
	"sync"
)

/*
	FillMapSyscalls = se encarga de llenar el MAP de las llamadas del sistema
*/
func FillMapSyscalls() {

	globals.MapStrace = make(map[string]int)
	log.Println("INICIO de MAP")
	globals.FlagActiveMap = true

	var i int = 0
	var name string = ""

	data := helpers.ReadFileSyscalls()
	lines := strings.Split(data, "\n")

	for i = 0; i < len(lines); i ++ {
		name = strings.Split(lines[i], "(")[0]
		if len(name) > 2 {addSyscallMap(name) }
	}

	globals.FlagActiveMap = false
	log.Println("Fin de MAP")
}

/*
	FillArraSyscalls = se encarga de llenar en un string las llamadas del sistema seleccionadas desde el frontend
*/
func FillASyscallsSelected(array []string) {
	globals.ArrayStrace2 = ""
	log.Println("Inicio de ARRAY")
	globals.FlagActiveArray = true

	var wg sync.WaitGroup
	var i int = 0
	var j int = 0

	data := helpers.ReadFileSyscalls()
	lines := strings.Split(data, "\n")

	for i = 0; i < len(lines); i += 10 {
		for j = 0; j < 10 && (i+j) < len(lines); j++ {
			wg.Add(1)
			go func(i int, j int, array []string) {
				defer wg.Done()
				
				name := strings.Split(lines[i+j], "(")[0]
				for x := 0; x < len(array); x++ {
					if name == array[x] {
						globals.ArrayStrace2 += fmt.Sprintf("%s\n",lines[i+j])
					}
				}

			}(i, j, array)
		}
		wg.Wait()
	}			
	globals.FlagActiveArray = false
	log.Println("Fin de ARRAY")
}



/*
	addSyscallMap = agrega el nombre de un systemCall al MAP
*/
func addSyscallMap(name string) {

	if el, ok := globals.MapStrace[name]; ok {
		el++
		globals.MapStrace[name] = el
	} else {
		globals.MapStrace[name] = 1
	}
}
