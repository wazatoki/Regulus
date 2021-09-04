import { Component } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatListModule } from '@angular/material';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import { OperatorUsableConditionsComponent } from '../../complex-search/operator-usable-conditions/operator-usable-conditions.component';
import { CancelComponent } from '../../form/buttons/cancel/cancel.component';

import { OperatorUsableConditionsDialogComponent } from './operator-usable-conditions-dialog.component';

describe('OperatorUsableConditionsDialogComponent', () => {
  let component: OperatorUsableConditionsDialogComponent;
  let fixture: ComponentFixture<OperatorUsableConditionsDialogComponent>;

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
      ],
      providers: [
        {
          provide: MatDialogRef,
          useValue: MatDialogRef
        },
        {
          provide: MAT_DIALOG_DATA, useValue: {}
        },
      ],
    }).overrideModule(BrowserDynamicTestingModule, {
      set: {
        entryComponents: [OperatorUsableConditionsDialogComponent],
      }
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OperatorUsableConditionsDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  fit('should create', () => {
    expect(component).toBeTruthy();
  });
});

@Component({
  template: `
  <div>test host component</div>`
})
class TestHostComponent {


  openDialog() {
    const dialogRef = this.dialog.open(OperatorUsableConditionsDialogComponent, {
      data: {
        operatorUsableConditions: []
      }
    });
  }

  constructor(
    public dialog: MatDialog) { }

}
