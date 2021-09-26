import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search.component';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { Subject } from 'rxjs';
import { MatButtonModule, MatDialog } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginService } from 'src/app/services/api/login.service';

describe('ComplexSearchConditionSearchComponent', () => {
  let component: ComplexSearchConditionSearchComponent;
  let fixture: ComponentFixture<ComplexSearchConditionSearchComponent>;
  const complexSearchConditionServiceSpy: jasmine.SpyObj<ComplexSearchConditionService> = jasmine.createSpyObj(
    'ComplexSearchConditionService',
    ['findByCondition']);
  const complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService> = jasmine.createSpyObj(
    'ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

  complexSearchServiceSpy.initSaveDataObj.and.returnValue({
    id: '',
    patternName: '',
    category: null,
    isDisclose: false,
    discloseGroups: [],
    ownerID: '',
    conditionData: {
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    },
    owner: {
      id: '',
      name: '',
      operatorUsableConditions: [],
    }
  });
  complexSearchServiceSpy.initConditionDataObj.and.returnValue({
    searchStrings: [],
    displayItemList: [],
    searchConditionList: [],
    orderConditionList: [],
  });
  complexSearchServiceSpy.complexSearchOrdered$ = new Subject<ConditionData>().asObservable();

  beforeEach(async(() => {

    TestBed.configureTestingModule({
      declarations: [ComplexSearchConditionSearchComponent],
      imports: [
        BrowserAnimationsModule,
        LayoutModule,
        FlexLayoutModule,
        MatButtonModule,
      ],
      providers: [
        { provide: ComplexSearchConditionService, useValue: complexSearchConditionServiceSpy },
        { provide: ComplexSearchService, useValue: complexSearchServiceSpy },
        { provide: MatDialog, useValue: MatDialog },
        { provide: LoginService, useValue: LoginService },
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchConditionSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
