import { Component, OnInit, Injectable } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { MatBottomSheetRef } from '@angular/material/bottom-sheet';

@Component({
  selector: 'app-maker-input-form',
  templateUrl: './maker-input-form.component.html',
  styleUrls: ['./maker-input-form.component.css']
})
export class MakerInputFormComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private _bottomSheetRef: MatBottomSheetRef<MakerInputFormComponent>) { }

  ngOnInit() {
  }

  makerForm: FormGroup = this.fb.group({
    name: ['', [Validators.required]],
  });

  get name(): AbstractControl {
    return this.makerForm.get('name');
  }

  onCancelClick() {

  }
}
