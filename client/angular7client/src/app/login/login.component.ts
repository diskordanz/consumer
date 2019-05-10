import { Router } from '@angular/router';
import { Component, OnInit, Input } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { HttpErrorResponse } from '@angular/common/http';
import { LoginUser } from '../models';

@Component({
  selector: 'app-login', 
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  @Input() user: LoginUser; 
  
  constructor(
    private authService: AuthService,
    private router: Router) { }

  ngOnInit() {
    this.user = {
      id: 0,
      password: '',
      token: '',
      login:''
    };
  }

  onLogin() {
    this.authService.login(this.user)
  }
  
}