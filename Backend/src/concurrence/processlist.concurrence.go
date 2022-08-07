package concurrence

import (
	"api/src/structs"
	"log"
	"os/user"
	"sync"
)

/*
	NameUser = se encarga de obtener el nombre del usuario mediante concurrencia
*/
func NameUser(processlist []structs.ProcessList) (structs.ResponseProcessList) {
	var wg sync.WaitGroup
	var i int = 0
	var j int = 0

	for i = 0; i < len(processlist); i += 10 {
		for j = 0; j < 10 && (i+j) < len(processlist); j++ {
			wg.Add(1)
			go func(i int, j int ) {
				defer wg.Done()
				
				username := ""
				uid := processlist[i+j].Usuario
				user, err := user.LookupId(uid)

				if err != nil {
					log.Println("Error al obtener el nombre del usuario: ", uid)
					username = uid
				} else {
					username = user.Username
				}		

				processlist[i+j].Usuario = username
			}(i, j)
		}
		wg.Wait()
	}


	var res structs.ResponseProcessList
	res.Procesos = processlist
	return res

}