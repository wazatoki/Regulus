import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup, AbstractControl } from '@angular/forms';
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
  fromDisplayItemArray: fieldAttr[];

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
    this.groupList.forEach(g => {
      this.discloseGroupFormArray.push(this.fb.control(''));
    })
    this.fromDisplayItemArray = this.displayItemList;
    this.selectedDisplayItemArray = [];
  }

  displayItemDrop(event: CdkDragDrop<fieldAttr[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex);
    }
  }

  controlDrop(event: CdkDragDrop<AbstractControl[]>) {
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
    const data: saveData = this.createSaveData();
    this.onSave.emit(data)
  }

  initSaveDataObj(): saveData {
    return {
      patternName: '',
      category: '',
      isDisclose: false,
      discloseGroups: [],
      ownerID: '',
      conditionData: this.initConditionDataObj(),
    };
  }

  initConditionDataObj() {
    return {
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    }
  }

  createSearchCondition(): searchCondition[] {
    const result: searchCondition[] = [];
    this.searchConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: fieldAttr;
      this.searchConditionList.forEach((v, i) => {
        if (v.id == formGroup.get('fieldSelected').value) {
          field = v;
        }
      });
      const condition: searchCondition = {
        field: field,
        conditionValue: formGroup.get('conditionValue').value,
        matchType: formGroup.get('matchTypeSelected').value,
        operator: formGroup.get('operatorSelected').value,
      };
      result.push(condition);
    });
    return result;
  }

  createOrderCondition(): orderCondition[] {
    const result: orderCondition[] = [];
    this.orderConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: fieldAttr;
      this.orderConditionList.forEach((v, i) => {
        if (v.id == formGroup.get('orderFieldSelected').value) {
          field = v;
        }
      });
      const condition: orderCondition = {
        orderField: field,
        orderFieldKeyWord: formGroup.get('orderFieldKeyWordSelected').value,
      }
      result.push(condition);
    });
    return result;
  }

  createSaveData(): saveData {

    const data: saveData = this.initSaveDataObj();

    if (this.isShowSaveCondition) {
      data.patternName = this.saveConditions.get('patternName').value;
      data.isDisclose = this.saveConditions.get('isDisclose').value;
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        if (v.value === true) {
          data.discloseGroups.push(this.groupList[i].id);
        }
      });
    }

    if (this.isShowDisplayItem) {
      data.conditionData.displayItemList = this.selectedDisplayItemArray;
    }

    data.conditionData.searchConditionList = this.createSearchCondition();

    if (this.isShowOrderCondition) {
      data.conditionData.orderConditionList = this.createOrderCondition();
    }

    return data;
  }

}// end of class

export interface fieldAttr {
  id: string,
  entityName: string,
  fieldName: string,
  viewValue: string,
  fieldType: string,
}

export interface saveData {
  patternName: string;
  category: string;
  isDisclose: boolean;
  discloseGroups: string[];
  ownerID: string;
  conditionData: conditionData;
}

interface conditionData {
  displayItemList: fieldAttr[],
  searchConditionList: searchCondition[],
  orderConditionList: orderCondition[],
}

interface searchCondition {
  field: fieldAttr,
  conditionValue: string,
  matchType: string,
  operator: string,
}

interface orderCondition {
  orderField: fieldAttr,
  orderFieldKeyWord: string,
}


