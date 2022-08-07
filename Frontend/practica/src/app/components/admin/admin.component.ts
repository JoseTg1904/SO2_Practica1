import { Component, OnInit, ViewChild  } from '@angular/core';
import { ServicesService } from '../../services.service'
import { DataCPU } from './../../cpu.model';
import { ChartConfiguration, ChartData, ChartType } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';
import DataLabelsPlugin from 'chartjs-plugin-datalabels';
import { Proceso } from '../../proceso.model';

interface datos {
  name:string;
  value:number;
}
@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})


export class AdminComponent implements OnInit {
  @ViewChild(BaseChartDirective) chart: BaseChartDirective | undefined;


  public barChartOptions: ChartConfiguration['options'] = {
    responsive: true,

    // We use these empty structures as placeholders for dynamic theming.
    scales: {

      xAxes: {},
      yAxes: {
        min: 0,



      }
    },
    plugins: {
      legend: {
        display: true,
      },
      datalabels: {
        anchor: 'end',
        align: 'end',

      }
    }
  };
  public barChartType: ChartType = 'bar';
  public barChartPlugins = [
    DataLabelsPlugin
  ];

  public barChartData: ChartData<'bar'> = {
    labels: [ ],
    datasets: [
      { data: [ ], label: 'Frecuencia' },

    ]
  };

  dataCPU: DataCPU = {
    procesos: 0,
    ejecucion: 0,
    suspendidos: 0,
    detenidos: 0,
    zombies: 0,
  };

  foods: datos[] = [

  ];

  valores!: Proceso;

  constructor(
    private servicesService: ServicesService
  ) { }

  ngOnInit(): void {
    this.fetchCPU();
    this.getProcess();
  }
  fetchCPU() {

    this.servicesService.getCPU().subscribe( datos => {
      this.dataCPU=datos;
    });
    console.log("a")



  }
  async getProcess() {
    this.valores = await this.servicesService.getProcessList();
    this.valores.procesos.map(({pid,nombre})=>

      this.foods.push({name:nombre,value:pid})

    );

  }

  async visualizar(valor:number){
    let datos = await this.servicesService.getHistograma(valor);
    console.log(datos);
    this.barChartData.labels = datos.names;
    this.barChartData.datasets[0].data = datos.values;
    this.chart?.chart?.update();

  }

}
