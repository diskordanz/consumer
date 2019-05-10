import { Component, OnInit } from '@angular/core';
import { Consumer } from '../models';
import { CommonService } from '../services/common.service';
import { Location } from '@angular/common' 
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  created: boolean;
  user: Consumer = {
    id: 0,
    first_name: "",
    last_name: "",
    password: "",
    phone_number: "",
    city: "",
    address: "",
    login: "",
    mail: ""
  };

  constructor(private userService: CommonService, private router: Router) {
    this.created = false;
   }

  ngOnInit() {
  }

  createUser():void {
    this.userService.createConsumer(this.user).subscribe(user => {
      this.created = true;
      this.user = user;
    });
    this.router.navigate(['login']);
  }
}
