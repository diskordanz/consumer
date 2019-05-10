import { Component, OnInit } from '@angular/core';
import { CartItem } from '../models';
import { CommonService } from '../services/common.service';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {

  constructor(private s: CommonService, private auth: AuthService,private router: Router) { }
  cart: CartItem[];
  
  ngOnInit() {
    var id = this.auth.getUserID(localStorage.getItem('token')) 
    this.s.getCart(id).subscribe((data: CartItem[]) => {
      this.cart = data;
  })
  }

  getProduct(id) {
    this.router.navigateByUrl('/products/'+id)
  }

}
