import { Component, OnInit, Inject, Output, EventEmitter } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { FieldAttr } from '../../../../services/models/search/field-attr';
import { ConditionData } from '../../../../services/models/search/condition-data';
import { ComplexSearchItems } from '../../../../services/models/search/complex-search-items';
import {Group } from '../../../../services/models/group/group';

@Component({
  selector: 'app-complex-search-dialog',
  templateUrl: './complex-search-dialog.component.html',
  styleUrls: ['./complex-search-dialog.component.css']
})
export class ComplexSearchDialogComponent implements OnInit {

  constructor(
    public dialogRef: MatDialogRef<ComplexSearchDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: ComplexSearchItems
  ) { }

  ngOnInit() {
    this.dialogRef.updateSize("1100px")
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

  onSearch(data: ConditionData): void {
    this.dialogRef.close();
  }

}

