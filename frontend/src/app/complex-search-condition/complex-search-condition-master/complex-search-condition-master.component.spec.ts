import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionMasterComponent } from './complex-search-condition-master.component';
import { ComplexSearchConditionSearchComponent } from '../complex-search-condition-search/complex-search-condition-search.component';
import { LayoutModule } from 'src/app/layout/layout.module';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatPaginatorModule } from '@angular/material/paginator';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDialogModule, MatDialog } from '@angular/material/dialog';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { HttpClient } from '@angular/common/http';

describe('ComplexSearchConditionMasterComponent', () => {
  let component: ComplexSearchConditionMasterComponent;
  let fixture: ComponentFixture<ComplexSearchConditionMasterComponent>;

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('ComplexSearchConditionService', ['delete']);
    const dialogspy = jasmine.createSpyObj('MatDialog', ['open']);

    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchConditionMasterComponent,
        ComplexSearchConditionSearchComponent
      ],
      imports: [
        BrowserAnimationsModule,
        LayoutModule,
        MatTableModule,
        MatListModule,
        MatDialogModule,
        MatCheckboxModule,
        MatPaginatorModule,
      ],
      providers: [
        { provide: ComplexSearchConditionService, useValue: spy },
        { provide: MatDialog, useValue: dialogspy },
        { provide: HttpClient, useValue: {} },
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchConditionMasterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
