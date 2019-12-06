import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';

import { MakerInputFormComponent } from './maker-input-form.component';
import { LayoutModule } from '../../layout/layout.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';

describe('MakerInputFormComponent', () => {
  let component: MakerInputFormComponent;
  let fixture: ComponentFixture<MakerInputFormComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakerInputFormComponent ],
      imports: [
        BrowserAnimationsModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatDialogModule,
        MatGridListModule,
        LayoutModule,
      ],
      providers: [
        {
          provide: MatDialogRef,
          useValue: { close: (dialogResult: any) => { } }
        },
        {
          provide: MAT_DIALOG_DATA, useValue: {} 
        },
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerInputFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should not dispaly varidation alert at created', () => {
    dbElement = fixture.debugElement;
    element = dbElement.nativeElement; 
    expect(element.textContent).not.toContain('製造販売業者名称は必須項目です。');
  });

  it('should dispaly varidation alert at blue with inut is null', () => {
    dbElement = fixture.debugElement;
    const inputElement: HTMLInputElement = dbElement.query(By.css('input[type="text"]')).nativeElement;
    inputElement.dispatchEvent(new Event('blur'));
    fixture.detectChanges();
    element = dbElement.nativeElement; 
    expect(element.textContent).toContain('製造販売業者名称は必須項目です。');
  });

  it('should not dispaly varidation alert at input value', () => {
    dbElement = fixture.debugElement;
    const inputElement: HTMLInputElement = dbElement.query(By.css('input[type="text"]')).nativeElement;
    inputElement.value = 'test value';
    inputElement.dispatchEvent(new Event('input'));
    inputElement.dispatchEvent(new Event('blur'));
    fixture.detectChanges();
    element = dbElement.nativeElement; 
    expect(element.textContent).not.toContain('製造販売業者名称は必須項目です。');
  });
});
