import { Component, OnInit } from '@angular/core';
import { CommonService } from '../common.service';
import { Product } from '../models';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {

  products: Product[];

  constructor(private s: CommonService) { }
 
  ngOnInit() {
    this.s
      .listProducts()
      .subscribe((data: Product[]) => {
        this.products = data;
    });
  }
}