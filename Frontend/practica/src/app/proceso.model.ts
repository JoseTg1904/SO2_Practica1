export interface Proceso {
  procesos: dato[]
}
interface dato {
  pid:number;
  nombre:string;
  usuario:string;
  estado:string;
  hijos: dato[]
}
