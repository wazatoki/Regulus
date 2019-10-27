import { Component, OnInit } from '@angular/core';
import { MakerService } from '../../services/api/maker.service';
import { Maker } from '../../services/models/maker/maker';
import { MakerCondition } from '../../services/models/maker/maker-condition';

@Component({
  selector: 'app-maker-master',
  templateUrl: './maker-master.component.html',
  styleUrls: ['./maker-master.component.css']
})
export class MakerMasterComponent implements OnInit {

  displayedColumns: string[];
  dataSource: Maker[];

  constructor(
    private makerService:MakerService,
    private makerCondition: MakerCondition) {
    this.displayedColumns = ['name'];
    this.dataSource = [];
  }

  ngOnInit() {
  }

  onFetchedMakers(data: Maker[]){
    this.dataSource = data;
  }
}
