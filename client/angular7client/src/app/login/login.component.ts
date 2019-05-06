import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { AuthService } from '../services/auth.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  private loginForm;
  loginUserData;
  
  constructor(private fb: FormBuilder,
    private authService: AuthService,
    private router: Router) { }

  ngOnInit() {
    this.loginForm = this.fb.group({
      name: ['', [Validators.required]],
      password: ['', [Validators.required]]
    });
  }

  onLogin() {
    this.loginUserData = {
      name: this.loginForm.value.name,
      password: this.loginForm.value.password
    };
    this.authService.loginUser(this.loginUserData)
      .subscribe(res => {
        localStorage.setItem('token', res['token']);
        this.router.navigate(['/profile/view']);
      },
        err => {
          if (err instanceof HttpErrorResponse) {
            if (err.status === 401) {
            }
          }
        });
  }

  logout() {
    localStorage.removeItem('token');
    this.router.navigate(['/login']);
  }

  isLoggedIn(): boolean {
    return this.authService.loggedIn();
  }
}