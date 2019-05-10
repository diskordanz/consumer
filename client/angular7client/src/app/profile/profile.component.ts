
import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { Consumer } from '../models';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
})
export class ProfileComponent implements OnInit {
  user: any = Consumer;
  

  constructor(
    private auth: AuthService,
    private userService: CommonService,
    private router: Router) { }

  ngOnInit() {
    this.getUser();
  }

  getUser(): void {
  var id = this.auth.getUserID(localStorage.getItem('token')) 
  this.userService.getProfile(id).subscribe(user => {
    this.user = user;
  })
  }

}
