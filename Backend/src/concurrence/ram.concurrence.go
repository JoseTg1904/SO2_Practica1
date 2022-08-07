package concurrence

import (
	"api/src/helpers"
	"time"
)

/*
	GetDataRama = se encarga de ejecutar cada 3 segundo la lectura del modulo de RAM
*/
func GetDataRama()  {
	go func() {
		for {
			helpers.ReadModuleRam()
			time.Sleep(3*time.Second)
		}
	}()
}