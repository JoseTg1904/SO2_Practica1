package structs

/*
	struct ResponseMessage = para respuesta inicial, errores y del kill
*/
type ResponseMessage struct {
	Message string `json:"message"`
}

/*
	struct Graphic = para el tiempo y consumo de la memoria ram en un determinado tiempo
*/
type Graphic struct {
	Tiempo string `json:"tiempo"`
	Consumo float32 `json:"consumo"`
}

/*
	struct ResponseRAM = para respuesta del modulo de RAM
*/
type ResponseRAM struct {
	Consumida float32 `json:"consumida"`
	Porcentaje float32 `json:"porcentaje"`
	Total float32 `json:"total"`
	Grafica []Graphic `json:"grafica"`
}

/*
	struct ResponseProcessQuantity = para respuesta de la cantidad de procesos
*/
type ResponseProcessQuantity struct {
	Procesos float32 `json:"procesos"`
	Ejecucion float32 `json:"ejecucion"`
	Suspendidos float32 `json:"suspendidos"`
	Detenidos float32 `json:"detenidos"`
	Zombies float32 `json:"zombies"`
}

/*
	struct ResponseProcessList = para respuesta de la lista de procesos
*/
type ResponseProcessList struct {
	Procesos []ProcessList `json:"procesos"`
}

/*
	struct ResponseStrace = para respuesta del strace -p pid
*/
type ResponseStrace struct {
	Data string `json:"data"`
}

/*
	struct GetStraceHistograma = para respuesta del strace para el histograma
*/
type ResponseStraceHistograma struct {
	Names []string `json:"names"`
	Values []int `json:"values"`
}


/*
	RequestStraceSyscalls = para respuesta de las syscalls seleccionadas
*/
type RequestStraceSyscalls struct {
	Syscalls []string `json:"syscalls"`
}