import { Component, OnInit, Inject, Input } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Category } from 'src/app/services/models/search/category';
import { SaveData } from 'src/app/services/models/search/save-data';
import { FormBuilder, FormGroup, FormControl, FormArray, AbstractControl } from '@angular/forms';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { FieldAttr } from 'src/app/services/models/search/field-attr';
import { Group } from 'src/app/services/models/group/group';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SearchCondition } from 'src/app/services/models/search/search-condition';
import { OrderCondition } from 'src/app/services/models/search/order-condition';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-complex-search-condition-input-form-dialog',
  templateUrl: './complex-search-condition-input-form-dialog.component.html',
  styleUrls: ['./complex-search-condition-input-form-dialog.component.css']
})
export class ComplexSearchConditionInputFormDialogComponent implements OnInit {

  selectedCategory: Category;
  selectedDisplayItemArray: FieldAttr[] = [];
  fromDisplayItemArray: FieldAttr[] = [];
  isShowDisplayItem: boolean = false;
  isShowOrderCondition: boolean = false;
  isShowSaveCondition: boolean = true;
  groupList: Group[] = [];
  displayItemList: FieldAttr[] = [];
  searchConditionList: FieldAttr[] = [];
  orderConditionList: FieldAttr[] = [];

  form: FormGroup;


  get saveData(): SaveData {
    return this.data.saveData;
  }

  get categories(): Category[] {
    return this.data.categories;
  }
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




  onClearClick(): void {
    this.form = this.initForm()
    this.selectedDisplayItemArray = [];
    this.fromDisplayItemArray = [];
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

  removeOrderCondition(i: number) {
    this.orderConditionFormArray.removeAt(i);
  }

  pushOrderCondition() {
    this.orderConditionFormArray.push(new FormGroup({
      orderFieldSelected: new FormControl(''),
      orderFieldKeyWordSelected: new FormControl(''),
    }));
  }

  setSavedDataToForm() {
    // カテゴリーの反映
    this.categorySelected.setValue(this.saveData.category)
    this.onSelectCategory()

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
      this.selectedCategory.searchItems.displayItemList.filter(item => {
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

  removeSearchCondition(i: number) {
    this.searchConditionFormArray.removeAt(i);
  }

  controlDrop(event: CdkDragDrop<AbstractControl[]>) {
    moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
  }

  pushSearchCondition() {
    this.searchConditionFormArray.push(new FormGroup({
      fieldSelected: new FormControl(''),
      conditionValue: new FormControl(''),
      matchTypeSelected: new FormControl(''),
      operatorSelected: new FormControl(''),
    }));
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

  initSelectedDisplayItems() {
    if (this.selectedCategory) {
      this.fromDisplayItemArray = [];
      this.selectedCategory.searchItems.displayItemList.forEach(item => {
        this.fromDisplayItemArray.push(item)
      })
      this.selectedDisplayItemArray = [];
    } else {
      this.fromDisplayItemArray = [];
      this.selectedDisplayItemArray = [];
    }
  }

  initGroups() {
    if (this.selectedCategory && this.selectedCategory.searchItems.groups) {
      this.discloseGroupFormArray.clear();
      this.selectedCategory.searchItems.groups.forEach(g => {
        this.discloseGroupFormArray.push(this.fb.control(''));
      })
      this.groupList = this.selectedCategory.searchItems.groups
    } else {
      this.discloseGroupFormArray.clear();
      this.groupList = [];
    }

  }

  onSelectCategory() {
    this.selectedCategory = this.data.categories.find((c) => {
      return (c.name === this.categorySelected.value)
    })

    this.initGroups()

    this.isShowDisplayItem = this.selectedCategory.searchItems.isShowDisplayItem
    this.isShowOrderCondition = this.selectedCategory.searchItems.isShowOrderCondition
    this.searchConditionList = this.selectedCategory.searchItems.searchConditionList
    this.orderConditionList = this.selectedCategory.searchItems.orderConditionList
    this.initSelectedDisplayItems();
  }


  createOrderCondition(): OrderCondition[] {
    const result: OrderCondition[] = [];
    this.orderConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      this.selectedCategory.searchItems.orderConditionList.forEach((v, i) => {
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

  createSearchCondition(): SearchCondition[] {
    const result: SearchCondition[] = [];
    this.searchConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;

      this.selectedCategory.searchItems.searchConditionList.forEach((v, i) => {
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

  createSaveData(): void {

    if (this.data.saveData === null || this.data.saveData === undefined) {
      this.data.saveData = this.complexSearchDataShereService.initSaveDataObj();
    }

    if (this.isShowSaveCondition) {
      this.data.saveData.category = this.selectedCategory.name
      this.data.saveData.patternName = this.saveConditions.get('patternName').value;
      this.data.saveData.isDisclose = this.saveConditions.get('isDisclose').value;
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        if (v.value === true) {
          this.data.saveData.discloseGroupIDs.push(this.groupList[i].id);
        }
      });
    }

    this.data.saveData.conditionData = this.createSearchData();

  }

  onSubmit() {
    this.createSaveData();
    if (this.data.saveData.id) {
      this.complexSearchDataShereService.updateSearchCondition(this.data.saveData).subscribe(data => {
        this.dialog.open(NoticeDialogComponent, {
          data: { contents: '検索条件を保存しました。' }
        });
      });
    } else {
      this.complexSearchDataShereService.addSearchCondition(this.data.saveData).subscribe(data => {
        this.dialog.open(NoticeDialogComponent, {
          data: { contents: '検索条件を保存しました。' }
        });
      });
    }
  }

  initForm(): FormGroup {
    return this.fb.group({
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

  constructor(
    private dialogRef: MatDialogRef<ComplexSearchConditionInputFormDialogComponent>,
    @Inject(MAT_DIALOG_DATA) private data: DialogData,
    private fb: FormBuilder,
    private dialog: MatDialog,
    private complexSearchDataShereService: ComplexSearchService,
  ) {
    this.form = this.initForm();
  }

  ngOnInit() {
    this.dialogRef.updateSize("1100px")
    // saveDataの編集のときは値をフォームに反映する
    if (this.saveData !== null && this.saveData !== undefined && this.saveData.id !== '') {
      this.setSavedDataToForm()
    }
  }
}

export interface DialogData {
  categories: Category[];
  saveData: SaveData;
}
