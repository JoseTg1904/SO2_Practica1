export interface DataRam {
  total: number;
  consumida: number;
  porcentaje: number;
  grafica: grafa[]
}

interface grafa {
  tiempo: string,
  consumo: number
}