import { Component, OnInit } from '@angular/core';
import { CommonService } from '../common.service';
import { Order } from '../models';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-order-list',
  templateUrl: './order-list.component.html',
  styleUrls: ['./order-list.component.css']
})
export class OrderListComponent implements OnInit {

  orders: Order[];

  constructor(private s: CommonService) { }
 
  ngOnInit() {
    this.s
      .listOrders()
      .subscribe((data: Order[]) => {
        this.orders = data;
    });
  }
}
