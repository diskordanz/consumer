import { TokenInterceptorService } from './services/token-interceptor.service';
import { AuthGuard } from './guards/auth.guard';
 
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { SlimLoadingBarModule } from 'ng2-slim-loading-bar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';
import { CommonService } from './services/common.service';
import { FormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS, HttpClient } from '@angular/common/http';
import {AlertModule } from 'ngx-bootstrap';
import { FranchiseGetComponent } from './franchise-get/franchise-get.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatSelectModule} from '@angular/material/select';
import { CartComponent } from './cart/cart.component';
import { OrderGetComponent } from './order-get/order-get.component';
import { OrderListComponent } from './order-list/order-list.component';
import { ProductListComponent } from './product-list/product-list.component';
import { ProductGetComponent } from './product-get/product-get.component';

import { ProfileComponent } from './profile/profile.component';
import { HeaderComponent } from './header/header.component';
import { LoginComponent } from './login/login.component';
import { RegistrationComponent } from './registration/registration.component';


@NgModule({
  declarations: [
    AppComponent,
    FranchiseListComponent,
    FranchiseGetComponent,
    CartComponent,
    OrderGetComponent,
    OrderListComponent,
    ProductListComponent,
    ProductGetComponent,
    ProfileComponent,
    HeaderComponent,
    LoginComponent,
    RegistrationComponent,
    ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    SlimLoadingBarModule,
    FormsModule,
    HttpClientModule,
    AlertModule.forRoot(),
    BrowserAnimationsModule,
    MatSelectModule
  ],
  providers: [AuthGuard,/*
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptorService,
      multi: true
    },*/
    CommonService],
  bootstrap: [AppComponent]
})
export class AppModule { }
