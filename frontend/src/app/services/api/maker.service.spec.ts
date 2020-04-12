import { TestBed } from '@angular/core/testing';

import { MakerService } from './maker.service';
import { HttpService } from '../http.service';

import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { Maker } from '../models/maker/maker';
import { ComplexSearchItems } from '../models/search/complex-search-items';
import { ConditionData } from '../models/search/condition-data';
import { of } from 'rxjs';

import { createTestInstance1 as createConditionData1 } from 'src/app/services/models/search/condition-data.spec';
import { createTestInstance1 } from 'src/app/services/models/search/complex-search-items.spec';

describe('MakerService', () => {

  let httpTestingController: HttpTestingController;
  let makerService: MakerService;
  let httpServiceSpy: jasmine.SpyObj<HttpService>;

  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpService', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule(
      {
        imports: [
          HttpClientTestingModule,
        ],
        providers: [
          { provide: HttpService, useValue: spy },
        ]
      }
    )

    httpTestingController = TestBed.get(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });


  it('should be created', () => {
    const service: MakerService = TestBed.get(MakerService);
    expect(service).toBeTruthy();
  });

  it('findById method', () => {
    const testData: Maker = { id: 'testid', name: 'Test Maker' };
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData);
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: Maker;

    makerService.findById('testid').subscribe(data => {
      result = data
    })

    expect(result).toEqual(testData);

    const id: Map<string, string> = new Map();
    id.set('id', 'testid');
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/maker', id);
  });

  it('findByCondition method', () => {
    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    makerService = TestBed.get(MakerService);
    const conditionData: ConditionData = createConditionData1();
    
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData)
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: Maker[];

    makerService.findByCondition(conditionData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(testData);

    const data: Map<string, string> = new Map();
    data.set('condition',JSON.stringify(conditionData));
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/maker', data);
  });

  it('findAll method', () => {
    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData)
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: Maker[];

    makerService.findAll().subscribe(data => {
      result = data
    })

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/maker');
  });

  it('add method', () => {
    const testData: Maker = { id: '', name: 'Test Maker1' };
    const resultData: Maker = { id: 'testid1', name: 'Test Maker1' };
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.post.and.returnValue(of(resultData));

    let result: Maker;

    makerService.add(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.post).toHaveBeenCalledWith('/maker', testData);
  });

  it('update method', () => {
    const testData: Maker = { id: 'testid1', name: 'Test Maker1' };
    const resultData: Maker = { id: 'testid1', name: 'Test Maker1' };
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.put.and.returnValue(of(resultData));

    let result: Maker;

    makerService.update(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.put).toHaveBeenCalledWith('/maker', testData);
  });

  it('delete method', () => {
    const testData: string[] = ['id1', 'id2'];
    const resultData: Maker[] = [];
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.delete.and.returnValue(of(resultData));

    let result: Maker[];

    makerService.delete(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.delete).toHaveBeenCalledWith('/maker/delete', testData);
  });

  it('findComplexSearchItems method', () => {
    const testData: ComplexSearchItems = createTestInstance1();
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData)
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: ComplexSearchItems;

    makerService.findComplexSearchItems().subscribe(data => {
      result = data
    })

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/maker/ComplexSearchItems');
  });

});
