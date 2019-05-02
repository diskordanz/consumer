import { Component, OnInit } from '@angular/core';
import { FranchiseService } from '../franchise.service';
import Franchise from '../Franchise';

@Component({
  selector: 'app-franchise-list',
  templateUrl: './franchise-list.component.html',
  styleUrls: ['./franchise-list.component.css']
})
export class FranchiseListComponent implements OnInit {

  franchises: Franchise[];

  constructor(private fr: FranchiseService) { }

  ngOnInit() {
    this.fr
      .listFranchises()
      .subscribe((data: Franchise[]) => {
        this.franchises = data;
    });
  }
} 
  