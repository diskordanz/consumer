import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { Product } from '../models';
import { ActivatedRoute } from '@angular/router';
import {Observable} from 'rxjs';
import { map } from 'rxjs/operators';


@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {
  state$: Observable<object>; 
  products: Product[];

  constructor(public router: ActivatedRoute, private s: CommonService) { }
 
  ngOnInit() {
    this.state$ = this.router.paramMap
    .pipe(map(() => window.history.state))

    console.log(window.history.state.name)
    this.s.listProducts(window.history.state.name)
    .subscribe((data: Product[]) => {
      this.products = data;
  });
  }
   
}