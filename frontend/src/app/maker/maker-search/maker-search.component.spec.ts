import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { LayoutModule } from '../../layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';
import { MakerSearchComponent } from './maker-search.component';
import { MakerService } from '../../services/api/maker.service';
import { MakerCondition } from '../../services/models/maker/maker-condition';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { SearchComponent } from 'src/app/layout/search/search.component';
import { Maker } from '../../services/models/maker/maker';
import { of,Subject } from 'rxjs';
import { ConditionData } from 'src/app/services/models/search/condition-data';

describe('MakerSearchComponent', () => {
  let component: MakerSearchComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<MakerSearchComponent>;
  let makerServiceSpy: jasmine.SpyObj<MakerService>;
  let complexSearchServiceSpy:jasmine.SpyObj<ComplexSearchService>;


  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('MakerService', ['findByCondition']);
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearchSave','orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

    TestBed.configureTestingModule({
      declarations: [ MakerSearchComponent ],
      imports: [
        LayoutModule,
        MatButtonModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: MakerService, useValue: spy },
        { provide: ComplexSearchService, useValue: complexSearchServiceSpy}
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    complexSearchServiceSpy = TestBed.get(ComplexSearchService);
    complexSearchServiceSpy.initSaveDataObj.and.returnValue({
      patternName: '',
      category: '',
      isDisclose: false,
      discloseGroups: [],
      ownerID: '',
      conditionData: {
        searchStrings: [],
        displayItemList: [],
        searchConditionList: [],
        orderConditionList: [],
      },
    });
    complexSearchServiceSpy.initConditionDataObj.and.returnValue({
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    });
    complexSearchServiceSpy.complexSearchOrdered$ = new Subject<ConditionData>().asObservable();

    fixture = TestBed.createComponent(MakerSearchComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement; 
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should onSearch set keywords to makerCondition.searchStrings', () => {
    makerServiceSpy = TestBed.get(MakerService);
    complexSearchServiceSpy = TestBed.get(ComplexSearchService);
    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    const stubValue = of(testData)
    const condition = {
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
    makerServiceSpy.findByCondition.and.returnValue(stubValue);
    complexSearchServiceSpy.initConditionDataObj.and.returnValue(condition);
    component.fetched.subscribe( (data: Maker[]) => {
      expect(data).toBe(testData);
    })
    component.onSearch('search word');
    condition.searchStrings = ['search', 'word']
    expect(makerServiceSpy.findByCondition).toHaveBeenCalledWith(condition);
  });
});
