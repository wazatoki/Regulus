import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { TRUE, FALSE } from '../../../services/models/enum/boolean';

@Component({
  selector: 'app-notice-dialog',
  templateUrl: './notice-dialog.component.html',
  styleUrls: ['./notice-dialog.component.css']
})
export class NoticeDialogComponent implements OnInit {

  RESULT_TRUE = TRUE;

  constructor(
    public dialogRef: MatDialogRef<NoticeDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) { }

  ngOnInit() {
  }
}

export interface DialogData {
  title: string;
  contents: string;
}
