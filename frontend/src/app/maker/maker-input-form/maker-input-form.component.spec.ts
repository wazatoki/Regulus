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
import { CancelComponent } from '../../layout/form/buttons/cancel/cancel.component';
import { ClearComponent } from '../../layout/form/buttons/clear/clear.component';
import { SubmitComponent } from '../../layout/form/buttons/submit/submit.component';
import { NoticeDialogComponent } from '../../layout/dialog/notice-dialog/notice-dialog.component';
import { MakerService } from '../../services/api/maker.service';
import { Maker } from '../../services/models/maker/maker';
import { of, Observable, throwError } from 'rxjs';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';

describe('MakerInputFormComponent', () => {
  let component: MakerInputFormComponent;
  let fixture: ComponentFixture<MakerInputFormComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    const matDialogRefSpy = jasmine.createSpyObj('MatDialogRef', ['close']);
    const matMakerServiceSpy = jasmine.createSpyObj('MakerService', ['add']);

    TestBed.configureTestingModule({
      declarations: [
        MakerInputFormComponent,
      ],
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
          useValue: matDialogRefSpy
        },
        {
          provide: MAT_DIALOG_DATA, useValue: {} 
        },
        {
          provide: MakerService,
          useValue: matMakerServiceSpy
        },
      ],
    }).overrideModule(BrowserDynamicTestingModule, {
      set: {
        entryComponents: [ NoticeDialogComponent ],
      }
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

  it('should close form at click cancel button', () => {
    dbElement = fixture.debugElement;
    const buttonDebugElement: DebugElement = dbElement.query(By.directive(CancelComponent));
    buttonDebugElement.triggerEventHandler('clicked', null);
    fixture.detectChanges();
    expect(component.dialogRef.close).toHaveBeenCalled();
  });

  it('should clear form at click clear button', () => {
    dbElement = fixture.debugElement;
    const inputElement: HTMLInputElement = dbElement.query(By.css('input[type="text"]')).nativeElement;
    const buttonDebugElement: DebugElement = dbElement.query(By.directive(ClearComponent));
    inputElement.value = 'test value';
    inputElement.dispatchEvent(new Event('input'));
    inputElement.dispatchEvent(new Event('blur'));
    buttonDebugElement.triggerEventHandler('clicked', null);
    fixture.detectChanges();
    expect(inputElement.value).toEqual('');
  });

  it('should save form date at click save button', () => {
    const testData: Maker = { id: 'testid', name: 'Test Maker' };
    const spy: jasmine.SpyObj<MakerService> = TestBed.get(MakerService);
    const stubValue = of(testData);
    spy.add.and.returnValue(stubValue);

    dbElement = fixture.debugElement;
    const inputElement: HTMLInputElement = dbElement.query(By.css('input[type="text"]')).nativeElement;
    const buttonDebugElement: DebugElement = dbElement.query(By.directive(SubmitComponent));
    inputElement.value = 'test value';
    inputElement.dispatchEvent(new Event('input'));
    inputElement.dispatchEvent(new Event('blur'));
    buttonDebugElement.triggerEventHandler('clicked', null);
    fixture.detectChanges();
    expect(component.makerService.add).toHaveBeenCalled();
  });

  it('should save form date at click save button with error', () => {
    const testData: Maker = { id: 'testid', name: 'Test Maker' };
    const spy: jasmine.SpyObj<MakerService> = TestBed.get(MakerService);
    spy.add.and.returnValue(throwError({status: 404}));

    dbElement = fixture.debugElement;
    const inputElement: HTMLInputElement = dbElement.query(By.css('input[type="text"]')).nativeElement;
    const buttonDebugElement: DebugElement = dbElement.query(By.directive(SubmitComponent));
    inputElement.value = 'test value';
    inputElement.dispatchEvent(new Event('input'));
    inputElement.dispatchEvent(new Event('blur'));
    buttonDebugElement.triggerEventHandler('clicked', null);
    fixture.detectChanges();
    expect(component.makerService.add).toHaveBeenCalled();
  });
});
