import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { TRUE, FALSE } from '../../../services/models/enum/boolean';

@Component({
  selector: 'app-alert-diarog',
  templateUrl: './alert-diarog.component.html',
  styleUrls: ['./alert-diarog.component.css']
})
export class AlertDiarogComponent implements OnInit {

  RESULT_TRUE = TRUE;
  RERULT_FALSE = FALSE;

  constructor(
    public dialogRef: MatDialogRef<AlertDiarogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) { }

  ngOnInit() {
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }
}

export interface DialogData {
  title: string;
  contents: string;
}
