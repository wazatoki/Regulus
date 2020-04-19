import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionInputFormComponent } from './complex-search-condition-input-form.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { MatDialogModule, MatDialog } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { MatButtonModule } from '@angular/material/button';
import { ViewChild, Component, DebugElement } from '@angular/core';

import { ceateTestArrayForMasterMaintenanceTest } from 'src/app/services/models/search/category.spec'
import { createTestArray as createFielsAttrTeatArray } from 'src/app/services/models/search/field-attr.spec';
import { ceateTestArray as createSearchConditionArray } from 'src/app/services/models/search/search-condition.spec'
import { ceateTestArray as createOrderConditionArray } from 'src/app/services/models/search/order-condition.spec'
import { ceateTestArray as createGroupArray } from 'src/app/services/models/group/group.spec';
import { createTestInstance1 as createSaveData } from 'src/app/services/models/search/save-data.spec';
import { createInitSaveData } from 'src/app/services/models/search/save-data.spec'
import { createInitConditionData } from 'src/app/services/models/search/condition-data.spec';

import { Category } from 'src/app/services/models/search/category';
import { FieldAttr } from 'src/app/services/models/search/field-attr';
import { Group } from 'src/app/services/models/group/group';
import { SaveData } from 'src/app/services/models/search/save-data';
import { By } from '@angular/platform-browser';
import { ComplexSearchConditionItemComponent } from 'src/app/layout/complex-search/complex-search-condition-item/complex-search-condition-item.component';
import { ComplexSearchOrderItemComponent } from 'src/app/layout/complex-search/complex-search-order-item/complex-search-order-item.component';
import { Subject } from 'rxjs';


describe('ComplexSearchConditionInputFormComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;
  let spy: jasmine.SpyObj<ComplexSearchService>;
  let saveData: SaveData
  // let component: ComplexSearchConditionInputFormComponent;
  // let fixture: ComponentFixture<ComplexSearchConditionInputFormComponent>;

  beforeEach(async(() => {
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'updateSearchCondition', 'addSearchCondition']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionInputFormComponent,
        TestHostComponent,
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
          provide: MatDialog,
          useValue: {}
        },
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;

    spy = TestBed.get(ComplexSearchService);
    spy.initSaveDataObj.and.returnValue(createInitSaveData());
    spy.initConditionDataObj.and.returnValue(createInitConditionData());
    spy.updateSearchCondition.and.returnValue(new Subject<SaveData>().asObservable());
    spy.addSearchCondition.and.returnValue(new Subject<SaveData>().asObservable());

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

      expect(component.complexSearchConditionInputFormComponent.isShowDisplayItem).toBe(true)
      expect(component.complexSearchConditionInputFormComponent.isShowOrderCondition).toBe(true)
    });
  });

  it('should click add condition button', () => {

    component.complexSearchConditionInputFormComponent.isShowDisplayItem = false;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = false;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = false;
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

    component.complexSearchConditionInputFormComponent.isShowDisplayItem = false;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = true;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = false;
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
    component.complexSearchConditionInputFormComponent.isShowDisplayItem = true;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = true;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = true;
    fixture.detectChanges();

    const formDebugElement: DebugElement = fixture.debugElement.query(By.css('form'));
    formDebugElement.triggerEventHandler('submit', null);
    fixture.detectChanges();

    if(component.complexSearchConditionInputFormComponent.saveData.id){
      expect(spy.updateSearchCondition).toHaveBeenCalled();
    }else{
      expect(spy.addSearchCondition).toHaveBeenCalled();
    }
  });

  it('should create save data', () => {

    component.complexSearchConditionInputFormComponent.isShowDisplayItem = true;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = true;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = true;
    fixture.detectChanges();

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

    component.complexSearchConditionInputFormComponent.pushSearchCondition();
    component.complexSearchConditionInputFormComponent.pushOrderCondition();
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    fixture.detectChanges();

    component.complexSearchConditionInputFormComponent.createSaveData();
    saveData = component.complexSearchConditionInputFormComponent.saveData;
    expect(saveData.patternName).toBe('sample pattern name');
    expect(saveData.isDisclose).toBe(false);
    expect(saveData.discloseGroupIDs).toEqual(['test-group-id-1', 'test-group-id-2']);
    expect(saveData.conditionData.searchConditionList[0].field).toEqual(component.complexSearchConditionInputFormComponent.searchConditionList[0]);
    expect(saveData.conditionData.searchConditionList[0].conditionValue).toEqual('value1');
    expect(saveData.conditionData.searchConditionList[0].matchType).toEqual('match');
    expect(saveData.conditionData.searchConditionList[0].operator).toEqual('and');
    expect(saveData.conditionData.orderConditionList[0].orderField).toEqual(component.complexSearchConditionInputFormComponent.orderConditionList[1]);
    expect(saveData.conditionData.orderConditionList[0].orderFieldKeyWord).toEqual('asc');
  });

  it('should delete condition', () => {

    component.complexSearchConditionInputFormComponent.isShowDisplayItem = false;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = false;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = false;
    fixture.detectChanges();


    component.complexSearchConditionInputFormComponent.pushSearchCondition();
    component.complexSearchConditionInputFormComponent.pushSearchCondition();
    fixture.detectChanges();
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[1] as FormGroup).get('fieldSelected').setValue('fieldid2');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[1] as FormGroup).get('conditionValue').setValue('value2');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[1] as FormGroup).get('matchTypeSelected').setValue('unmatch');
    (component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[1] as FormGroup).get('operatorSelected').setValue('or');
    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".search-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.complexSearchConditionInputFormComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').value).toEqual('fieldid2');
  });

  it('should delete order condition', () => {

    component.complexSearchConditionInputFormComponent.isShowDisplayItem = false;
    component.complexSearchConditionInputFormComponent.isShowOrderCondition = true;
    component.complexSearchConditionInputFormComponent.isShowSaveCondition = false;
    fixture.detectChanges();


    component.complexSearchConditionInputFormComponent.pushOrderCondition();
    component.complexSearchConditionInputFormComponent.pushOrderCondition();
    fixture.detectChanges();
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid1');
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
    (component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldKeyWordSelected').setValue('desc');

    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".order-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.complexSearchConditionInputFormComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').value).toEqual('fieldid2');
  });

});


@Component({
  template: `
  <app-complex-search-condition-input-form
    [categories]="categories"
    [displayItemList]="displayItemList"
    [searchConditionList]="searchConditionList"
    [orderConditionList]="orderConditionList"
    [groupList]="groupList"
    [saveData]="saveData"
  >
  </app-complex-search-condition-input-form>`
})
class TestHostComponent {

  @ViewChild(ComplexSearchConditionInputFormComponent, {static: true})
  complexSearchConditionInputFormComponent: ComplexSearchConditionInputFormComponent;

  categories: Category[] = ceateTestArrayForMasterMaintenanceTest();
  displayItemList: FieldAttr[] = createFielsAttrTeatArray();
  searchConditionList: FieldAttr[] = createFielsAttrTeatArray();
  orderConditionList: FieldAttr[] = createFielsAttrTeatArray();
  groupList: Group[] = createGroupArray();
  saveData: SaveData = createSaveData();
}

