
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
//import { validateName } from '../validators/name.validator';
//import { validateAge } from '../validators/age.validator';
//import { validateDate } from '../validators/date.validator';
import { ConsumerService } from '../services/user.service';
import { Consumer } from '../models';
//import * as moment from 'moment';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
})
export class ProfileComponent implements OnInit {
  user: Consumer;
  public loading = false;
  private editForm: FormGroup;

  constructor(private fb: FormBuilder,
    private userService: ConsumerService,
    private router: Router) { }

  ngOnInit() {
    this.getUserInfo();
    this.initForm();
  }

  resetForm() {
    this.editForm.reset();
  }

  onSubmit() {
    if (this.editForm.valid) {
      this.userService.updateConsumer(
        { 
          id: this.user.id,
          first_name: this.editForm.value.first_name,
          last_name: this.editForm.value.last_name,
          phone_number: this.editForm.value.phone_number,
          city: this.editForm.value.city,
          address: this.editForm.value.address,
          cart: this.user.cart,
          login: this.editForm.value.login,
          mail: this.editForm.value.mail,
          password: this.user.password,
        }
      ).subscribe(() => {
        this.router.navigate(['/profile']);
      });
    }
  }

  initForm() {
    this.editForm = this.fb.group({
      first_name: ['', [Validators.required]],
      last_name: ['', [Validators.required]],
      phone_number: ['', [Validators.required]],
      city: ['', [Validators.required]],
      address: ['', [Validators.required]],
      login: ['', [Validators.required]],
      mail: ['', [Validators.required]],

    });
  }

  setFormValue() {
    this.editForm.patchValue({
      first_name: this.user.first_name,
      last_name: this.user.last_name,
      phone_number: this.user.phone_number,
      city: this.user.city,
      address: this.user.address,
      login: this.user.login,
      mail: this.user.mail
    });
  }

  getUserInfo() {
    this.loading = true;
    this.userService.getAuthorizedConsumer()
      .subscribe((data: Consumer) => {
        this.user = data;
        this.setFormValue();
        this.loading = false;
      });
  }

}
