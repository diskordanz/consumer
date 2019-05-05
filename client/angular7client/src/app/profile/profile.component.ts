import { Component, OnInit } from '@angular/core';
import { CommonService } from '../common.service';
import { Consumer } from '../models';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  consumer: Consumer;

  constructor(private route: ActivatedRoute, private s: CommonService) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.s.getProfile().subscribe((res: Consumer) => {
        this.consumer = res;});
    });
  }
}
