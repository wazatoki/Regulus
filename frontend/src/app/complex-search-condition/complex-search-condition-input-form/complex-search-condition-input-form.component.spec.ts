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
import { SearchCondition } from 'src/app/services/models/search/search-condition';
import { OrderCondition } from 'src/app/services/models/search/order-condition';
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
  searchConditionList: SearchCondition[] = createSearchConditionArray();
  orderConditionList: OrderCondition[] = createOrderConditionArray();
  groupList: Group[] = createGroupArray();
  saveData: SaveData = createSaveData();
}

