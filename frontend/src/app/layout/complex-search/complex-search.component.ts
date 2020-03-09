import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { Group } from '../../services/models/group/group';

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
  @Input() groupList: Group[];
  @Output() onSave = new EventEmitter<saveData>();

  get searchConditionFormArray() {
    return this.form.get('searchCondition') as FormArray;
  }

  get orderConditionFormArray() {
    return this.form.get('orderCondition') as FormArray;
  }

  get saveConditions() {
    return this.form.get('saveCondition') as FormGroup;
  }

  get discloseGroupFormArray() {
    return this.saveConditions.get('discloseGroups') as FormArray;
  }

  constructor(private fb: FormBuilder) {
    this.form = this.fb.group({
      searchCondition: this.fb.array([]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({
        patternName: this.fb.control(""),
        isDisclose: this.fb.control(""),
        discloseGroups: this.fb.array([]),
      }),
    });
  }

  ngOnInit() {
    this.groupList.forEach( g => {
      this.discloseGroupFormArray.push(this.fb.control(''));
    })
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
    this.searchConditionFormArray.push(new FormGroup({
      fieldSelected: new FormControl(''),
      conditionValue: new FormControl(''),
      matchTypeSelected: new FormControl(''),
      operatorSelected: new FormControl(''),
    }));
  }

  pushOrderCondition() {
    this.orderConditionFormArray.push(new FormGroup({
      orderFieldSelected: new FormControl(''),
      orderFieldKeyWordSelected: new FormControl(''),
    }));
  }

  clickSave() {
    let data: saveData;
    data.patternName = this.saveConditions.get('patternName').value;
    data.isDisclose = this.saveConditions.get('isDisclose').value;

    this.onSave.emit(data)
  }

}

export interface fieldAttr {
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}

interface saveData {
  patternName: string;
  category: string;
  isDisclose: boolean;
  discloseGroups: string[];
  ownerID: string;
  conditionData: conditionData;
}

interface conditionData {

}
