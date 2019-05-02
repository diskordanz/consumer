import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';

const routes: Routes = [
  {
    path: 'franchises',
    component: FranchiseListComponent,
    data: { title: 'List of Franchises' }
  }
];

@NgModule({ 
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
