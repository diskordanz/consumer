import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { OrderItemWithProduct, OrderItem, Order , OrderWithItems, OrderWithItemsAndProducts, Product} from '../models';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-order-get',
  templateUrl: './order-get.component.html',
  styleUrls: ['./order-get.component.css']
})
export class OrderGetComponent implements OnInit {

  order: OrderWithItemsAndProducts;
  constructor(private route: ActivatedRoute, private s: CommonService, private router: Router) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getOrder(params['id']).subscribe((res: OrderWithItemsAndProducts) => {
        this.order = res});});
  }

  getProduct(id) {
    this.router.navigateByUrl('/products/'+id)
  }
}