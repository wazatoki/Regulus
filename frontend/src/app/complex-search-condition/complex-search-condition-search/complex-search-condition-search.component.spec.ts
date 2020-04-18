import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search.component';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { LayoutModule } from 'src/app/layout/layout.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { ComplexSearchService } from 'src/app/services/share/complex-search.service';
import { SaveData } from 'src/app/services/models/search/save-data';
import { of } from 'rxjs';

import { createTestArray } from 'src/app/services/models/search/save-data.spec';

describe('ComplexSearchConditionSearchComponent', () => {
  let component: ComplexSearchConditionSearchComponent;
  let fixture: ComponentFixture<ComplexSearchConditionSearchComponent>;
  let complexSearchConditionServiceSpy: jasmine.SpyObj<ComplexSearchConditionService>;
  let complexSearchServiceSpy: jasmine.SpyObj<ComplexSearchService>

  beforeEach(async(() => {

    const complexSearchConditionServiceSpy = jasmine.createSpyObj('ComplexSearchConditionService', ['findByCondition']);
    const complexSearchServiceSpy = jasmine.createSpyObj('ComplexSearchService', ['initConditionDataObj']);

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

    fixture = TestBed.createComponent(ComplexSearchConditionSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('called api when onSearch execute', () => {
    complexSearchConditionServiceSpy = TestBed.get(ComplexSearchConditionService);
    const searchWords = 'aaa bbb ccc';
    const condition = {
      searchStrings: ['aaa','bbb','ccc'],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
    const data: SaveData[] = createTestArray();
    complexSearchConditionServiceSpy.findByCondition.and.returnValue(of(data))
    component.onSearch(searchWords);

    component.fetched.subscribe( (res: SaveData[]) => {
      expect(res).toBe(data);
    });

    expect(complexSearchConditionServiceSpy.findByCondition).toHaveBeenCalledWith(condition);
  });
});
