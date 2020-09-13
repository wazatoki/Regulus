import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatCardModule } from '@angular/material/card';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';
import { ComplexSearchComponent } from './complex-search.component';
import { FieldAttr } from '../../services/models/search/field-attr';
import { SaveData } from '../../services/models/search/save-data';
import { ComplexSearchConditionItemComponent } from './complex-search-condition-item/complex-search-condition-item.component';
import { ComplexSearchOrderItemComponent } from "./complex-search-order-item/complex-search-order-item.component";
import { DeleteComponent } from '../form/buttons/delete/delete.component';
import { DebugElement, Component, ViewChild } from '@angular/core';
import { By } from '@angular/platform-browser';
import { Group } from '../../services/models/group/group';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { MatDialog } from '@angular/material/dialog';
import { Subject } from 'rxjs';

import { createTestArray } from 'src/app/services/models/search/field-attr.spec';
import { ceateTestArray as ceateTestArrayGroup } from 'src/app/services/models/group/group.spec';
import { createTestInstance1 as createTestInstanceSaveData } from 'src/app/services/models/search/save-data.spec';
import { createInitSaveData } from 'src/app/services/models/search/save-data.spec'
import { createInitConditionData } from 'src/app/services/models/search/condition-data.spec';
import { ClearComponent } from 'src/app/layout/form/buttons/clear/clear.component'
import { SubmitComponent } from 'src/app/layout/form/buttons/submit/submit.component'
import { Category } from 'src/app/services/models/search/category';
import { ceateTestArrayForMasterMaintenanceTest as createCategoryArrayData } from 'src/app/services/models/search/category.spec'

describe('ComplexSearchComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;
  let spy: jasmine.SpyObj<ComplexSearchService>;
  let saveData: SaveData

  beforeEach(async(() => {
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
      ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'updateSearchCondition', 'addSearchCondition']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchComponent,
        TestHostComponent,
        ComplexSearchConditionItemComponent,
        ComplexSearchOrderItemComponent,
        DeleteComponent,
        ClearComponent,
        SubmitComponent,
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatCheckboxModule,
        MatCardModule,
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

  it('should click add condition button', () => {

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
    fixture.detectChanges();

    const formDebugElement: DebugElement = fixture.debugElement.query(By.css('form'));
    formDebugElement.triggerEventHandler('submit', null);
    // const buttonDe: DebugElement = fixture.debugElement.query(By.css(".complex-search-condition-save-button"));
    // const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    // buttonEl.click();
    fixture.detectChanges();

    if (component.searchComponent.saveData.id) {
      expect(spy.updateSearchCondition).toHaveBeenCalled();
    } else {
      expect(spy.addSearchCondition).toHaveBeenCalled();
    }
  });

  it('should click search button', () => {

    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".complex-search-button"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    fixture.detectChanges();

    expect(spy.orderComplexSearch).toHaveBeenCalled();
  });


  it('should create save data', async () => {

    fixture.detectChanges();

    const patternNameDe: DebugElement = fixture.debugElement.query(By.css("input.pattern-name"));
    const patternNameEl: HTMLInputElement = patternNameDe.nativeElement;
    patternNameEl.value = 'sample pattern name';
    patternNameEl.dispatchEvent(new Event('input'));
    component.searchComponent.saveConditions.get('isDisclose').setValue(true);
    component.searchComponent.discloseGroupFormArray.controls[0].setValue(true);
    component.searchComponent.discloseGroupFormArray.controls[1].setValue(true);

    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('fieldid1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('fieldid2');
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');

    fixture.whenStable().then(() => {

      saveData = component.searchComponent.createSaveData();
      expect(saveData.patternName).toBe('sample pattern name');
      expect(saveData.isDisclose).toBe(true);
      expect(saveData.discloseGroups[0].id).toEqual('test-group-id-1');
      expect(saveData.discloseGroups[1].id).toEqual('test-group-id-2');
      expect(saveData.conditionData.searchConditionList[0].searchField).toEqual(component.category.searchItems.searchConditionList[0]);
      expect(saveData.conditionData.searchConditionList[0].conditionValue).toEqual('value1');
      expect(saveData.conditionData.searchConditionList[0].matchType).toEqual('match');
      expect(saveData.conditionData.searchConditionList[0].operator).toEqual('and');
      expect(saveData.conditionData.orderConditionList[0].orderField).toEqual(component.category.searchItems.orderConditionList[1]);
      expect(saveData.conditionData.orderConditionList[0].orderFieldKeyWord).toEqual('asc');

    });

  });

  it('should delete condition', () => {

    fixture.detectChanges();


    component.searchComponent.pushSearchCondition();
    component.searchComponent.pushSearchCondition();
    fixture.detectChanges();
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('id1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.searchComponent.searchConditionFormArray.controls[1] as FormGroup).get('fieldSelected').setValue('id2');
    (component.searchComponent.searchConditionFormArray.controls[1] as FormGroup).get('conditionValue').setValue('value2');
    (component.searchComponent.searchConditionFormArray.controls[1] as FormGroup).get('matchTypeSelected').setValue('unmatch');
    (component.searchComponent.searchConditionFormArray.controls[1] as FormGroup).get('operatorSelected').setValue('or');
    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".search-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').value).toEqual('id2');
  });

  it('should delete order condition', () => {

    fixture.detectChanges();


    component.searchComponent.pushOrderCondition();
    component.searchComponent.pushOrderCondition();
    fixture.detectChanges();
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('id1');
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    (component.searchComponent.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldSelected').setValue('id2');
    (component.searchComponent.orderConditionFormArray.controls[1] as FormGroup).get('orderFieldKeyWordSelected').setValue('desc');

    fixture.detectChanges();

    const deleteDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".order-condition .item-list .delete-button button"));
    const deleteEl0: HTMLButtonElement = deleteDe[0].nativeElement;
    deleteEl0.click();

    fixture.detectChanges();

    expect((component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').value).toEqual('id2');
  });

});

@Component({
  template: `
  <app-complex-search
    [category]="category"
    [saveData]="saveData"
    >
  </app-complex-search>`
})
class TestHostComponent {

  @ViewChild(ComplexSearchComponent, { static: true })
  searchComponent: ComplexSearchComponent;

  category: Category = createCategoryArrayData()[0]
  saveData: SaveData = createTestInstanceSaveData()
}
