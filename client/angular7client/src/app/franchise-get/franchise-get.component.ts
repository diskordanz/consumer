import { Component, OnInit } from '@angular/core';
import { CommonService } from '../services/common.service';
import {Franchise, Location} from '../models';
import { ActivatedRoute } from '@angular/router';

const pathToImage = "../../assets/img/"

@Component({
  selector: 'app-franchise-get',
  templateUrl: './franchise-get.component.html',
  styleUrls: ['./franchise-get.component.css']
})
export class FranchiseGetComponent implements OnInit {

  public pathToImage = pathToImage;
  franchise: Franchise;
  locations: Location[];

  constructor(private route: ActivatedRoute, private s: CommonService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getFranchise(params['id']).subscribe((res: Franchise) => {
        this.franchise = res;})
        this.s.listLocations(params['id'], 20, 0).subscribe((data: Location[]) => {
        this.locations = data;});
    });
  }
}

      