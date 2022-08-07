package structs

/*
	struct RAM = estructura del contenido que tendrán el archivo de modulo de RAM
*/
type RAM struct {
	Consumida float32 `json:"consumida"`
	Porcentaje float32 `json:"porcentaje"`
	Total float32 `json:"total"`
}

/*
	struct processListChild = hijos de un procesos
*/
type processListChild struct {
	Pid int `json:"pid"`
	Nombre string `json:"nombre"`
	Hijos []processListChild `json:"hijos"`
}

/*
	struct ProcessList = padre superior de todos los procesos
*/
type ProcessList struct {
	Pid int `json:"pid"`
	Nombre string `json:"nombre"`
	Usuario string `json:"usuario"`
	Estado string `json:"estado"`
	Hijos []processListChild `json:"hijos"`
}
