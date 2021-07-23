import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { of } from 'rxjs';
import { HttpService } from '../http.service';
import { StaffGroup } from '../models/group/staff-group';

import { StaffGroupService } from './staff-group.service';
import { createTestInstance1 as createTestInstanceStaffGroup, ceateTestArray as createTestArrayStaffGroupData } from '../models/group/staff-group.spec';
import { createTestInstance1 as createTestInstanceConditionData } from 'src/app/services/models/search/condition-data.spec';
import { HttpErrorResponse } from '@angular/common/http';

describe('StaffGroupService', () => {

  let httpTestingController: HttpTestingController;
  let httpServiceSpy: jasmine.SpyObj<HttpService>;
  let staffGroupService: StaffGroupService;
  
  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpService', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,
      ],
      providers: [
        { provide: HttpService, useValue: spy },
      ]
    })

    httpTestingController = TestBed.get(HttpTestingController);

  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should call httpService.get with /staffGroup and condition when called findByCondition method', () => {
    const testData: StaffGroup[] = createTestArrayStaffGroupData();
    staffGroupService = TestBed.get(StaffGroupService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData);
    httpServiceSpy.get.and.returnValue(stubValue);

    let result: StaffGroup[] | HttpErrorResponse;

    const condition = createTestInstanceConditionData();
    staffGroupService.findByCondition(condition).subscribe(data => {
      result = data
    })

    expect(result).toEqual(testData);

    const data: Map<string, string> = new Map();
    data.set('condition', JSON.stringify(condition));
    expect(httpServiceSpy.get).toHaveBeenCalledWith('/staffGroup', data);
  });

    it('should call httpService.post with /staffGroup when called add method', () => {
    const testData: StaffGroup = createTestInstanceStaffGroup();
    testData.id = '';
    const resultData: StaffGroup = createTestInstanceStaffGroup();
    staffGroupService = TestBed.get(StaffGroupService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.post.and.returnValue(of(resultData));

    let result: StaffGroup | HttpErrorResponse;

    staffGroupService.add(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.post).toHaveBeenCalledWith('/staffGroup', testData);
  });

  it('should call httpService.put with /staffGroup when called update method', () => {
    const testData: StaffGroup = createTestInstanceStaffGroup();
    const resultData: StaffGroup = createTestInstanceStaffGroup();
    staffGroupService = TestBed.get(StaffGroupService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.put.and.returnValue(of(resultData));

    let result: StaffGroup | HttpErrorResponse;

    staffGroupService.update(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.put).toHaveBeenCalledWith('/staffGroup', testData);
  });


  it('should call httpService.delete with /staffGroup when called delete method', () => {
    const testData: string[] = ['ID1', 'ID2'];
    const resultData: StaffGroup[] = [];
    staffGroupService = TestBed.get(StaffGroupService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.delete.and.returnValue(of(resultData));

    let result: StaffGroup[] | HttpErrorResponse;

    staffGroupService.delete(testData).subscribe(data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.delete).toHaveBeenCalledWith('/staffGroup', testData);
  });


  it('should be created', () => {
    const service: StaffGroupService = TestBed.get(StaffGroupService);
    expect(service).toBeTruthy();
  });
});
