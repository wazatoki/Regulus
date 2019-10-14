import { Component, OnInit } from '@angular/core';

export interface MakerElement {
  id: string;
  name: string;
}

@Component({
  selector: 'app-maker-master',
  templateUrl: './maker-master.component.html',
  styleUrls: ['./maker-master.component.css']
})
export class MakerMasterComponent implements OnInit {

  displayedColumns: string[];
  dataSource: MakerElement[];

  constructor() {
    this.displayedColumns = ['name'];
    this.dataSource = [];
  }


  ngOnInit() {
  }

}
