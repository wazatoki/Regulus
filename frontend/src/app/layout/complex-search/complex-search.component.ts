import { Component, OnInit, Input } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';
import {CdkDragDrop, moveItemInArray, transferArrayItem} from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-complex-search',
  templateUrl: './complex-search.component.html',
  styleUrls: ['./complex-search.component.css']
})
export class ComplexSearchComponent implements OnInit {

  form: FormGroup;

  selectedDisplayItemArray: fieldAttr[];

  @Input() displayItemList: fieldAttr[];
  @Input() searchConditionList: fieldAttr[];
  @Input() orderConditionList: fieldAttr[];

  get searchConditionArray() {
    return this.form.get('searchCondition') as FormArray;
  }

  get orderConditionArray() {
    return this.form.get('orderCondition') as FormArray;
  }

  get saveCondition() {
    return this.form.get('saveCondition') as FormGroup;
  }

  constructor(private fb: FormBuilder) { 
    this.form = this.fb.group({
      searchCondition: this.fb.array([new FormGroup({
        fieldSelected: new FormControl(''),
        conditionValue: new FormControl(''),
        matchTypeSelected: new FormControl(''),
        operatorSelected: new FormControl(''),
      })]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({}),
    });
  }

  ngOnInit() {
  }

  drop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(event.previousContainer.data,
                        event.container.data,
                        event.previousIndex,
                        event.currentIndex);
    }
  }



}

export interface fieldAttr {
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}
