import { HttpClient } from '@angular/common/http';
import { DebugElement } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatCheckboxModule, MatDialog, MatDialogModule, MatListModule, MatPaginatorModule, MatTableModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { LayoutModule } from 'src/app/layout/layout.module';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroupSearchComponent } from '../staff-group-search/staff-group-search.component';
import { StaffGroup } from 'src/app/services/models/group/staff-group';
import { ceateTestArray } from 'src/app/services/models/group/staff-group.spec';

import { StaffGroupMasterComponent } from './staff-group-master.component';

describe('StaffGroupMasterComponent', () => {
  let component: StaffGroupMasterComponent;
  let fixture: ComponentFixture<StaffGroupMasterComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;
  let testData: StaffGroup[];

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('StaffGroupService', ['delete']);
    const dialogspy = jasmine.createSpyObj('MatDialog', ['open']);

    TestBed.configureTestingModule({
      declarations: [ 
        StaffGroupMasterComponent,
        StaffGroupSearchComponent,
      ],
      imports: [
        BrowserAnimationsModule,
        LayoutModule,
        RouterTestingModule,
        MatTableModule,
        MatListModule,
        MatDialogModule,
        MatCheckboxModule,
        MatPaginatorModule,
      ],
      providers: [
        { provide: StaffGroupService, useValue: spy },
        { provide: MatDialog, useValue: dialogspy },
        { provide: HttpClient, useValue: {} },
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StaffGroupMasterComponent);
    component = fixture.componentInstance;
    dbElement = fixture.debugElement;
    element = dbElement.nativeElement;
    fixture.detectChanges();
    testData = ceateTestArray();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain table', () => {

    expect(element.textContent).toContain('グループ名称');

    component.onFetchedSearchConditions(testData);
    fixture.detectChanges();
    expect(element.textContent).toContain('TEST_GROUP_NAME_1');
    expect(element.textContent).toContain('TEST_GROUP_NAME_2');
  });

  it('select item as all checked', () => {

    component.onFetchedSearchConditions(testData);
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input')
      const checkbox = checkboxList[0];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();
      expect(component.selection.selected.length).toBe(2);

    });
  });
});
