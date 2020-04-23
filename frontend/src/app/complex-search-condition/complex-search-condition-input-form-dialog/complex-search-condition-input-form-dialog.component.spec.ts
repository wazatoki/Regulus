import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionInputFormDialogComponent } from './complex-search-condition-input-form-dialog.component';
import { Component, DebugElement, OnInit } from '@angular/core';
import { MatDialog, MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

import { ceateTestArrayForMasterMaintenanceTest } from 'src/app/services/models/search/category.spec'
import { createTestInstance1 as createSaveData } from 'src/app/services/models/search/save-data.spec';
import { Category } from 'src/app/services/models/search/category';
import { SaveData } from 'src/app/services/models/search/save-data';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClient } from '@angular/common/http';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import { By } from '@angular/platform-browser';
import { OverlayContainer } from '@angular/cdk/overlay';
import { ComplexSearchConditionItemComponent } from 'src/app/layout/complex-search/complex-search-condition-item/complex-search-condition-item.component';
import { ComplexSearchOrderItemComponent } from 'src/app/layout/complex-search/complex-search-order-item/complex-search-order-item.component';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { createInitSaveData } from 'src/app/services/models/search/save-data.spec'
import { createInitConditionData } from 'src/app/services/models/search/condition-data.spec';
import { Subject } from 'rxjs';

describe('ComplexSearchConditionInputFormDialogComponent', () => {
  let component: ComplexSearchConditionInputFormDialogComponent;
  let fixture: ComponentFixture<ComplexSearchConditionInputFormDialogComponent>;
  const matDialogRefSpy = jasmine.createSpyObj('MatDialogRef', ['close', 'updateSize']);
  let complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService>;
  let dialog: MatDialog;
  const categories: Category[] = ceateTestArrayForMasterMaintenanceTest();
  const saveData: SaveData = createSaveData();
  const dialogPassedData = {
    categories: categories,
    saveData: saveData,
  };

  beforeEach(async(() => {
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
      ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'updateSearchCondition', 'addSearchCondition']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionInputFormDialogComponent,
      ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
        FormsModule,
        ReactiveFormsModule,
        MatDialogModule,
        MatCardModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatCheckboxModule,
        MatSelectModule,
        MatRadioModule,
        MatGridListModule,
        DragDropModule,
        NoopAnimationsModule,
      ],
      providers: [
        {
          provide: FormBuilder,
          useValue: new FormBuilder()
        },
        {
          provide: ComplexSearchService,
          useValue: complexSearchServiceSpy
        },
        {
          provide: MatDialogRef,
          useValue: matDialogRefSpy,
        },
        {
          provide: MAT_DIALOG_DATA,
          useValue: dialogPassedData,
        },
        {
          provide: HttpClient, useValue: {}
        }
      ],
    })
      .overrideModule(BrowserDynamicTestingModule, {
        set: {
          entryComponents: [ComplexSearchConditionInputFormDialogComponent],
        }
      })
      .compileComponents();

    dialog = TestBed.get(MatDialog);
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchConditionInputFormDialogComponent);
    component = fixture.componentInstance;

    complexSearchServiceSpy = TestBed.get(ComplexSearchService);
    complexSearchServiceSpy.initSaveDataObj.and.returnValue(createInitSaveData());
    complexSearchServiceSpy.initConditionDataObj.and.returnValue(createInitConditionData());
    complexSearchServiceSpy.updateSearchCondition.and.returnValue(new Subject<SaveData>().asObservable());
    complexSearchServiceSpy.addSearchCondition.and.returnValue(new Subject<SaveData>().asObservable());

    fixture.detectChanges();
  });

  it('should create', () => {

    expect(component).toBeTruthy();
  });

  it('should select category', async () => {
    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-category-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;

    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()
      fixture.detectChanges();
    });

    expect(component.isShowDisplayItem).toBe(true)
    expect(component.isShowOrderCondition).toBe(true)

  });

  it('should click add condition button', () => {

    component.isShowDisplayItem = false;
    component.isShowOrderCondition = false;
    component.isShowSaveCondition = false;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".push-search-condition"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    buttonEl.click();
    fixture.detectChanges();

    const itemBoxDeArray: DebugElement[] = fixture.debugElement.queryAll(By.directive(ComplexSearchConditionItemComponent));

    expect(itemBoxDeArray.length).toBe(4);
  });

  it('should click add order button', () => {

    component.isShowDisplayItem = false;
    component.isShowOrderCondition = true;
    component.isShowSaveCondition = false;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".push-order-condition"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    buttonEl.click();
    fixture.detectChanges();

    const itemBoxDeArray: DebugElement[] = fixture.debugElement.queryAll(By.directive(ComplexSearchOrderItemComponent));

    expect(itemBoxDeArray.length).toBe(4);
  });

  it('should click save button', () => {
    const spy: jasmine.SpyObj<ComplexSearchService> = TestBed.get(ComplexSearchService);
    component.isShowDisplayItem = true;
    component.isShowOrderCondition = true;
    component.isShowSaveCondition = true;
    fixture.detectChanges();

    const formDebugElement: DebugElement = fixture.debugElement.query(By.css('form'));
    formDebugElement.triggerEventHandler('submit', null);
    fixture.detectChanges();

    if (component.saveData.id) {
      expect(spy.updateSearchCondition).toHaveBeenCalled();
    } else {
      expect(spy.addSearchCondition).toHaveBeenCalled();
    }
  });

  it('should create save data', async () => {

    component.isShowDisplayItem = true;
    component.isShowOrderCondition = true;
    component.isShowSaveCondition = true;
    fixture.detectChanges();

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-category-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;

    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()
      fixture.detectChanges();
    });

    const patternNameDe: DebugElement = fixture.debugElement.query(By.css("input.pattern-name"));
    const patternNameEl: HTMLInputElement = patternNameDe.nativeElement;
    patternNameEl.value = 'sample pattern name';
    patternNameEl.dispatchEvent(new Event('input'));
    const isDiscloseDe: DebugElement = fixture.debugElement.query(By.css(".is-disclose label"));
    const isDiscloseEl: HTMLInputElement = isDiscloseDe.nativeElement;
    isDiscloseEl.click();
    const groupDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".disclosure-destination-group label"));
    const groupEl0: HTMLInputElement = groupDe[0].nativeElement;
    const groupEl1: HTMLInputElement = groupDe[1].nativeElement;
    groupEl0.click();
    groupEl1.click();

    component.pushSearchCondition();
    component.pushOrderCondition();
    (component.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
    (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    fixture.detectChanges();

    component.createSaveData();
    const expectSaveData = component.saveData;
    expect(expectSaveData.patternName).toBe('sample pattern name');
    expect(expectSaveData.isDisclose).toBe(false);
    expect(expectSaveData.discloseGroupIDs).toEqual(['test-group-id-1', 'test-group-id-2']);
    expect(expectSaveData.conditionData.searchConditionList[0].field).toEqual(component.selectedCategory.searchItems.searchConditionList[0]);
    expect(expectSaveData.conditionData.searchConditionList[0].conditionValue).toEqual('value1');
    expect(expectSaveData.conditionData.searchConditionList[0].matchType).toEqual('match');
    expect(expectSaveData.conditionData.searchConditionList[0].operator).toEqual('and');
    expect(expectSaveData.conditionData.orderConditionList[0].orderField).toEqual(component.selectedCategory.searchItems.orderConditionList[1]);
    expect(expectSaveData.conditionData.orderConditionList[0].orderFieldKeyWord).toEqual('asc');
  });

  it('should delete condition', () => {

    component.isShowDisplayItem = false;
    component.isShowOrderCondition = false;
    component.isShowSaveCondition = false;
    fixture.detectChanges();


    component.pushSearchCondition();
    component.pushSearchCondition();
    fixture.detectChanges();
    (component.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.searchConditionFormArray.controls[1] as FormGroup).get('fieldSelected').setValue('fieldid2');
    (component.searchConditionFormArray.controls[1] as FormGroup).get('conditionValue').setValue('value2');
    (component.searchConditionFormArray.controls[1] as FormGroup).get('matchTypeSelected').setValue('unmatch');
    (component.searchConditionFormArray.controls[1] as FormGroup).get('operatorSelected').setValue('or');
    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".search-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').value).toEqual('fieldid2');
  });

  it('should delete order condition', () => {

    component.isShowDisplayItem = false;
    component.isShowOrderCondition = true;
    component.isShowSaveCondition = false;
    fixture.detectChanges();


    component.pushOrderCondition();
    component.pushOrderCondition();
    fixture.detectChanges();
    (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid1');
    (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    (component.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
    (component.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldKeyWordSelected').setValue('desc');

    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".order-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').value).toEqual('fieldid2');
  });

  it('should clear form', async () => {
    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-category-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;

    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()
      fixture.detectChanges();
    });

    await fixture.whenStable().then(() => {
      const patternNameDe: DebugElement = fixture.debugElement.query(By.css("input.pattern-name"));
      const patternNameEl: HTMLInputElement = patternNameDe.nativeElement;
      patternNameEl.value = 'sample pattern name';
      patternNameEl.dispatchEvent(new Event('input'));
      const isDiscloseDe: DebugElement = fixture.debugElement.query(By.css(".is-disclose label"));
      const isDiscloseEl: HTMLInputElement = isDiscloseDe.nativeElement;
      isDiscloseEl.click();
      const groupDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".disclosure-destination-group label"));
      const groupEl0: HTMLInputElement = groupDe[0].nativeElement;
      const groupEl1: HTMLInputElement = groupDe[1].nativeElement;
      groupEl0.click();
      groupEl1.click();

      component.pushSearchCondition();
      component.pushOrderCondition();
      (component.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
      (component.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
      (component.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
      (component.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
      (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
      (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
      fixture.detectChanges();
    });

    await fixture.whenStable().then(() => {
      component.pushOrderCondition();
      component.pushOrderCondition();
      fixture.detectChanges();
      (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid1');
      (component.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
      (component.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
      (component.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldKeyWordSelected').setValue('desc');
      fixture.detectChanges();
    });

    component.onClearClick();
    fixture.detectChanges();
    expect(selectEl.textContent).not.toContain('TEST_CATEGORY_1');
  });

});
