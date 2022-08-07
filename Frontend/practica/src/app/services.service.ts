import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DataRam } from './ram.model';
import { DataCPU } from './cpu.model';
import { Proceso } from './proceso.model'

@Injectable({
  providedIn: 'root',
})
export class ServicesService {
  constructor(public http: HttpClient) {}
  api_url = 'http://34.125.201.170:5000';

  async getProcessList(): Promise<any> {
    return this.http.get<any>(this.api_url + '/procesos/lista').toPromise();
  }

  async getGeneralStrace(pid: number): Promise<any> {
    return this.http.get<any>(this.api_url + '/strace/' + pid).toPromise();
  }

  async getKill(pid: number): Promise<any> {
    return this.http.get<any>(this.api_url + '/kill/' + pid).toPromise();
  }

  async getSyscall(pid: number, list: any[]): Promise<any> {
    return this.http
      .post<any>(this.api_url + '/strace/syscalls/' + pid, {
        syscalls: list,
      })
      .toPromise();
  }

  async getRam(): Promise<any> {
    return this.http.get<DataRam>(this.api_url + '/ram').toPromise();
  }

  getCPU(){
    return this.http.get<DataCPU>(this.api_url + '/procesos/numeros');
  }
  async getList(): Promise<any> {
    return this.http.get<Proceso>(this.api_url + '/procesos/lista').toPromise();
  }
  async getHistograma(id: number) : Promise<any> {
    return this.http.get<any>(`${this.api_url}/strace/histograma/${id}`).toPromise();
  }
}
