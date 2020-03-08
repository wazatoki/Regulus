import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchOrderItemComponent } from './complex-search-order-item.component';

import { FormsModule, ReactiveFormsModule, FormControl, FormGroup } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
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

  fit('should click order field select', async () => {

    const selectDe: DebugElement = fixture.debugElement.query(By.css(".select-order-field-name"));
    const selectEl: HTMLSelectElement = selectDe.nativeElement;
    const orderFieldFormControll: FormControl = component.formGroup.get('orderFieldSelected') as FormControl;
    selectEl.click();
    fixture.detectChanges();

    await fixture.whenStable().then(() => {
      const inquiryOptions = fixture.debugElement.queryAll(By.css('.mat-option-text'));
      inquiryOptions[0].nativeElement.click()

      fixture.detectChanges();
      expect(orderFieldFormControll.value).toBe(component.fields[0])
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

});


@Component({
  template: `
  <app-complex-search-order-item
    [fields]="fields" [formGroup]="formGroup">
  </app-complex-search-order-item>`
})
class TestHostComponent {

  formGroup: FormGroup = new FormGroup({
    orderFieldSelected: new FormControl(''),
    orderFieldKeyWordSelected: new FormControl(''),
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
  ]
}

