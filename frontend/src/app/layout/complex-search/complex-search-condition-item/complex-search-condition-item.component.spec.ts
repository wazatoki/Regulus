import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionItemComponent } from './complex-search-condition-item.component';

import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';

describe('ComplexSearchConditionItemComponent', () => {
  let component: ComplexSearchConditionItemComponent;
  let fixture: ComponentFixture<ComplexSearchConditionItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ComplexSearchConditionItemComponent],
      imports: [
        FormsModule,
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
    fixture = TestBed.createComponent(ComplexSearchConditionItemComponent);
    component = fixture.componentInstance;
    component.fields = [
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

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should click field select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-field-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(component.fieldSelected).toBe(component.fields[0])
      selectDe.triggerEventHandler('selectionChange', null);
      expect(component.matchTypes).toEqual(component.matchTypesForNumber)
    });
  });

  it('should input value select', async () => {

    fixture.detectChanges();
    const inputDe: DebugElement = fixture.debugElement.query(By.css(".input-value"));
    const inputEl: HTMLInputElement = inputDe.nativeElement;
    inputEl.value = "aaa"
    inputEl.dispatchEvent(new Event('input'));
    fixture.detectChanges();

    expect(component.conditionValue).toBe("aaa")
  });

  it('should click match type select', async () => {

    component.matchTypeSelected = "";
    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-match-type"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(component.matchTypeSelected).toBe("match")
    });
  });

  it('should click operator radio button', async () => {

    component.operatorSelected = '';
    const radioDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".radio-operator-option .mat-radio-input"));
    const radioEl: HTMLInputElement = radioDe[0].nativeElement;
    radioEl.click()
    fixture.detectChanges();
    expect(component.operatorSelected).toBe('and')
  });

});
