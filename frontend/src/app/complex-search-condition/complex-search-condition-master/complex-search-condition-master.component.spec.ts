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
import { SaveData } from 'src/app/services/models/search/save-data';
import { DebugElement } from '@angular/core';
import { NoticeDialogComponent } from 'src/app/layout/dialog/notice-dialog/notice-dialog.component';
import { of } from 'rxjs';

describe('ComplexSearchConditionMasterComponent', () => {
  let component: ComplexSearchConditionMasterComponent;
  let fixture: ComponentFixture<ComplexSearchConditionMasterComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;
  let testData: SaveData[];

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
    dbElement = fixture.debugElement;
    element = dbElement.nativeElement;
    fixture.detectChanges();

    testData = [
      {
        id: 'testID1',
        ownerID: 'testOwnerID1',
        patternName: 'testName1',
        category: 'testCategory1',
        conditionData: {
          displayItemList: [],
          orderConditionList: [],
          searchConditionList: [],
          searchStrings: [],
        },
        discloseGroups: [],
        isDisclose: false,
        owner: {
          id: '',
          name: 'testStaffName1',
        },
      },
      {
        id: 'testID2',
        ownerID: 'testOwnerID2',
        patternName: 'testName2',
        category: 'testCategory2',
        conditionData: {
          displayItemList: [],
          orderConditionList: [],
          searchConditionList: [],
          searchStrings: [],
        },
        discloseGroups: [],
        isDisclose: false,
        owner: {
          id: '',
          name: 'testStaffName1',
        },
      },
    ];
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain table', () => {

    expect(element.textContent).toContain('検索条件パターン名');

    component.onFetchedSearchConditions(testData);
    fixture.detectChanges();
    expect(element.textContent).toContain('testName1');
    expect(element.textContent).toContain('testName2');
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

  it('select item as checked', () => {

    component.onFetchedSearchConditions(testData);
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input')
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
    })
  });

  it('click delete with select', () => {

    const spy: jasmine.SpyObj<ComplexSearchConditionService> = TestBed.get(ComplexSearchConditionService);
    const stubValue = of(testData);
    spy.delete.and.returnValue(stubValue);
    
    component.onFetchedSearchConditions(testData);
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input')
      const checkbox = checkboxList[1];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();

      component.deleteItems();
      expect(spy.delete).toHaveBeenCalled();

    });
  });

});
