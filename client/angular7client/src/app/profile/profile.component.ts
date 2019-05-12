
import { Component, OnInit, Input } from '@angular/core';
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
  @Input() user: Consumer;
  updated = false;

  constructor(
    private auth: AuthService,
    private userService: CommonService,
    private router: Router) { }

  ngOnInit() {
    this.getUser();
  }

  updateUser(){
    this.userService.updateConsumer(this.user).subscribe(resp => {
      this.updated = true;
      this.ngOnInit();
    })
  }

  getUser(): void {
  var id = this.auth.getUserID(localStorage.getItem('token')) 
  this.userService.getProfile(id).subscribe((user:Consumer) => {
    this.user = user;
  })
  }

}
