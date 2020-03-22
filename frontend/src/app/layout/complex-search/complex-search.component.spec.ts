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

describe('ComplexSearchComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;
  let spy: jasmine.SpyObj<ComplexSearchService>;
  let saveData: SaveData

  beforeEach(async(() => {
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearchSave','orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchComponent,
        TestHostComponent,
        ComplexSearchConditionItemComponent,
        ComplexSearchOrderItemComponent,
        DeleteComponent,
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
        }
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    spy = TestBed.get(ComplexSearchService);
    spy.initSaveDataObj.and.returnValue({
      patternName: '',
      category: '',
      isDisclose: false,
      discloseGroups: [],
      ownerID: '',
      conditionData: {
        searchStrings: [],
        displayItemList: [],
        searchConditionList: [],
        orderConditionList: [],
      },
    });
    spy.initConditionDataObj.and.returnValue({
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    });
  });

  it('should create', () => {
    const fb = new FormBuilder();
    component.searchComponent.form = fb.group({
      searchCondition: fb.array([new FormGroup({
        fieldSelected: new FormControl(''),
        conditionValue: new FormControl(''),
        matchTypeSelected: new FormControl(''),
        operatorSelected: new FormControl(''),
      })]),
      orderCondition: fb.array([new FormGroup({
        orderFieldSelected: new FormControl(''),
        orderFieldKeyWordSelected: new FormControl(''),
      })]),
      saveCondition: fb.group({
        patternName: fb.control(""),
        isDisclose: fb.control(""),
        discloseGroups: fb.array([new FormControl('')]),
      }),
    });
    component.searchComponent.isShowDisplayItem = true;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = true;
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should click add condition button', () => {

    component.searchComponent.isShowDisplayItem = false;
    component.searchComponent.isShowOrderCondition = false;
    component.searchComponent.isShowSaveCondition = false;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".push-search-condition"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    buttonEl.click();
    fixture.detectChanges();

    const itemBoxDeArray: DebugElement[] = fixture.debugElement.queryAll(By.directive(ComplexSearchConditionItemComponent));

    expect(itemBoxDeArray.length).toBe(2);
  });

  it('should click add order button', () => {

    component.searchComponent.isShowDisplayItem = false;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = false;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".push-order-condition"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    buttonEl.click();
    fixture.detectChanges();

    const itemBoxDeArray: DebugElement[] = fixture.debugElement.queryAll(By.directive(ComplexSearchOrderItemComponent));

    expect(itemBoxDeArray.length).toBe(2);
  });

  it('should click save button', () => {
    component.searchComponent.isShowDisplayItem = true;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = true;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".complex-search-condition-save-button"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    fixture.detectChanges();

    expect(spy.orderComplexSearchSave).toHaveBeenCalled();
  });

  it('should click search button', () => {

    component.searchComponent.isShowDisplayItem = true;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = true;
    fixture.detectChanges();

    const buttonDe: DebugElement = fixture.debugElement.query(By.css(".complex-search-button"));
    const buttonEl: HTMLSelectElement = buttonDe.nativeElement;
    buttonEl.click();
    fixture.detectChanges();

    expect(spy.orderComplexSearch).toHaveBeenCalled();
  });


  it('should create save data', () => {

    component.searchComponent.isShowDisplayItem = true;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = true;
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

    component.searchComponent.pushSearchCondition();
    component.searchComponent.pushOrderCondition();
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('fieldSelected').setValue('id1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('conditionValue').setValue('value1');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('matchTypeSelected').setValue('match');
    (component.searchComponent.searchConditionFormArray.controls[0] as FormGroup).get('operatorSelected').setValue('and');
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldSelected').setValue('id2');
    (component.searchComponent.orderConditionFormArray.controls[0] as FormGroup).get('orderFieldKeyWordSelected').setValue('asc');
    fixture.detectChanges();

    saveData = component.searchComponent.createSaveData();
    expect(saveData.patternName).toBe('sample pattern name');
    expect(saveData.isDisclose).toBe(true);
    expect(saveData.discloseGroups).toEqual(['id1', 'id2']);
    expect(saveData.conditionData.searchConditionList[0].field).toEqual(component.searchConditionList[0]);
    expect(saveData.conditionData.searchConditionList[0].conditionValue).toEqual('value1');
    expect(saveData.conditionData.searchConditionList[0].matchType).toEqual('match');
    expect(saveData.conditionData.searchConditionList[0].operator).toEqual('and');
    expect(saveData.conditionData.orderConditionList[0].orderField).toEqual(component.orderConditionList[1]);
    expect(saveData.conditionData.orderConditionList[0].orderFieldKeyWord).toEqual('asc');
  });

  it('should delete condition', () => {

    component.searchComponent.isShowDisplayItem = false;
    component.searchComponent.isShowOrderCondition = false;
    component.searchComponent.isShowSaveCondition = false;
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

    component.searchComponent.isShowDisplayItem = false;
    component.searchComponent.isShowOrderCondition = true;
    component.searchComponent.isShowSaveCondition = false;
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
    [displayItemList]="displayItemList"
    [searchConditionList]="searchConditionList"
    [orderConditionList]="orderConditionList"
    [groupList]="groupList"
    >
  </app-complex-search>`
})
class TestHostComponent {

  @ViewChild(ComplexSearchComponent, {static: true})
  searchComponent:ComplexSearchComponent;

  displayItemList = [
    {
      entityName: 'aaa',
      fieldName: 'AAA',
      viewValue: 'aaa-AAA',
      fieldType: 'number',
    },
    {
      entityName: 'bbb',
      fieldName: 'BBB',
      viewValue: 'bbb-BBB',
      fieldType: 'string',
    },
  ];
  searchConditionList: FieldAttr[] =  [
    {
      id: 'id1',
      entityName: 'aaa',
      fieldName: 'AAA',
      viewValue: 'aaa-AAA',
      fieldType: 'number',
    },
    {
      id: 'id2',
      entityName: 'bbb',
      fieldName: 'BBB',
      viewValue: 'bbb-BBB',
      fieldType: 'string',
    },
  ];
  orderConditionList =  [
    {
      id: 'id1',
      entityName: 'aaa',
      fieldName: 'AAA',
      viewValue: 'aaa-AAA',
      fieldType: 'number',
    },
    {
      id: 'id2',
      entityName: 'bbb',
      fieldName: 'BBB',
      viewValue: 'bbb-BBB',
      fieldType: 'string',
    },
  ]
  groupList: Group[] = [
    {
      id: 'id1',
      name: 'name1',
    },
    {
      id: 'id2',
      name: 'name2',
    },
    {
      id: 'id3',
      name: 'name3',
    },
  ];
}
