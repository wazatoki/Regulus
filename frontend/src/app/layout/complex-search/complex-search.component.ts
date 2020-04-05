import { Component, OnInit, Input } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup, AbstractControl } from '@angular/forms';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { Group } from '../../services/models/group/group';
import { FieldAttr } from '../../services/models/search/field-attr';
import { ConditionData } from '../../services/models/search/condition-data';
import { SearchCondition } from '../../services/models/search/search-condition';
import { OrderCondition } from '../../services/models/search/order-condition';
import { SaveData } from '../../services/models/search/save-data';
import { ComplexSearchService } from '../../services/share/complex-search.service';

@Component({
  selector: 'app-complex-search',
  templateUrl: './complex-search.component.html',
  styleUrls: ['./complex-search.component.css']
})
export class ComplexSearchComponent implements OnInit {

  form: FormGroup;

  selectedDisplayItemArray: FieldAttr[];
  fromDisplayItemArray: FieldAttr[];

  @Input() displayItemList: FieldAttr[] = [];
  @Input() searchConditionList: FieldAttr[] = [];
  @Input() orderConditionList: FieldAttr[] = [];
  @Input() isShowDisplayItem: boolean = false;
  @Input() isShowOrderCondition: boolean = false;
  @Input() isShowSaveCondition: boolean = false;
  @Input() groupList: Group[] = [];
  @Input() saveData: SaveData;

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

  constructor(private fb: FormBuilder,
    private complexSearchDataShereService: ComplexSearchService) {

    this.form = this.fb.group({
      searchCondition: this.fb.array([]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({
        patternName: this.fb.control(""),
        isDisclose: this.fb.control(""),
        discloseGroups: this.fb.array([]),
      }),
    });

    this.saveData = this.complexSearchDataShereService.initSaveDataObj();
  }

  ngOnInit() {
    this.groupList.forEach(g => {
      this.discloseGroupFormArray.push(this.fb.control(''));
    })
    this.fromDisplayItemArray = this.displayItemList;
    this.selectedDisplayItemArray = [];
  }

  displayItemDrop(event: CdkDragDrop<FieldAttr[]>) {
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

  removeSearchCondition(i: number) {
    this.searchConditionFormArray.removeAt(i);
  }

  pushOrderCondition() {
    this.orderConditionFormArray.push(new FormGroup({
      orderFieldSelected: new FormControl(''),
      orderFieldKeyWordSelected: new FormControl(''),
    }));
  }

  removeOrderCondition(i: number) {
    this.orderConditionFormArray.removeAt(i);
  }

  clickSave() {
    this.createSaveData();
    this.complexSearchDataShereService.orderComplexSearchSave(this.saveData);
  }

  clickSearch(): void {
    const data: ConditionData = this.createSearchData();
    this.complexSearchDataShereService.orderComplexSearch(data);
  }

  createSearchCondition(): SearchCondition[] {
    const result: SearchCondition[] = [];
    this.searchConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      this.searchConditionList.forEach((v, i) => {
        if (v.id == formGroup.get('fieldSelected').value) {
          field = v;
        }
      });
      const condition: SearchCondition = {
        field: field,
        conditionValue: formGroup.get('conditionValue').value,
        matchType: formGroup.get('matchTypeSelected').value,
        operator: formGroup.get('operatorSelected').value,
      };
      result.push(condition);
    });
    return result;
  }

  createOrderCondition(): OrderCondition[] {
    const result: OrderCondition[] = [];
    this.orderConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      this.orderConditionList.forEach((v, i) => {
        if (v.id == formGroup.get('orderFieldSelected').value) {
          field = v;
        }
      });
      const condition: OrderCondition = {
        orderField: field,
        orderFieldKeyWord: formGroup.get('orderFieldKeyWordSelected').value,
      }
      result.push(condition);
    });
    return result;
  }

  createSaveData(): void {

    if (this.saveData === null || this.saveData === undefined) {
      this.saveData = this.complexSearchDataShereService.initSaveDataObj();
    }

    if (this.isShowSaveCondition) {
      this.saveData.patternName = this.saveConditions.get('patternName').value;
      this.saveData.isDisclose = this.saveConditions.get('isDisclose').value;
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        if (v.value === true) {
          this.saveData.discloseGroups.push(this.groupList[i].id);
        }
      });
    }

    this.saveData.conditionData = this.createSearchData();

  }

  createSearchData(): ConditionData {
    const data: ConditionData = this.complexSearchDataShereService.initConditionDataObj();

    if (this.isShowDisplayItem) {
      data.displayItemList = this.selectedDisplayItemArray;
    }

    data.searchConditionList = this.createSearchCondition();

    if (this.isShowOrderCondition) {
      data.orderConditionList = this.createOrderCondition();
    }

    return data;
  }

}// end of class
