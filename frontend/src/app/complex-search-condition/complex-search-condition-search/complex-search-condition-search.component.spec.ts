import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search.component';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { SaveData } from 'src/app/services/models/search/save-data';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { of, Subject } from 'rxjs';

import { createTestArray } from 'src/app/services/models/search/save-data.spec';

describe('ComplexSearchConditionSearchComponent', () => {
  let component: ComplexSearchConditionSearchComponent;
  let fixture: ComponentFixture<ComplexSearchConditionSearchComponent>;
  let complexSearchConditionServiceSpy: jasmine.SpyObj<ComplexSearchConditionService>;
  let complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService>

  beforeEach(async(() => {

    const complexSearchConditionServiceSpy = jasmine.createSpyObj('ComplexSearchConditionService', ['findByCondition']);
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

    TestBed.configureTestingModule({
      declarations: [ ComplexSearchConditionSearchComponent ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: ComplexSearchConditionService, useValue: complexSearchConditionServiceSpy },
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

    fixture = TestBed.createComponent(ComplexSearchConditionSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
