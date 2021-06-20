import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroup } from 'src/app/services/models/group/staff-group';

import { StaffGroupSearchComponent } from './staff-group-search.component';
import { createTestInstance1 as createTestInstanceStaffGroup, ceateTestArray as createTestArrayStaffGroupData } from '../../services/models/group/staff-group.spec';
import { of, Subject } from 'rxjs';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ConditionData } from 'src/app/services/models/search/condition-data';

describe('StaffGroupSearchComponent', () => {
  let component: StaffGroupSearchComponent;
  let fixture: ComponentFixture<StaffGroupSearchComponent>;
  let staffGroupServiceSpy: jasmine.SpyObj<StaffGroupService>
  let complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService>

  beforeEach(async(() => {

    const staffGroupServiceSpy = jasmine.createSpyObj('StaffGroupService', ['findByCondition']);
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

    TestBed.configureTestingModule({
      declarations: [ StaffGroupSearchComponent ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: StaffGroupService, useValue: staffGroupServiceSpy },
        { provide: ComplexSearchService, useValue: complexSearchServiceSpy },
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    complexSearchServiceSpy = TestBed.get(ComplexSearchService);
    complexSearchServiceSpy.initConditionDataObj.and.returnValue({
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    });
    complexSearchServiceSpy.complexSearchOrdered$ = new Subject<ConditionData>().asObservable();

    fixture = TestBed.createComponent(StaffGroupSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('called api when onSearch execute', () => {
    staffGroupServiceSpy = TestBed.get(StaffGroupService);
    const searchWords = 'aaa bbb ccc';
    const condition = {
      searchStrings: ['aaa','bbb','ccc'],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
    const data: StaffGroup[] = createTestArrayStaffGroupData();
    staffGroupServiceSpy.findByCondition.and.returnValue(of(data))
    component.onSearch(searchWords);

    component.fetched.subscribe( (res: StaffGroup[]) => {
      expect(res).toBe(data);
    });

    expect(staffGroupServiceSpy.findByCondition).toHaveBeenCalledWith(condition);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
