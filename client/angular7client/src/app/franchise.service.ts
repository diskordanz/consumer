import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap, map } from 'rxjs/operators';
import  Franchise  from './Franchise';

@Injectable({
  providedIn: 'root'
})
export class FranchiseService {

  uri = 'http://localhost:8080/api/franchises?count=10&offset=0';

  constructor(private http: HttpClient) { }
  
  listFranchises() {
    return this
           .http
           .get(`${this.uri}`);
  }
}
