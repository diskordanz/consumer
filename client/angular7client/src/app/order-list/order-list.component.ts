import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';import { Order } from '../models';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-order-list',
  templateUrl: './order-list.component.html',
  styleUrls: ['./order-list.component.css']
})
export class OrderListComponent implements OnInit {

  orders: Order[];

  constructor(private s: CommonService, private auth: AuthService,private router: Router) { }
 
  ngOnInit() {
    var id = this.auth.getUserID(localStorage.getItem('token')) 
    this.s
      .listOrders(id)
      .subscribe((data: Order[]) => {
        this.orders = data;
    });
  }
}
