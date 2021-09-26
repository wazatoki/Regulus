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
import { ceateTestArray as createTestArrayStaffGroupData } from '../../services/models/group/staff-group.spec';
import { StaffGroupMasterComponent } from './staff-group-master.component';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { of } from 'rxjs';

describe('StaffGroupMasterComponent', () => {
  let component: StaffGroupMasterComponent;
  let fixture: ComponentFixture<StaffGroupMasterComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;
  let testData: StaffGroup[];

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('StaffGroupService', ['delete', 'findByCondition']);
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

    component.dataSource.data = testData;
    fixture.detectChanges();
    expect(element.textContent).toContain('TEST_GROUP_NAME_1');
    expect(element.textContent).toContain('TEST_GROUP_NAME_2');
  });

  it('select item as all checked', () => {

    component.dataSource.data = testData;
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input');
      const checkbox = checkboxList[0];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();
      expect(component.selection.selected.length).toBe(2);

    });
  });

  it('select item as checked', () => {

    component.dataSource.data = testData;
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input');
      const checkbox = checkboxList[1];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();
      expect(component.selection.selected.length).toBe(1);

    });
  });

  it('click delete without select', () => {
    component.deleteItems();
    expect(component.dialog.open).toHaveBeenCalledWith(NoticeDialogComponent, {
      data: { contents: '削除対象が選択されていません。' }
    });
  });

  it('click delete with select', () => {

    const spy: jasmine.SpyObj<StaffGroupService> = TestBed.get(StaffGroupService);
    const stubValue = of(testData);
    spy.delete.and.returnValue(stubValue);

    component.dataSource.data = testData;
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input');
      const checkbox = checkboxList[1];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();

      component.execDeleteItems();
      expect(spy.delete).toHaveBeenCalled();

    });
  });

  it('called api when search execute', () => {
    const spy: jasmine.SpyObj<StaffGroupService> = TestBed.get(StaffGroupService);
    const searchWords = 'aaa bbb ccc';
    const condition = {
      searchStrings: ['aaa', 'bbb', 'ccc'],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
    const data: StaffGroup[] = createTestArrayStaffGroupData();
    spy.findByCondition.and.returnValue(of(data));
    component.search(condition);

    expect(spy.findByCondition).toHaveBeenCalledWith(condition);
  });

});
