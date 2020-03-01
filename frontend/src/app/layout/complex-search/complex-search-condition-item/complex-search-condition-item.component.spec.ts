import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionItemComponent } from './complex-search-condition-item.component';

import { FormsModule, ReactiveFormsModule, FormControl, FormArray, FormGroup } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { Component, DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';

describe('ComplexSearchConditionItemComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionItemComponent,
        TestHostComponent,
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatSelectModule,
        MatRadioModule,
        MatGridListModule,
        NoopAnimationsModule,
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  fit('should click field select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-field-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const selectDe2: DebugElement = fixture.debugElement.query(By.css(".select-match-type"));
    const selectEl2: HTMLSelectElement = selectDe2.nativeElement;
    const matchTypeControll: FormControl = component.formGroup.get('matchTypeSelected') as FormControl;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.option-field-name'));
      const fieldControll: FormControl = component.formGroup.get('fieldSelected') as FormControl;
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(fieldControll.value).toBe(component.fields[0])
      selectDe.triggerEventHandler('selectionChange', null);
      fixture.detectChanges();
      selectEl2.click();
      fixture.detectChanges();

      const inquiryOptions2: DebugElement[] = fixture.debugElement.queryAll(By.css('.option-match-type'));
      inquiryOptions2[2].nativeElement.click()
      expect(matchTypeControll.value).toBe('gt')
    });
  });

  it('should input value select', async () => {

    fixture.detectChanges();
    const inputDe: DebugElement = fixture.debugElement.query(By.css(".input-value"));
    const inputEl: HTMLInputElement = inputDe.nativeElement;
    const conditionValueControll: FormControl = component.formGroup.get('conditionValue') as FormControl;
    inputEl.value = "aaa"
    inputEl.dispatchEvent(new Event('input'));
    fixture.detectChanges();

    expect(conditionValueControll.value).toBe("aaa")
  });

  it('should click match type select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-match-type"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const matchTypeControll: FormControl = component.formGroup.get('matchTypeSelected') as FormControl;
    
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(matchTypeControll.value).toBe("match")
    });
  });

  it('should click operator radio button', async () => {

    //component.operatorSelected = '';
    const radioDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".radio-operator-option .mat-radio-input"));
    const radioEl: HTMLInputElement = radioDe[0].nativeElement;
    const operatorSelectedControll: FormControl = component.formGroup.get('operatorSelected') as FormControl;
    radioEl.click()
    fixture.detectChanges();
    expect(operatorSelectedControll.value).toBe('and')
  });

});

@Component({
  template: `
  <app-complex-search-condition-item
    [fields]="fields" [formGroup]="formGroup">
  </app-complex-search-condition-item>`
})
class TestHostComponent {

  formGroup: FormGroup = new FormGroup({
    fieldSelected: new FormControl(''),
    conditionValue: new FormControl(''),
    matchTypeSelected: new FormControl(''),
    operatorSelected: new FormControl(''),
  });

  fields = [
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
}
