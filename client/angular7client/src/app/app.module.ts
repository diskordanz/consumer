import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { SlimLoadingBarModule } from 'ng2-slim-loading-bar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FranchiseListComponent } from './franchise-list/franchise-list.component';
import { FranchiseService } from './franchise.service';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {AlertModule } from 'ngx-bootstrap';

@NgModule({
  declarations: [
    AppComponent,
    FranchiseListComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    SlimLoadingBarModule,
    FormsModule,
    HttpClientModule,
    AlertModule.forRoot()
  ],
  providers: [FranchiseService],
  bootstrap: [AppComponent]
})
export class AppModule { }
