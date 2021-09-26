import { Component, OnInit, Inject, Output, EventEmitter } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ComplexSearchItems } from '../../../../services/models/search/complex-search-items';

@Component({
  selector: 'app-complex-search-dialog',
  templateUrl: './complex-search-dialog.component.html',
  styleUrls: ['./complex-search-dialog.component.css']
})
export class ComplexSearchDialogComponent implements OnInit {

  constructor(
    public dialogRef: MatDialogRef<ComplexSearchDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) { }

  ngOnInit() {
    this.dialogRef.updateSize('1100px');
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

}

export interface DialogData {
  title: string;
  complexSearchItems: ComplexSearchItems;
  saveData: SaveData;
}
