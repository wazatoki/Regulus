import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { LayoutModule } from '../../layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';
import { MakerSearchComponent } from './maker-search.component';
import { MakerService } from '../../services/api/maker.service';
import { ComplexSearchService } from '../../services/share/complex-search.service';
import { Maker } from '../../services/models/maker/maker';
import { of,Subject } from 'rxjs';
import { ConditionData } from 'src/app/services/models/search/condition-data';
import { ComplexSearchItems } from 'src/app/services/models/search/complex-search-items';
import { SaveData } from 'src/app/services/models/search/save-data';

describe('MakerSearchComponent', () => {
  let component: MakerSearchComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<MakerSearchComponent>;
  let makerServiceSpy: jasmine.SpyObj<MakerService>;
  let complexSearchServiceSpy:jasmine.SpyObj<ComplexSearchService>;


  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('MakerService', ['findByCondition', 'findComplexSearchItems']);
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService',
    ['orderComplexSearch', 'initSaveDataObj', 'initConditionDataObj', 'complexSearchOrdered$']);

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
      id: '',
      patternName: '',
      category: null,
      isDisclose: false,
      discloseGroupIDs: [],
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
      }
    });
    complexSearchServiceSpy.initConditionDataObj.and.returnValue({
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    });
    complexSearchServiceSpy.complexSearchOrdered$ = new Subject<SaveData>().asObservable();

    fixture = TestBed.createComponent(MakerSearchComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement; 
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    makerServiceSpy = TestBed.get(MakerService);
    const complexSearchItems: ComplexSearchItems = {
      displayItemList: [],
      searchConditionList:[],
      orderConditionList: [],
      isShowDisplayItem: false,
      isShowOrderCondition: false,
      isShowSaveCondition: false,
      groups: [],
    }
    const stubValue = of(complexSearchItems)

    makerServiceSpy.findComplexSearchItems.and.returnValue(stubValue)
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
    const condition: ConditionData = {
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
      searchStrings: [],
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
