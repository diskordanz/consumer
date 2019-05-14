import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import {Product, CartItem} from '../models';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from '../services/auth.service';

const pathToImage = "../../assets/img/"

@Component({
  selector: 'app-product-get',
  templateUrl: './product-get.component.html',
  styleUrls: ['./product-get.component.css']
})
export class ProductGetComponent implements OnInit {

  public pathToImage = pathToImage;
  product: Product;
  item: CartItem;
  item2: CartItem;

  constructor(private route: ActivatedRoute, private s: CommonService,  private auth: AuthService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getProduct(params['id']).subscribe((res: Product) => {
        this.product = res;});
    });
  }

  addToCart(product: Product){

    this.item = {
        id: 0,
        consumer_id: this.auth.getUserID(localStorage.getItem('token')),
        product: product,
        count: 1,
      };

      this.s.getCartItem(this.item).subscribe((data: CartItem) => {
        this.item2 = data;
        console.log(this.item2);
      
      if(this.item2.id == 0){
        this.s.createCartItem(this.item).subscribe((data :CartItem) => {
          this.item = data;
          console.log("create");
        });
      } 
      else {
        this.item2.count++;
        this.s.updateCartItem(this.item2).subscribe((data :CartItem) => {
          this.item2 = data;
        console.log("update"); 
      });
    }
    });
  } 
}