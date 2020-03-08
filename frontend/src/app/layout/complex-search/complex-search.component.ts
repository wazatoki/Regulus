import { Component, OnInit, Input } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';

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
  @Input() isShowDisplayItem: boolean = false;
  @Input() isShowOrderCondition: boolean = false;
  @Input() isShowSaveCondition: boolean = false;

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
      searchCondition: this.fb.array([]),
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

  pushSearchCondition() {
    this.searchConditionArray.push(new FormGroup({
      fieldSelected: new FormControl(''),
      conditionValue: new FormControl(''),
      matchTypeSelected: new FormControl(''),
      operatorSelected: new FormControl(''),
    }));
  }

  pushOrderCondition() {
    this.orderConditionArray.push(new FormGroup({
      orderFieldSelected: new FormControl(''),
      orderFieldKeyWordSelected: new FormControl(''),
    }));
  }

}

export interface fieldAttr {
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}
