import { Component } from '@angular/core';
import { NestedTreeControl } from '@angular/cdk/tree';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { ServicesService } from 'src/app/services.service';

interface FoodNode {
  pid: number;
  usuario?: number;
  estado?: string;
  nombre: string;
  hijos?: FoodNode[];
}

@Component({
  selector: 'app-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css'],
})
export class TreeComponent {
  panelOpenState = false;
  process: any;
  syscalls: any[] = [];

  treeControl = new NestedTreeControl<FoodNode>((node) => node.hijos);
  dataSource = new MatTreeNestedDataSource<FoodNode>();

  constructor(private servicesService: ServicesService) {
    this.getProcess();
  }

  async getProcess() {
    const list = await this.servicesService.getProcessList();
    this.dataSource.data = list.procesos;
  }

  async getGeneralStrace(pid: number) {
    const list = await this.servicesService.getGeneralStrace(pid);
    let txt = document.getElementById('generalStrace') as HTMLElement;
    txt.innerHTML = list.data == '' ? 'Ninguno' : list.data;
  }

  add() {
    let txt = document.getElementById('syscallInput') as HTMLInputElement;
    this.syscalls.push(txt.value);
    txt.value = '';
  }

  delete() {
    this.syscalls = [];
  }

  async getSyscall(pid: number) {
    const list = await this.servicesService.getSyscall(pid, this.syscalls);
    let txt = document.getElementById('resultSyscall') as HTMLElement;
    txt.innerHTML = list.data == '' ? 'Ninguno' : list.data;
  }

  async getKill(pid: number) {
    const list = await this.servicesService.getKill(pid);
    alert(list.message);
    this.getProcess();
  }

  hasChild = (_: number, node: FoodNode) =>
    !!node.hijos && node.hijos.length > 0;
}
