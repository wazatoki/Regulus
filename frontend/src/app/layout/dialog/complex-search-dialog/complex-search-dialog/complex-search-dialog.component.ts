import { Component, OnInit, Inject, Output, EventEmitter } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Category } from 'src/app/services/models/search/category';

@Component({
  selector: 'app-complex-search-dialog',
  templateUrl: './complex-search-dialog.component.html',
  styleUrls: ['./complex-search-dialog.component.css']
})
export class ComplexSearchDialogComponent implements OnInit {

  constructor(
    public dialogRef: MatDialogRef<ComplexSearchDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) { }

  ngOnInit() {
    this.dialogRef.updateSize("1100px")
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

}

