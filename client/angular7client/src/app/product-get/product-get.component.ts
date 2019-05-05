import { Component, OnInit } from '@angular/core';
import { CommonService } from '../common.service';
import {Product} from '../models';
import { ActivatedRoute } from '@angular/router';

const pathToImage = "../../assets/img/"

@Component({
  selector: 'app-product-get',
  templateUrl: './product-get.component.html',
  styleUrls: ['./product-get.component.css']
})
export class ProductGetComponent implements OnInit {

  public pathToImage = pathToImage;
  product: Product;

  constructor(private route: ActivatedRoute, private s: CommonService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getProduct(params['id']).subscribe((res: Product) => {
        this.product = res;});
    });
  }
}