import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap, map } from 'rxjs/operators';
import  Location  from './location';

@Injectable({
  providedIn: 'root'
})
export class LocationService {

  uri = 'http://localhost:8080/api/franchises';

  constructor(private http: HttpClient) { }
  
  listLocations(id: number) {
    return this
           .http
           .get(`${this.uri}/${id}/locations?count=10&offset=0`);
  }
}
