import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { StaffGroupSearchComponent } from './staff-group-search.component';
import { Subject } from 'rxjs';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { MatDialog } from '@angular/material';
import { LoginService } from 'src/app/services/api/login.service';

describe('StaffGroupSearchComponent', () => {
  let component: StaffGroupSearchComponent;
  let fixture: ComponentFixture<StaffGroupSearchComponent>;
  const complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService> = jasmine.createSpyObj(
    'ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);
  complexSearchServiceSpy.initConditionDataObj.and.returnValue({
    searchStrings: [],
    displayItemList: [],
    searchConditionList: [],
    orderConditionList: [],
  });
  complexSearchServiceSpy.complexSearchOrdered$ = new Subject<ConditionData>().asObservable();

  beforeEach(async(() => {

    TestBed.configureTestingModule({
      declarations: [StaffGroupSearchComponent],
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
    fixture = TestBed.createComponent(StaffGroupSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
