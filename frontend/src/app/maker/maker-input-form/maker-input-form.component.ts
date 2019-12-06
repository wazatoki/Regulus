import { Component, OnInit, Inject } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

export interface DialogData {
  name: string;
}

@Component({
  selector: 'app-maker-input-form',
  templateUrl: './maker-input-form.component.html',
  styleUrls: ['./maker-input-form.component.css']
})
export class MakerInputFormComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<MakerInputFormComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) { }

  ngOnInit() {
  }

  makerForm: FormGroup = this.fb.group({
    name: ['', [Validators.required]],
  });

  get name(): AbstractControl {
    return this.makerForm.get('name');
  }

  onCancelClick() {
    this.dialogRef.close();
  }

  onClearClick(){}

  onSaveClick(){}
}
