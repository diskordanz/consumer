import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { Product, CartItem } from '../models';
import { ActivatedRoute } from '@angular/router';
import {Observable} from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthService } from '../services/auth.service';


@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {
  state$: Observable<object>; 
  products: Product[];

  item: CartItem = {
    id: 0,
    id_consumer: 0,
    product: null,
    count: 1,
  };

  item2: CartItem;

  constructor(public router: ActivatedRoute, private s: CommonService, private auth: AuthService) { }
 
  ngOnInit() {
    this.state$ = this.router.paramMap
    .pipe(map(() => window.history.state))

    if(window.history.state.category_id){

      if(window.history.state.category_id == 1){

      this.s.listProducts(window.history.state.name)
      .subscribe((data: Product[]) => {
        this.products = data;});
      }
      else {
      this.s.listProductsByCategory(window.history.state.category_id, window.history.state.name)
      .subscribe((data: Product[]) => {
        this.products = data;});
      } 
  }
  else{
    this.s.listProducts("")
    .subscribe((data: Product[]) => {
      this.products = data;});
  }
}
  addToCart(product: Product){
    
    this.item.id_consumer = this.auth.getUserID(localStorage.getItem('token')); 
    this.item.product=product;


    this.s.getCartItem(this.item).subscribe((data :CartItem)=> {
      this.item2 = data;})

    if(this.item2){
        this.item2.count++;
        this.s.updateCartItem(this.item2).subscribe(item => {
          this.item2 = item;
        });
        console.log(this.item2)
        console.log("update")
      }    
    else {
      this.s.createCartItem(this.item).subscribe(item => {
      this.item = item;
    });
  }
}
}