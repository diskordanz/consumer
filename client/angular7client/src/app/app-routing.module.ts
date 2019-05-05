import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';
import { FranchiseGetComponent } from './franchise-get/franchise-get.component';
import { AppComponent } from './app.component';
import { ProductListComponent } from './product-list/product-list.component';
import { ProductGetComponent } from './product-get/product-get.component';
import { ProfileComponent } from './profile/profile.component';
import { OrderListComponent } from './order-list/order-list.component';
import { OrderGetComponent } from './order-get/order-get.component';
import { CartComponent } from './cart/cart.component';

const routes: Routes = [
  {
    path: 'franchises',
    component: FranchiseListComponent,
  },
  { 
    path: 'franchises/:id', 
    component: FranchiseGetComponent,
  },
  {
    path: 'products',
    component: ProductListComponent,
  },
  {
    path: 'products/:id',
    component: ProductGetComponent,
  },
  {
    path: 'profile',
    component: ProfileComponent,
  },
  {
    path: 'orders',
    component: OrderListComponent,
  },
  {
    path: 'orders/:id',
    component: OrderGetComponent,
  },
  {
    path: 'cart',
    component: CartComponent,
  },
];

@NgModule({ 
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
