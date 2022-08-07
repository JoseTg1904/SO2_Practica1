import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HttpClient } from '@angular/common/http';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSliderModule } from '@angular/material/slider';
import { AppRoutingModule } from './app-routing.module';
import { RouterModule } from '@angular/router';
import { RamComponent } from './components/ram/ram.component';
import { TreeComponent } from './components/tree/tree.component';
import { AdminComponent } from './components/admin/admin.component';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatTreeModule } from '@angular/material/tree';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatDividerModule } from '@angular/material/divider';
import { NgChartsModule } from 'ng2-charts';
import { AgGridModule } from 'ag-grid-angular';
import { AgChartsAngularModule } from 'ag-charts-angular';
import { MatSelectModule } from '@angular/material/select'
@NgModule({
  declarations: [AppComponent, RamComponent, TreeComponent, AdminComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatSliderModule,
    AppRoutingModule,
    RouterModule,
    MatExpansionModule,
    MatTreeModule,
    MatIconModule,
    MatButtonModule,
    MatDividerModule,
    NgChartsModule,
    HttpClientModule,
    AgGridModule.withComponents([]),
    AgChartsAngularModule,
    MatSelectModule

  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
