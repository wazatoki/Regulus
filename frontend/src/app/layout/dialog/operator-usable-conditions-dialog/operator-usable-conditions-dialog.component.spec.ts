import { Component, DebugElement } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatListModule, MatIconModule } from '@angular/material';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { By } from '@angular/platform-browser';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { OperatorUsableConditionsComponent } from '../../complex-search/operator-usable-conditions/operator-usable-conditions.component';
import { CancelComponent } from '../../form/buttons/cancel/cancel.component';

import { OperatorUsableConditionsDialogComponent } from './operator-usable-conditions-dialog.component';
import { HttpClient } from '@angular/common/http';

@Component({
  template: `
  <div>test host component</div>`
})
class TestHostComponent {


  openDialog() {
    const dialogRef = this.dialog.open(OperatorUsableConditionsDialogComponent, {
      data: {
        title: 'test_title',
        operatorUsableConditions: []
      }
    });
  }

  constructor(
    public dialog: MatDialog) { }

}

describe('OperatorUsableConditionsDialogComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;
  let debugElement: DebugElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        TestHostComponent,
        OperatorUsableConditionsDialogComponent,
        OperatorUsableConditionsComponent,
        CancelComponent,
      ],
      imports: [
        MatDialogModule,
        MatButtonModule,
        MatListModule,
        MatIconModule,
        NoopAnimationsModule,
      ],
      providers: [
        {
          provide: MatDialogRef,
          useValue: MatDialogRef
        },
        {
          provide: MAT_DIALOG_DATA, useValue: {}
        },
        {
          provide: HttpClient, useValue: {}
        }
      ],
    }).overrideModule(BrowserDynamicTestingModule, {
      set: {
        entryComponents: [OperatorUsableConditionsDialogComponent],
      }
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    debugElement = fixture.debugElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    component.openDialog();
    fixture.detectChanges();
    expect(component).toBeTruthy();

  });
});
