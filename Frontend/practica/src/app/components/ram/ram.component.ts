import { Component, OnInit,ViewChild ,SimpleChanges   } from '@angular/core';
import { ChartConfiguration,  ChartType } from 'chart.js';
import { BaseChartDirective } from 'ng2-charts';
import { ServicesService } from '../../services.service'
import { DataRam } from '../../ram.model'
import { interval } from 'rxjs';

@Component({
  selector: 'app-ram',
  templateUrl: './ram.component.html',
  styleUrls: ['./ram.component.css']
})
export class RamComponent implements OnInit {

  dataRam: DataRam = {
    total: 0,
    consumida: 0,
    porcentaje: 0,
    grafica: [{ "tiempo": "22:48:55", "consumo": 48 },
    { "tiempo": "22:49:00", "consumo": 60 },
    { "tiempo": "22:49:05", "consumo": 50 }
]
  };
  public timerInterval:any;

  public lineChartData: ChartConfiguration['data'] = {
    datasets: [
      {
        data: [ 65, 59, 80, 81, 56, 55, 40 ],
        label: 'MB',
        backgroundColor: 'rgba(148,159,177,0.2)',
        borderColor: 'rgba(148,159,177,1)',
        pointBackgroundColor: 'rgba(148,159,177,1)',
        pointBorderColor: '#fff',
        pointHoverBackgroundColor: '#fff',
        pointHoverBorderColor: 'rgba(148,159,177,0.8)',
        fill: 'origin',
      },
    ],
    labels: [ 'January', 'February', 'March', 'April', 'May', 'June', 'July' ]
  };

  public lineChartOptions: ChartConfiguration['options'] = {
    elements: {
      line: {
        tension: 0.1
      }
    },
    scales: {
      // We use this empty structure as a placeholder for dynamic theming.
      x: {},
      'y-axis-0':
        {
          position: 'left',
        },
      'y-axis-1': {
        position: 'right',
        grid: {
          color: 'rgba(255,0,0,0.3)',
        },
        ticks: {
          color: 'red'
        }
      }
    },

    plugins: {
      legend: { display: true },

    }
  };

  public lineChartType: ChartType = 'line';

  @ViewChild(BaseChartDirective) chart?: BaseChartDirective;
  constructor(
    private servicesService: ServicesService
  ) { }

  ngOnInit(): void {
    this.timerInterval = interval(1000);
    this.timerInterval.subscribe( () =>{

      this.fetchProducts();

    });


  }
  ngOnDestroy() {
    // Will clear when component is destroyed e.g. route is navigated away from.
    clearInterval(this.timerInterval);
 }

  async fetchProducts() {

    this.dataRam = await this.servicesService.getRam();


    this.lineChartData.labels = this.dataRam.grafica.map(({tiempo})=>tiempo);
    this.lineChartData.datasets[0].data =this.dataRam.grafica.map(({consumo})=>consumo);
    this.chart?.chart?.update();
  }

}
