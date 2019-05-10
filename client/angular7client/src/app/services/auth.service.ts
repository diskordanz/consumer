
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { LoginUser } from '../models';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Router } from '@angular/router';
import { Observable, of } from 'rxjs';

import * as jwt_decode from 'jwt-decode'

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private URL = 'http://localhost:8081';
  constructor(private http: HttpClient, private router: Router) {
  }

  login(user: LoginUser):void{
    var url = this.URL + '/api/consumer/auth'
    console.log(url)    
    this.http.post<any>(url, user, {observe: 'response'}).subscribe(resp =>{  
      localStorage.setItem('token', resp.headers.get('Token'));
    }); 
    this.router.navigate(["profile"])     
  }
 
  isLogged():boolean{
    return localStorage.getItem('token') == ""
  }

  logout(): void{
    localStorage.removeItem('token')
  }  

  getTokenExpDate(token: string):Date{
    const decoded = jwt_decode(token);

    if(decoded.exp === undefined) return null;

    const date = new Date(0)
    date.setUTCSeconds(decoded.exp);
    return date;
  }
  getUserID(token : string): number{
    const decoded = jwt_decode(token);
    var userId= decoded['user_id']
    if(userId === undefined) return null;    
    return userId;
  }

  refresh():Observable<boolean>{
    var url = this.URL + '/api/consumer/refresh'
    console.log(url) 
    let body = {
      token: localStorage.getItem('token')
    } 
    console.log(body)
    this.http.post<any>(url, body, {observe: 'response'}).subscribe(resp =>{  
      localStorage.setItem('token', resp.headers.get('Token'));         
    });
    return of(true);
  }

  isTokenExpired(token: string): boolean{
    if (!token) return true

    const date = this.getTokenExpDate(token)

    if(date === undefined) return false
    return !(date.valueOf()> new Date().valueOf());
  }

  public isAuthenticated(): boolean {
    const token = localStorage.getItem('token');
    const helper = new JwtHelperService()
    console.log(token)
    if (token == null){
      return false
    }
    return !helper.isTokenExpired(token);
  }
}