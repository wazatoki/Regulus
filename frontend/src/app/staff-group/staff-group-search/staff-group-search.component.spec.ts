import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';
import { StaffGroup } from 'src/app/services/models/group/staff-group';

import { StaffGroupSearchComponent } from './staff-group-search.component';
import { of, Subject } from 'rxjs';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { SelectSearchConditionComponent } from 'src/app/layout/form/buttons/select-search-condition/select-search-condition.component';
import { MatDialog } from '@angular/material';
import { LoginService } from 'src/app/services/api/login.service';

describe('StaffGroupSearchComponent', () => {
  let component: StaffGroupSearchComponent;
  let fixture: ComponentFixture<StaffGroupSearchComponent>;
  let complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService>;

  beforeEach(async(() => {

    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

    TestBed.configureTestingModule({
      declarations: [ StaffGroupSearchComponent],
      imports: [
        LayoutModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: ComplexSearchService, useValue: complexSearchServiceSpy },
        { provide: MatDialog },
        { provide: LoginService },
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

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
