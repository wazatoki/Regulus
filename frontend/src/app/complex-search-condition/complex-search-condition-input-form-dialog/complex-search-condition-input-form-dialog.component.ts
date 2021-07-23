import { Component, OnInit, Inject, EventEmitter } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Category } from 'src/app/services/models/search/category';
import { SaveData } from 'src/app/services/models/search/save-data';
import { FormBuilder, FormGroup, FormControl, FormArray, AbstractControl, Validators } from '@angular/forms';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { FieldAttr } from 'src/app/services/models/search/field-attr';
import { StaffGroup } from 'src/app/services/models/group/staff-group';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SearchCondition } from 'src/app/services/models/search/search-condition';
import { OrderCondition } from 'src/app/services/models/search/order-condition';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { HttpErrorResponse } from '@angular/common/http';

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
  groupList: StaffGroup[] = [];
  displayItemList: FieldAttr[] = [];
  searchConditionList: FieldAttr[] = [];
  orderConditionList: FieldAttr[] = [];

  form: FormGroup;

  submitted: EventEmitter<string> = new EventEmitter();

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
    this.onSelectCategory();
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
    this.categorySelected.setValue(this.saveData.category.name)
    this.onSelectCategory()

    this.saveConditions.get('patternName').setValue(this.saveData.patternName)
    this.saveConditions.get('isDisclose').setValue(this.saveData.isDisclose)
    if (this.saveData.isDisclose && this.saveData.discloseGroups) {
      this.discloseGroupFormArray.controls.forEach((v, i) => {
        this.saveData.discloseGroups.forEach(g => {
          if (this.groupList[i].id === g.id) {
            v.setValue(true)
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
      this.pushSearchCondition(condition);
      // const fgroup = this.searchConditionFormArray.at(this.searchConditionFormArray.length - 1);
      // fgroup.get('fieldSelected').setValue(condition.searchField.id);
      // fgroup.get('conditionValue').setValue(condition.conditionValue);
      // fgroup.get('matchTypeSelected').setValue(condition.matchType);
      // fgroup.get('operatorSelected').setValue(condition.operator);
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
        fieldSelected: new FormControl(this.searchConditionList[0].id),
        conditionValue: new FormControl('', [Validators.required]),
        matchTypeSelected: new FormControl(''),
        operatorSelected: new FormControl(''),
      }));
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
    if (this.selectedCategory && this.selectedCategory.searchItems.staffGroups) {
      this.discloseGroupFormArray.clear();
      this.selectedCategory.searchItems.staffGroups.forEach(g => {
        this.discloseGroupFormArray.push(this.fb.control(''));
      })
      this.groupList = this.selectedCategory.searchItems.staffGroups
    } else {
      this.discloseGroupFormArray.clear();
      this.groupList = [];
    }

  }

  onSelectCategory() {

    if (this.categorySelected.value) {
      this.selectedCategory = this.data.categories.find((c) => {
        return (c.name === this.categorySelected.value)
      })
      this.initGroups()

      if (this.selectedCategory) {
        this.isShowDisplayItem = this.selectedCategory.searchItems.isShowDisplayItem
        this.isShowOrderCondition = this.selectedCategory.searchItems.isShowOrderCondition
        this.searchConditionList = this.selectedCategory.searchItems.searchConditionList
        this.orderConditionList = this.selectedCategory.searchItems.orderConditionList
        this.initSelectedDisplayItems();
        this.searchConditionFormArray.clear();
        this.orderConditionFormArray.clear();
      }
    }
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
        orderFieldKeyWord: {value: formGroup.get('orderFieldKeyWordSelected').value},
      }
      result.push(condition);
    });
    return result;
  }

  createSearchCondition(): SearchCondition[] {
    const result: SearchCondition[] = [];
    this.searchConditionFormArray.controls.forEach((formGroup: FormGroup, i) => {
      let field: FieldAttr;
      const conditionValue = (fieldType: string) => {
        if (fieldType == 'boolean') {
          if (formGroup.get('conditionValue').value) {
            return 'true';
          } else {
            return 'false';
          }
        } else {
          return formGroup.get('conditionValue').value;
        }
      }

      this.selectedCategory.searchItems.searchConditionList.forEach((v, i) => {
        if (v.id == formGroup.get('fieldSelected').value) {
          field = v;
        }
      });

      const condition: SearchCondition = {
        searchField: field,
        conditionValue: conditionValue(field.fieldType.value),
        matchType: {value: formGroup.get('matchTypeSelected').value},
        operator: {value: formGroup.get('operatorSelected').value},
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
      this.data.saveData.category = this.selectedCategory
      this.data.saveData.patternName = this.saveConditions.get('patternName').value;
      if (this.saveConditions.get('isDisclose').value) {
        this.data.saveData.isDisclose = true;
      } else {
        this.data.saveData.isDisclose = false;
      }

      this.discloseGroupFormArray.controls.forEach((v, i) => {
        if (v.value === true) {
          this.data.saveData.discloseGroups.push({id: this.groupList[i].id, name: ''});
        }
      });
    }

    this.data.saveData.conditionData = this.createSearchData();

  }

  onSubmit() {

    if (this.form.valid) {

      this.createSaveData();

      if (this.saveData.id) {

        this.complexSearchDataShereService.updateSearchCondition(this.saveData).subscribe((res: SaveData | HttpErrorResponse) => {

          if (res instanceof HttpErrorResponse == true) {

            this.dialog.open(NoticeDialogComponent, {
              data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。<br/>データの整合性を確認してください。' }
            });

          }else{

            this.dialog.open(NoticeDialogComponent, {
              data: { contents: '検索条件を修正しました。' }
            });

            this.submitted.emit('');
          
          }
        });

      } else {

        this.complexSearchDataShereService.addSearchCondition(this.saveData).subscribe((res: SaveData | HttpErrorResponse) => {
        
          if (res instanceof HttpErrorResponse == true) {

            this.dialog.open(NoticeDialogComponent, {
              data: { contents: 'エラーが発生したため処理が正常に完了しませんでした。<br/>データの整合性を確認してください。' }
            });

          }else{

            this.dialog.open(NoticeDialogComponent, {
              data: { contents: '検索条件を保存しました。' }
            });

            this.submitted.emit('');
          
          }
        });
      }
    }
  }

  getPatternNameErrorMessage() {
    return this.saveConditions.get('patternName').hasError('required') ? '検索パターン名称は必須項目です。' : '';
  }

  initForm(): FormGroup {
    return this.fb.group({
      searchCondition: this.fb.array([]),
      orderCondition: this.fb.array([]),
      saveCondition: this.fb.group({
        category: this.fb.control(this.categories[0].name),
        patternName: this.fb.control('', [Validators.required]),
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
    this.onSelectCategory();
  }

  ngOnInit() {
    this.dialogRef.updateSize("1100px")
    // saveDataの編集のときは値をフォームに反映する
    if (this.data.saveData !== null && this.data.saveData !== undefined && this.data.saveData.id !== '') {
      this.setSavedDataToForm()
    }
  }
}

export interface DialogData {
  categories: Category[];
  saveData: SaveData;
}
