import { Component, OnInit, Input, Inject } from '@angular/core';
import { Maker } from '../../services/models/maker/maker';
import { FormBuilder, FormGroup, AbstractControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-maker-input-form',
  templateUrl: './maker-input-form.component.html',
  styleUrls: ['./maker-input-form.component.css']
})
export class MakerInputFormComponent implements OnInit {

  constructor( private fb: FormBuilder) { }

  ngOnInit() {
  }

  makerForm: FormGroup = this.fb.group({
    name: ['', [Validators.required]],
  });

  get name(): AbstractControl {
    return this.makerForm.get('name');
  }

  
}
