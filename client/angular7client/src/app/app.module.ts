import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { SlimLoadingBarModule } from 'ng2-slim-loading-bar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';
import { CommonService } from './common.service';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
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
    ProfileComponent
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
  providers: [CommonService],
  bootstrap: [AppComponent]
})
export class AppModule { }
