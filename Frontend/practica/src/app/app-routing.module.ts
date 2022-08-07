import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminComponent } from './components/admin/admin.component';
import { RamComponent } from './components/ram/ram.component';
import { TreeComponent } from './components/tree/tree.component';

const routes: Routes = [
  { path: 'admin', component: AdminComponent },
  { path: 'tree', component: TreeComponent },
  { path: '', component: RamComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
