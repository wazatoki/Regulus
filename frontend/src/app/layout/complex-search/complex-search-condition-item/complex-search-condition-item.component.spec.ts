import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionItemComponent } from './complex-search-condition-item.component';
import { FieldAttr } from '../../../services/models/search/field-attr';
import { DeleteComponent } from '../../form/buttons/delete/delete.component';

import { FormsModule, ReactiveFormsModule, FormControl, FormGroup } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatCheckboxModule } from '@angular/material/checkbox';

import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { Component, DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';

@Component({
  template: `
  <app-complex-search-condition-item
    [fields]="fields"
    [formGroup]="formGroup"
    (deleted)="deleted()">
  </app-complex-search-condition-item>`
})
class TestHostComponent {

  formGroup: FormGroup;
  fields: FieldAttr[];

  deleted(): string {
    return 'deleted called';
  }

  constructor() {
    this.formGroup = new FormGroup({
      fieldSelected: new FormControl(''),
      conditionValue: new FormControl(''),
      matchTypeSelected: new FormControl(''),
      operatorSelected: new FormControl(''),
    });

    this.formGroup.setValue({
      fieldSelected: 'id1',
      conditionValue: 'abcdefg',
      matchTypeSelected: 'unmatch',
      operatorSelected: 'or',
    });

    this.fields = [
      {
        id: 'id1',
        viewValue: 'aaa-AAA',
        fieldType: { value: 'number' },
        optionItems: null,
      },
      {
        id: 'id2',
        viewValue: 'bbb-BBB',
        fieldType: { value: 'string' },
        optionItems: null,
      },
    ];
  }

}

describe('ComplexSearchConditionItemComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionItemComponent,
        TestHostComponent,
        DeleteComponent,
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatSelectModule,
        MatRadioModule,
        MatGridListModule,
        MatButtonModule,
        MatCheckboxModule,
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

  it('should click field select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css('.select-field-name'));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const selectDe2: DebugElement = fixture.debugElement.query(By.css('.select-match-type'));
    const selectEl2: HTMLSelectElement = selectDe2.nativeElement;
    const matchTypeControll: FormControl = component.formGroup.get('matchTypeSelected') as FormControl;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.option-field-name'));
      const fieldControll: FormControl = component.formGroup.get('fieldSelected') as FormControl;
      inquiryOptions[0].nativeElement.click();

      fixture.detectChanges();
      expect(fieldControll.value).toBe(component.fields[0].id);
      selectDe.triggerEventHandler('selectionChange', null);
      fixture.detectChanges();
      selectEl2.click();
      fixture.detectChanges();

      const inquiryOptions2: DebugElement[] = fixture.debugElement.queryAll(By.css('.option-match-type'));
      inquiryOptions2[2].nativeElement.click();
      expect(matchTypeControll.value).toBe('gt');
    });
  });

  it('should input value select', async () => {

    fixture.detectChanges();
    const inputDe: DebugElement = fixture.debugElement.query(By.css('.input-value'));
    const inputEl: HTMLInputElement = inputDe.nativeElement;
    const conditionValueControll: FormControl = component.formGroup.get('conditionValue') as FormControl;
    inputEl.value = 'aaa';
    inputEl.dispatchEvent(new Event('input'));
    fixture.detectChanges();

    expect(conditionValueControll.value).toBe('aaa');
  });

  it('should click match type select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css('.select-match-type'));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const matchTypeControll: FormControl = component.formGroup.get('matchTypeSelected') as FormControl;

    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click();

      fixture.detectChanges();
      expect(matchTypeControll.value).toBe('match');
    });
  });

  it('should click operator radio button', async () => {

    const radioDe: DebugElement[] = fixture.debugElement.queryAll(By.css('.radio-operator-option .mat-radio-input'));
    const radioEl: HTMLInputElement = radioDe[0].nativeElement;
    const operatorSelectedControll: FormControl = component.formGroup.get('operatorSelected') as FormControl;
    radioEl.click();
    fixture.detectChanges();
    expect(operatorSelectedControll.value).toBe('and');
  });

  it('should click delete button', () => {

    const deleteDe: DebugElement = fixture.debugElement.query(By.css('.delete-button button'));
    const deleteEl: HTMLButtonElement = deleteDe.nativeElement;
    deleteEl.click();
    fixture.detectChanges();
    expect(component.deleted()).toEqual('deleted called');
  });

});
