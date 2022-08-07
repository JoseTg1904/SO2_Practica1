package helpers

import (
	"api/src/globals"
	"api/src/structs"
	"strings"
	"time"
)

/*
	ReadModuleRam = Se encarga de leer el modulo de la RAM
*/
func ReadModuleRam() (error) {
	var dataRAM structs.RAM

	// Leo el Archivo de la RAM
	content, err := ReadFile("/proc/modulo_ram")
	if err != nil {
		return err
	}

	// Convertir a Struct
	dataRAM, err = ByteToRamStruct(content)
	if err != nil {
		return  err
	}

	globals.DataRAM.Consumida = dataRAM.Consumida
	globals.DataRAM.Porcentaje = dataRAM.Porcentaje
	globals.DataRAM.Total = dataRAM.Total

	if len(globals.DataRAM.Grafica) >= 30 {
		globals.DataRAM.Grafica = globals.DataRAM.Grafica[1:30]
	} 

	now := time.Now()
	currentTime := now.Format("15:04:05")

	globals.DataRAM.Grafica = append(globals.DataRAM.Grafica, structs.Graphic{Tiempo: currentTime, Consumo: dataRAM.Consumida})

	return  nil
}



/*
	ReadProcessQuantities = Obtiene la Cantidades de procesos
*/
func ReadProcessQuantities() (structs.ResponseProcessQuantity, error) {
	var dataProcessQuantity structs.ResponseProcessQuantity

	// Leo el Archivo de Procesos
	content, err := ReadFile("/proc/modulo_procesos_numero")
	if err != nil {
		return dataProcessQuantity, err
	}

	// Convertir a Struct
	dataProcessQuantity, err = ByteToResponseProcessQuantity(content)
	if err != nil {
		return  dataProcessQuantity, err
	}

	return  dataProcessQuantity, nil
}


/*
	ReadProcessList = Obtiene la lista de procesos
*/
func ReadProcessList() ([]structs.ProcessList, error) {
	var dataProcessQuantity []structs.ProcessList

	// Leo el Archivo de Procesos lista
	content, err := ReadFile("/proc/modulo_procesos_lista")

    b := strings.Join(strings.Fields(strings.TrimSpace(string(content))), "")
	b = strings.ReplaceAll(b, ",]", "]")


	if err != nil {
		return dataProcessQuantity, err
	}

	// Convertir a Struct
	c := []byte(b)
	dataProcessQuantity, err = ByteToProcessList(c)
	if err != nil {
		return  dataProcessQuantity, err
	}

	return  dataProcessQuantity, nil
}