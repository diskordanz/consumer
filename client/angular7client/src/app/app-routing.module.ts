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
import { LoginComponent } from './login/login.component';
import { RegistrationComponent } from './registration/registration.component';
import { TokenInterceptorService } from './services/token-interceptor.service';
import { AuthGuard } from './guards/auth.guard';
import { HTTP_INTERCEPTORS } from '@angular/common/http';

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
    path: 'categories/:id/products',
    component: ProductListComponent,
  },
  {
    path: 'products/:id',
    component: ProductGetComponent,
  },
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'orders',
    component: OrderListComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'orders/:id',
    component: OrderGetComponent,
    canActivate: [AuthGuard] 
  },
  {
    path: 'cart',
    component: CartComponent,
  },
  { 
    path: 'login', 
    component: LoginComponent, 
  },
  { 
    path: 'registration', 
    component: RegistrationComponent, 
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(
      routes
    )],
  exports: [RouterModule]/*,
  providers: [AuthGuard,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptorService,
      multi: true
    }
  ],*/
})
export class AppRoutingModule { }
