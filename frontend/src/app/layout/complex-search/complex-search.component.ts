import { Component, OnInit, Input } from '@angular/core';
import { FormBuilder, FormControl, FormArray, FormGroup, AbstractControl, Validators } from '@angular/forms';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { StaffGroup } from '../../services/models/group/staff-group';
import { FieldAttr } from '../../services/models/search/field-attr';
import { ConditionData } from '../../services/models/search/condition-data';
import { SearchCondition } from '../../services/models/search/search-condition';
import { OrderCondition } from '../../services/models/search/order-condition';
import { SaveData } from '../../services/models/search/save-data';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { MatDialog } from '@angular/material/dialog';
import { NoticeDialogComponent } from '../dialog/notice-dialog/notice-dialog.component';
import { HttpErrorResponse } from '@angular/common/http';
import { ComplexSearchItems } from 'src/app/services/models/search/complex-search-items';

@Component({
  selector: 'app-complex-search',
  templateUrl: './complex-search.component.html',
  styleUrls: ['./complex-search.component.css']
})
export class ComplexSearchComponent implements OnInit {

  form: FormGroup;

  selectedDisplayItemArray: FieldAttr[];
  fromDisplayItemArray: FieldAttr[];

  @Input() complexSearchItems: ComplexSearchItems = {
    displayItemList: [],
    searchConditionList: [],
    orderConditionList: [],
    isShowDisplayItem: false,
    isShowOrderCondition: false,
    isShowSaveCondition: false,
    staffGroups: []
  };
  @Input() saveData: SaveData = this.complexSearchDataShereService.initSaveDataObj();

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

  constructor(
    private fb: FormBuilder,
    private complexSearchDataShereService: ComplexSearchService,
    private dialog: MatDialog) {

    this.form = this.fb.group({
      searchCondition: this.fb.array([]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({
        patternName: this.fb.control(''),
        isDisclose: this.fb.control(''),
        discloseGroups: this.fb.array([]),
      }),
    });
  }

  ngOnInit() {
    this.complexSearchItems.staffGroups.forEach(g => {
      this.discloseGroupFormArray.push(this.fb.control(''));
    });
    this.initSelectedDisplayItems();

    // saveDataが設定されているときは値をフォームに反映する
    if (this.saveData !== null && this.saveData !== undefined) {
      this.setSavedDataToForm();
    }
  }

  setSavedDataToForm() {
    this.saveConditions.get('patternName').setValue(this.saveData.patternName);
    this.saveConditions.get('isDisclose').setValue(this.saveData.isDisclose);

    if (this.saveData.isDisclose && this.saveData.discloseGroups) {
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        this.saveData.discloseGroups.forEach(g => {
          if (this.complexSearchItems.staffGroups[i].id === g.id) {
            v.setValue(true);
          }
        });
      });
    }

    // 表示項目を反映する。
    if (this.saveData.conditionData.displayItemList !== null
      && this.saveData.conditionData.displayItemList !== undefined
      && this.saveData.conditionData.displayItemList.length > 0) {

      this.selectedDisplayItemArray = this.saveData.conditionData.displayItemList;
      this.fromDisplayItemArray = [];
      this.complexSearchItems.displayItemList.filter(item => {
        const flag = this.saveData.conditionData.displayItemList.some(savedItem => {
          return savedItem.id !== item.id;
        });
        if (flag) {
          this.fromDisplayItemArray.push(item);
        }
      });
    }

    // 検索条件を反映する
    this.saveData.conditionData.searchConditionList.forEach(condition => {
      this.pushSearchCondition(condition);
      // const fgroup = this.searchConditionFormArray.at(this.searchConditionFormArray.length - 1);
      // fgroup.get('fieldSelected').setValue(condition.searchField.id);
      // fgroup.get('conditionValue').setValue(condition.conditionValue);
      // fgroup.get('matchTypeSelected').setValue(condition.matchType.value);
      // fgroup.get('operatorSelected').setValue(condition.operator.value);
    });

    // 並び順を反映する
    if (this.saveData.conditionData.orderConditionList !== null
      && this.saveData.conditionData.orderConditionList !== undefined
      && this.saveData.conditionData.orderConditionList.length > 0) {

      this.saveData.conditionData.orderConditionList.forEach(orderCondition => {
        this.pushOrderCondition();
        const fgroup = this.orderConditionFormArray.at(this.orderConditionFormArray.length - 1);
        fgroup.get('orderFieldSelected').setValue(orderCondition.orderField.id);
        fgroup.get('orderFieldKeyWordSelected').setValue(orderCondition.orderFieldKeyWord);
      });
    }
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
    moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
  }

  pushSearchCondition(condition?: SearchCondition) {
    if (condition && condition.searchField && condition.searchField.id) {
      this.searchConditionFormArray.push(new FormGroup({
        fieldSelected: new FormControl(condition.searchField.id),
        conditionValue: new FormControl(condition.conditionValue, [Validators.required]),
        matchTypeSelected: new FormControl(condition.matchType.value),
        operatorSelected: new FormControl(condition.operator.value),
      }));
    } else {
      this.searchConditionFormArray.push(new FormGroup({
        fieldSelected: new FormControl(this.complexSearchItems.searchConditionList[0].id),
        conditionValue: new FormControl('', [Validators.required]),
        matchTypeSelected: new FormControl(''),
        operatorSelected: new FormControl(''),
      }));
    }

    // this.searchConditionFormArray.push(new FormGroup({
    //   fieldSelected: new FormControl(''),
    //   conditionValue: new FormControl(''),
    //   matchTypeSelected: new FormControl(''),
    //   operatorSelected: new FormControl(''),
    // }));
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
      this.complexSearchDataShereService.updateSearchCondition(this.saveData).subscribe((res: SaveData | HttpErrorResponse) => {
        if (res instanceof HttpErrorResponse === true) {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。<br/>データの整合性を確認してください。' }
          });

        } else {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: '検索条件を修正しました。' }
          });

        }
      });

    } else {
      this.complexSearchDataShereService.addSearchCondition(this.saveData).subscribe((res: SaveData | HttpErrorResponse) => {
        if (res instanceof HttpErrorResponse === true) {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。<br/>データの整合性を確認してください。' }
          });

        } else {

          this.dialog.open(NoticeDialogComponent, {
            data: { contents: '検索条件を保存しました。' }
          });

        }
      });
    }

  }

  clickSearch(): void {
    const data: ConditionData = this.createSearchData();
    this.complexSearchDataShereService.orderComplexSearch(data);
  }

  createSearchCondition(): SearchCondition[] {
    const result: SearchCondition[] = [];
    this.searchConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      const conditionValue = (fieldType: string) => {
        if (fieldType === 'boolean') {
          if (formGroup.get('conditionValue').value) {
            return 'true';
          } else {
            return 'false';
          }
        } else {
          return formGroup.get('conditionValue').value;
        }
      };

      this.complexSearchItems.searchConditionList.forEach((v, index) => {
        if (v.id === formGroup.get('fieldSelected').value) {
          field = v;
        }
      });

      if (field) {
        const condition: SearchCondition = {
          searchField: field,
          conditionValue: conditionValue(field.fieldType.value),
          matchType: { value: formGroup.get('matchTypeSelected').value },
          operator: { value: formGroup.get('operatorSelected').value },
        };
        result.push(condition);
      }

    });

    return result;
  }

  createOrderCondition(): OrderCondition[] {
    const result: OrderCondition[] = [];
    this.orderConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      this.complexSearchItems.orderConditionList.forEach((v, index) => {
        if (v.id === formGroup.get('orderFieldSelected').value) {
          field = v;
        }
      });

      if (field) {
        const condition: OrderCondition = {
          orderField: field,
          orderFieldKeyWord: { value: formGroup.get('orderFieldKeyWordSelected').value },
        };
        result.push(condition);
      }

    });

    return result;
  }

  createSaveData(): void {

    if (this.saveData === null || this.saveData === undefined) {
      this.saveData = this.complexSearchDataShereService.initSaveDataObj();
    }

    if (this.complexSearchItems.isShowSaveCondition) {
      this.saveData.patternName = this.saveConditions.get('patternName').value;
      this.saveData.isDisclose = this.saveConditions.get('isDisclose').value;
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        if (v.value === true) {
          this.saveData.discloseGroups.push({ id: this.complexSearchItems.staffGroups[i].id, name: '' });
        }
      });
    }

    this.saveData.conditionData = this.createSearchData();

  }

  createSearchData(): ConditionData {
    const data: ConditionData = this.complexSearchDataShereService.initConditionDataObj();

    if (this.complexSearchItems.isShowDisplayItem) {
      data.displayItemList = this.selectedDisplayItemArray;
    }

    data.searchConditionList = this.createSearchCondition();

    if (this.complexSearchItems.isShowOrderCondition) {
      data.orderConditionList = this.createOrderCondition();
    }
    return data;
  }

  initSelectedDisplayItems() {
    this.fromDisplayItemArray = [];
    this.complexSearchItems.displayItemList.forEach(item => {
      this.fromDisplayItemArray.push(item);
    });
    this.selectedDisplayItemArray = [];
  }

}// end of class
