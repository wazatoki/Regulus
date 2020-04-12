import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionInputFormComponent } from './complex-search-condition-input-form.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FormsModule, ReactiveFormsModule, FormBuilder, FormGroup } from '@angular/forms';
import { MatDialogModule, MatDialog } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { MatButtonModule } from '@angular/material/button';
import { ViewChild, Component } from '@angular/core';

import { ceateTestArrayForMasterMaintenanceTest } from 'src/app/services/models/search/category.spec'
import { Category } from 'src/app/services/models/search/category';

describe('ComplexSearchConditionInputFormComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;
  // let component: ComplexSearchConditionInputFormComponent;
  // let fixture: ComponentFixture<ComplexSearchConditionInputFormComponent>;

  beforeEach(async(() => {
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'updateSearchCondition', 'addSearchCondition']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionInputFormComponent,
        TestHostComponent,
      ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
        FormsModule,
        ReactiveFormsModule,
        MatDialogModule,
        MatCardModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatCheckboxModule,
        MatSelectModule,
        MatRadioModule,
        MatGridListModule,
        DragDropModule,
        NoopAnimationsModule,
      ],
      providers: [
        {
          provide: FormBuilder,
          useValue: new FormBuilder()
        },
        {
          provide: ComplexSearchService,
          useValue: complexSearchServiceSpy
        },
        {
          provide: MatDialog,
          useValue: {}
        },
      ],
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
});


@Component({
  template: `
  <app-complex-search-condition-input-form
    [categories]="categories"
    [displayItemList]="displayItemList"
    [searchConditionList]="searchConditionList"
    [orderConditionList]="orderConditionList"
    [groupList]="groupList"
    [saveData]="saveData"
  >
  </app-complex-search-condition-input-form>`
})
class TestHostComponent {

  @ViewChild(ComplexSearchConditionInputFormComponent, {static: true})
  complexSearchConditionInputFormComponent: ComplexSearchConditionInputFormComponent;

  categories: Category[] = ceateTestArrayForMasterMaintenanceTest();
}

