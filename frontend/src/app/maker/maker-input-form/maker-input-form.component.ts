import { Component, OnInit, Input, Inject } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-maker-input-form',
  templateUrl: './maker-input-form.component.html',
  styleUrls: ['./maker-input-form.component.css']
})
export class MakerInputFormComponent implements OnInit {

  form: FormGroup;
  maker: Maker;

  constructor(
    private fb: FormBuilder,
    private dialogRef: MatDialogRef<MakerInputFormComponent>,
    @Inject(MAT_DIALOG_DATA) data
  ) {
    this.maker = data.maker;
   }

  ngOnInit() {
    this.form = this.fb.group(this.maker)
  }

  onCancelClick(){
    this.dialogRef.close(this.maker)
  }

  onSaveClick(){
    this.dialogRef.close();
  }

}
