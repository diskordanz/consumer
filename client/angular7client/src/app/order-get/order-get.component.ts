import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { Order } from '../models';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-order-get',
  templateUrl: './order-get.component.html',
  styleUrls: ['./order-get.component.css']
})
export class OrderGetComponent implements OnInit {

  order: Order;

  constructor(private route: ActivatedRoute, private s: CommonService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getOrder(params['id']).subscribe((res: Order) => {
        this.order = res;});
    });
  }
}