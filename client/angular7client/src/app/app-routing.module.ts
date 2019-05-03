import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';
import { LocationListComponent } from './location-list/location-list.component';

const routes: Routes = [
  {
    path: 'franchises',
    component: FranchiseListComponent,
  },
  {
    path: 'franchises/:id/locations',
    component: LocationListComponent,
  }
];

@NgModule({ 
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
