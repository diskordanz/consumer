import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import { Franchise } from '../models';

const pathToImage = "../../assets/img/"

@Component({
  selector: 'app-franchise-list',
  templateUrl: './franchise-list.component.html',
  styleUrls: ['./franchise-list.component.css']
})

export class FranchiseListComponent implements OnInit {

  franchises: Franchise[];
  public pathToImage = pathToImage;

  constructor(private s: CommonService) { }
 
  ngOnInit() {
    this.s
      .listFranchises()
      .subscribe((data: Franchise[]) => {
        this.franchises = data;
    });
  }
} 
  