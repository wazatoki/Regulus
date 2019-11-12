import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { MakerService } from '../../services/api/maker.service';
import { Maker } from '../../services/models/maker/maker';
import { MakerCondition } from '../../services/models/maker/maker-condition';
import { MakerInputFormComponent } from '../maker-input-form/maker-input-form.component';

@Component({
  selector: 'app-maker-master',
  templateUrl: './maker-master.component.html',
  styleUrls: ['./maker-master.component.css']
})
export class MakerMasterComponent implements OnInit {

  displayedColumns: string[];
  dataSource: Maker[];

  name: string;

  constructor(
    private makerService:MakerService,
    private makerCondition: MakerCondition,
    private dialog: MatDialog) {
    this.displayedColumns = ['name'];
    this.dataSource = [];
  }

  ngOnInit() {
  }

  onFetchedMakers(data: Maker[]){
    this.dataSource = data;
  }

  openInputForm(): void {
    const dialogRef = this.dialog.open(MakerInputFormComponent, {
      width: '500px',
      data: {name: this.name}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.name = '';
    });
  }
}
