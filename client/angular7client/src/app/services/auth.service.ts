
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private URL = 'http://localhost:3000';
  constructor(private http: HttpClient) {
  }

  loginUser(user) {
    return this.http.post(`${this.URL}/login`, user);
  }

  loggedIn() {
    return !!localStorage.getItem('token');
  }

  getToken() {
    return localStorage.getItem('token');
  }

  restorePassword(data) {
    return this.http.post(`${this.URL}/forgot`, data);
  }

  isAdmin() {
    return this.http.get(`${this.URL}/admin`);
  }

}