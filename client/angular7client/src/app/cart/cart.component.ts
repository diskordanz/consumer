import { Component, OnInit } from '@angular/core';
import { CartItem, CartItemByFranchise, Order, OrderItem } from '../models';
import { CommonService } from '../services/common.service';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';
import { summaryFileName } from '@angular/compiler/src/aot/util';
import { Timestamp } from 'rxjs/internal/operators/timestamp';
import { containsElement } from '@angular/animations/browser/src/render/shared';
const pathToImage = "../../assets/img/"

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {
  public pathToImage = pathToImage;

  constructor(private s: CommonService, private auth: AuthService,private router: Router) { }
  cart: CartItemByFranchise[];
  order: Order;
  order_item: OrderItem;

  ngOnInit() {
    var id = this.auth.getUserID(localStorage.getItem('token')) 
    this.s.getCart(id, 20, 0).subscribe((data: CartItemByFranchise[]) => {
      this.cart = data;
  })
  }

  getProduct(id) {
    this.router.navigateByUrl('/products/'+id)
  }

  deleteCartItem(id){
    this.s.deleteCartItem(id).subscribe(() => this.ngOnInit())
}

addItem(item : CartItem){
  item.count++;
  this.s.updateCartItem(item).subscribe(() => this.ngOnInit())
}

delItem(item : CartItem){
  item.count--;
  if(item.count!=0){
    this.s.updateCartItem(item).subscribe(() => this.ngOnInit())
  }
  else{
    this.deleteCartItem(item.id)
  }
}
createOrder(item:CartItemByFranchise){
  this.order = {
  id: 0,
  consumer_id: this.auth.getUserID(localStorage.getItem('token')),
  franchise_id: item.franchise_id,
  time: new Date(),
  total: this.total(item.cart_items),
  status: 'NEW',
  }
  this.s.createOrder(this.order).subscribe((data: Order) => {
    this.order = data;
  
 /* this.s.findOrder(this.order).subscribe((data: Order) => {
    this.order = data;
  });*/
    console.log(this.order);
    item.cart_items.forEach(element => {
      this.order_item = {
        id: 0,
        product_id: element.product.id,
        order_id: this.order.id,
        count: element.count,
      }
      this.s.createOrderItem(this.order_item).subscribe(()=>{this.deleteCartItem(element.id)});
    });
  });
}

total(items: CartItem[]) : number{
  var sum = 0;
  items.forEach(element => {
  sum=sum+element.product.price*element.count;
});
return sum;
}

}
