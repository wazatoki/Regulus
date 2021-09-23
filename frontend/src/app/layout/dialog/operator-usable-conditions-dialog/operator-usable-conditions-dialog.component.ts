import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { SaveData } from 'src/app/services/models/search/save-data';

@Component({
  selector: 'app-operator-usable-conditions-dialog',
  templateUrl: './operator-usable-conditions-dialog.component.html',
  styleUrls: ['./operator-usable-conditions-dialog.component.css']
})
export class OperatorUsableConditionsDialogComponent implements OnInit {

  conditionSelected(data: SaveData) {
    this.dialogRef.close(data);
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

  constructor(
    public dialogRef: MatDialogRef<OperatorUsableConditionsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) { }

  ngOnInit() {
  }
}

export interface DialogData {
  title: string,
  operatorUsableConditions: SaveData[]
}
