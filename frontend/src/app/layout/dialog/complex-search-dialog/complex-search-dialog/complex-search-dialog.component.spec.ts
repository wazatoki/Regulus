import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatCardModule } from '@angular/material/card';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDialog, MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';

import { ComplexSearchDialogComponent } from './complex-search-dialog.component';
import { ComplexSearchComponent } from '../../../complex-search/complex-search.component';
import { ComplexSearchConditionItemComponent } from '../../../complex-search/complex-search-condition-item/complex-search-condition-item.component'
import { ComplexSearchOrderItemComponent } from '../../../complex-search/complex-search-order-item/complex-search-order-item.component';
import { CancelComponent } from '../../../form/buttons/cancel/cancel.component';
import { DeleteComponent } from '../../../form/buttons/delete/delete.component';
import { Group } from '../../../../services/models/group/group';
import { FieldAttr } from '../../../../services/models/search/field-attr';
import { DebugElement, Component, ViewChild } from '@angular/core';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import { HttpClient } from '@angular/common/http';
import { ClearComponent } from 'src/app/layout/form/buttons/clear/clear.component'

describe('ComplexSearchDialogComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        TestHostComponent,
        ComplexSearchDialogComponent,
        ComplexSearchComponent,
        ComplexSearchConditionItemComponent,
        ComplexSearchOrderItemComponent,
        CancelComponent,
        DeleteComponent,
        ClearComponent,
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
        MatDialogModule,
        DragDropModule,
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
        entryComponents: [ ComplexSearchDialogComponent ],
      }
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    component.openDialog();
    expect(component).toBeTruthy();
  });
});

@Component({
  template: `
  <div>test host component</div>`
})
class TestHostComponent {


  openDialog() {
    const dialogRef = this.dialog.open(ComplexSearchDialogComponent, {
      data: {
        displayItemList: [
          {
            id: 'id1',
            entityName: 'aaa',
            fieldName: 'AAA',
            viewValue: 'aaa-AAA',
            fieldType: 'number',
          },
          {
            id: 'id2',
            entityName: 'bbb',
            fieldName: 'BBB',
            viewValue: 'bbb-BBB',
            fieldType: 'string',
          },
        ],
        searchConditionList: [
          {
            id: 'id1',
            entityName: 'aaa',
            fieldName: 'AAA',
            viewValue: 'aaa-AAA',
            fieldType: 'number',
          },
          {
            id: 'id2',
            entityName: 'bbb',
            fieldName: 'BBB',
            viewValue: 'bbb-BBB',
            fieldType: 'string',
          },
        ],
        orderConditionList: [
          {
            id: 'id1',
            entityName: 'aaa',
            fieldName: 'AAA',
            viewValue: 'aaa-AAA',
            fieldType: 'number',
          },
          {
            id: 'id2',
            entityName: 'bbb',
            fieldName: 'BBB',
            viewValue: 'bbb-BBB',
            fieldType: 'string',
          },
        ],
        isShowDisplayItem: true,
        isShowOrderCondition: true,
        isShowSaveCondition: true,
        groupList: [
          { id: 'id1', name: 'name1' },
          { id: 'id2', name: 'name2' },
        ],
      }
    });
  }

  constructor(
    public dialog: MatDialog) { }

}
