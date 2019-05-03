import { Component, OnInit } from '@angular/core';
import { LocationService } from '../location.service';
import Location from '../location';
import { ActivatedRoute, Router } from '@angular/router';


@Component({
  selector: 'app-location-list',
  templateUrl: './location-list.component.html',
  styleUrls: ['./location-list.component.css']
})

export class LocationListComponent implements OnInit {

  locations: Location[];
  constructor(private route: ActivatedRoute, private lc: LocationService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.lc.listLocations(params['id']).subscribe((data: Location[]) => {
        this.locations = data;
      });
    });
  }
}