import { TestBed } from '@angular/core/testing';

import { ComplexSearchConditionService } from './complex-search-condition.service';
import { HttpTestingController, HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpService } from '../http.service';
import { Category } from '../models/search/category';
import { ceateTestArrayForMasterMaintenanceTest } from 'src/app/services/models/search/category.spec';
import { of } from 'rxjs';
import { SaveData } from '../models/search/save-data';
import { createTestInstance1 as createTestInstanceSaveData, createTestArray as createTestArraySaveData } from 'src/app/services/models/search/save-data.spec';
import { createTestInstance1 as createTestInstanceComplexSearchItem } from 'src/app/services/models/search/complex-search-items.spec';
import { ComplexSearchItems } from '../models/search/complex-search-items';
import { createTestInstance1 as createTestInstanceConditionData } from 'src/app/services/models/search/condition-data.spec';
import { HttpErrorResponse } from '@angular/common/http';

describe('ComplexSearchConditionService', () => {

  let httpTestingController: HttpTestingController;
  let complexSearchConditionService: ComplexSearchConditionService;
  let httpServiceSpy: jasmine.SpyObj<HttpService>;

  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpService', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,
      ],
      providers: [
        { provide: HttpService, useValue: spy },
      ]
    });

    httpTestingController = TestBed.get(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });



  it('should call httpService.get with /complexSearchCondition/dataInputFormItems when called findAllCategories method', () => {
    const testData: Category[] = ceateTestArrayForMasterMaintenanceTest();

    complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData);
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: Category[] | HttpErrorResponse;

    complexSearchConditionService.findAllCategories().subscribe(data => {
      result = data;
    });

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/complexSearchCondition/dataInputFormItems');
  });

  it('should call httpService.get with /complexSearchCondition/categories when called findComplexSearchItems method', () => {
    const testData: ComplexSearchItems = createTestInstanceComplexSearchItem();

    complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData);
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: ComplexSearchItems | HttpErrorResponse;

    complexSearchConditionService.findComplexSearchItems().subscribe(data => {
      result = data;
    });

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/complexSearchCondition/complexSearchItems');
  });

  it('should call httpService.get with /complexSearchCondition and condition when called findByCondition method', () => {
    const testData: SaveData[] = createTestArraySaveData();
    complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData);
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: SaveData[] | HttpErrorResponse;

    const condition = createTestInstanceConditionData();
    complexSearchConditionService.findByCondition(condition).subscribe(data => {
      result = data;
    });

    expect(result).toEqual(testData);

    const data: Map<string, string> = new Map();
    data.set('condition', JSON.stringify(condition));
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/complexSearchCondition', data);
  });

  // it('should call httpService.post with /complexSearchCondition when called add method', () => {
  //   const testData: SaveData = createTestInstanceSaveData();
  //   testData.id = '';
  //   const resultData: SaveData = createTestInstanceSaveData();
  //   complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
  //   httpServiceSpy = TestBed.get(HttpService);
  //   httpServiceSpy.post.and.returnValue(of(resultData));

  //   let result: SaveData;

  //   complexSearchConditionService.add(testData).subscribe(data => {
  //     result = data
  //   })

  //   expect(result).toEqual(resultData);
  //   expect(httpServiceSpy.post).toHaveBeenCalledWith('/complexSearchCondition', testData);
  // });

  // it('should call httpService.put with /complexSearchCondition when called update method', () => {
  //   const testData: SaveData = createTestInstanceSaveData();
  //   const resultData: SaveData = createTestInstanceSaveData();
  //   complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
  //   httpServiceSpy = TestBed.get(HttpService);
  //   httpServiceSpy.put.and.returnValue(of(resultData));

  //   let result: SaveData;

  //   complexSearchConditionService.update(testData).subscribe(data => {
  //     result = data
  //   })

  //   expect(result).toEqual(resultData);
  //   expect(httpServiceSpy.put).toHaveBeenCalledWith('/complexSearchCondition', testData);
  // });

  it('should call httpService.delete with /complexSearchCondition when called delete method', () => {
    const testData: string[] = ['saveID1', 'saveID2'];
    const resultData: SaveData[] = [];
    complexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.delete.and.returnValue(of(resultData));

    let result: SaveData[] | HttpErrorResponse;

    complexSearchConditionService.delete(testData).subscribe(data => {
      result = data;
    });

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.delete).toHaveBeenCalledWith('/complexSearchCondition', testData);
  });

  it('should be created', () => {
    const service: ComplexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    expect(service).toBeTruthy();
  });
});
