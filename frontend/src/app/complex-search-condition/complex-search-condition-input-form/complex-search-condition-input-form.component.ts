import { Component, OnInit, Input } from '@angular/core';
import { FormGroup, FormBuilder, FormArray, FormControl } from '@angular/forms';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { MatDialog } from '@angular/material/dialog';
import { FieldAttr } from 'src/app/services/models/search/field-attr';
import { SaveData } from 'src/app/services/models/search/save-data';
import { Group } from 'src/app/services/models/group/group';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { SearchCondition } from 'src/app/services/models/search/search-condition';
import { OrderCondition } from 'src/app/services/models/search/order-condition';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { Category } from 'src/app/services/models/search/category';

@Component({
  selector: 'app-complex-search-condition-input-form',
  templateUrl: './complex-search-condition-input-form.component.html',
  styleUrls: ['./complex-search-condition-input-form.component.css']
})
export class ComplexSearchConditionInputFormComponent implements OnInit {

  
  selectedDisplayItemArray: FieldAttr[];
  fromDisplayItemArray: FieldAttr[];
  isShowDisplayItem: boolean = false;
  isShowOrderCondition: boolean = false;
  isShowSaveCondition: boolean = true;

  form: FormGroup;


  get categorySelected() {
    return this.saveConditions.get('category') as FormControl;
  }

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

  @Input() categories: Category[];
  @Input() displayItemList: FieldAttr[] = [];
  @Input() searchConditionList: FieldAttr[] = [];
  @Input() orderConditionList: FieldAttr[] = [];
  @Input() groupList: Group[] = [];
  @Input() saveData: SaveData = this.complexSearchDataShereService.initSaveDataObj();

  constructor(
    private fb: FormBuilder,
    private dialog: MatDialog,
    private complexSearchDataShereService: ComplexSearchService
  ) {
    this.form = this.fb.group({
      searchCondition: this.fb.array([]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({
        category: this.fb.control(''),
        patternName: this.fb.control(''),
        isDisclose: this.fb.control(''),
        discloseGroups: this.fb.array([]),
      }),
    });
  }

  ngOnInit() {
    // saveDataの編集のときは値をフォームに反映する
    if (this.saveData !== null && this.saveData !== undefined && this.saveData.id !== '') {
      this.setSavedDataToForm()
    }
  }

  onSelectCategory() {
    const category = this.categories.find((c) => {
      return (c.name === this.categorySelected.value)
    })

    this.isShowDisplayItem = category.isShowDisplayItem
    this.isShowOrderCondition = category.isShowOrderCondition
  }

  setSavedDataToForm() {
    this.saveConditions.get('patternName').setValue(this.saveData.patternName)
    this.saveConditions.get('isDisclose').setValue(this.saveData.isDisclose)
    this.discloseGroupFormArray.controls.forEach((v, i) => {
      this.saveData.discloseGroupIDs.forEach(id => {
        if (this.groupList[i].id === id) {
          v.setValue(true)
        }
      });
    });

    // 表示項目を反映する。
    if (this.saveData.conditionData.displayItemList !== null
      && this.saveData.conditionData.displayItemList !== undefined
      && this.saveData.conditionData.displayItemList.length > 0) {

      this.selectedDisplayItemArray = this.saveData.conditionData.displayItemList;
      this.fromDisplayItemArray = [];
      this.displayItemList.filter(item => {
        const flag = this.saveData.conditionData.displayItemList.some(savedItem => {
          return savedItem.id !== item.id
        });
        if (flag) {
          this.fromDisplayItemArray.push(item);
        }
      })
    }

    // 検索条件を反映する
    this.saveData.conditionData.searchConditionList.forEach(condition => {
      this.pushSearchCondition();
      const fgroup = this.searchConditionFormArray.at(this.searchConditionFormArray.length - 1);
      fgroup.get('fieldSelected').setValue(condition.field.id);
      fgroup.get('conditionValue').setValue(condition.conditionValue);
      fgroup.get('matchTypeSelected').setValue(condition.matchType);
      fgroup.get('operatorSelected').setValue(condition.operator);
    });

    // 並び順を反映する
    if (this.saveData.conditionData.orderConditionList !== null
      && this.saveData.conditionData.orderConditionList !== undefined
      && this.saveData.conditionData.orderConditionList.length > 0) {
      this.saveData.conditionData.orderConditionList.forEach(orderCondition => {
        this.pushOrderCondition()
        const fgroup = this.orderConditionFormArray.at(this.orderConditionFormArray.length - 1);
        fgroup.get('orderFieldSelected').setValue(orderCondition.orderField.id);
        fgroup.get('orderFieldKeyWordSelected').setValue(orderCondition.orderFieldKeyWord);
      });
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
    if (this.saveData.id) {
      this.complexSearchDataShereService.updateSearchCondition(this.saveData).subscribe(data => {
        this.dialog.open(NoticeDialogComponent, {
          data: { contents: '検索条件を保存しました。' }
        });
      });
    } else {
      this.complexSearchDataShereService.addSearchCondition(this.saveData).subscribe(data => {
        this.dialog.open(NoticeDialogComponent, {
          data: { contents: '検索条件を保存しました。' }
        });
      });
    }

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
          this.saveData.discloseGroupIDs.push(this.groupList[i].id);
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

}
