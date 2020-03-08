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
import { ComplexSearchConditionItemComponent } from './complex-search-condition-item/complex-search-condition-item.component'
import { ComplexSearchOrderItemComponent } from "./complex-search-order-item/complex-search-order-item.component"
import { DebugElement, Component, ViewChild } from '@angular/core';
import { By } from '@angular/platform-browser';

describe('ComplexSearchComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchComponent,
        TestHostComponent,
        ComplexSearchConditionItemComponent,
        ComplexSearchOrderItemComponent,
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
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
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
      saveCondition: fb.group({}),
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

  fit('should click add order button', () => {

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
});

@Component({
  template: `
  <app-complex-search
    [displayItemList]="displayItemList"
    [searchConditionList]="searchConditionList"
    [orderConditionList]="orderConditionList"
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
  searchConditionList =  [
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
  orderConditionList =  [
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
  ]
}
