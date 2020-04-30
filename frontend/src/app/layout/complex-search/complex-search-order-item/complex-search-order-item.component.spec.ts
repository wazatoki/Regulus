import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchOrderItemComponent } from './complex-search-order-item.component';
import { FieldAttr } from '../../../services/models/search/field-attr';
import { DeleteComponent } from '../../form/buttons/delete/delete.component';

import { FormsModule, ReactiveFormsModule, FormControl, FormGroup } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { DebugElement, Component } from '@angular/core';
import { By } from '@angular/platform-browser';

describe('ComplexSearchOrderItemComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchOrderItemComponent,
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

  it('should click order field select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-order-field-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const orderFieldFormControll: FormControl = component.formGroup.get('orderFieldSelected') as FormControl;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(orderFieldFormControll.value).toBe(component.fields[0].id)
    });
  });

  it('should click order key word radio button', async () => {

    const radioDe: DebugElement[] = fixture.debugElement.queryAll(By.css(".radio-group-order-field-key-word-select .mat-radio-input"));
    const radioEl: HTMLInputElement = radioDe[0].nativeElement;
    const orderFieldKeyWordFormControll: FormControl = component.formGroup.get('orderFieldKeyWordSelected') as FormControl;
    radioEl.click()
    fixture.detectChanges();
    expect(orderFieldKeyWordFormControll.value).toBe('asc')
  });

  it('should click delete button', () => {

    const deleteDe: DebugElement = fixture.debugElement.query(By.css('.delete-button button'));
    const deleteEl: HTMLButtonElement = deleteDe.nativeElement;
    deleteEl.click()
    fixture.detectChanges();
    expect(component.onDelete()).toEqual("onDelete called");
  });

});


@Component({
  template: `
  <app-complex-search-order-item
    [fields]="fields"
    [formGroup]="formGroup"
    (onDelete)="onDelete()">
  </app-complex-search-order-item>`
})
class TestHostComponent {

  onDelete(): string{
    return 'onDelete called';
  }

  formGroup: FormGroup = new FormGroup({
    orderFieldSelected: new FormControl(''),
    orderFieldKeyWordSelected: new FormControl(''),
  });

  fields: FieldAttr[] = [
    {
      id: 'id1',
      viewValue: 'aaa-AAA',
      fieldType: 'number',
    },
    {
      id: 'id2',
      viewValue: 'bbb-BBB',
      fieldType: 'string',
    },
  ]
}

