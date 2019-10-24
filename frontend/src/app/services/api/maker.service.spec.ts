import { TestBed } from '@angular/core/testing';

import { MakerService } from './maker.service';
import { HttpService } from '../http.service'

import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { Maker} from '../models/maker';
import { of } from 'rxjs';

describe('MakerService', () => {

  let httpTestingController: HttpTestingController;
  let makerService: MakerService;
  let httpServiceSpy: jasmine.SpyObj<HttpService>;

  beforeEach(() => {
  
    const spy = jasmine.createSpyObj('HttpService', ['get', 'post']);

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

  it('findById mothod', () => {
    const testData: Maker = { id: 'testid', name: 'Test Maker' };
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData)
    httpServiceSpy.get.and.returnValue(stubValue);
  
    let result: Maker;

    makerService.findById('testid').subscribe( data => {
      result = data
    })

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('maker', {id: 'testid'});
  });

  it('findAll mothod', () => {
    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    const stubValue = of(testData)
    httpServiceSpy.get.and.returnValue(stubValue);
  
    let result: Maker[];

    makerService.findAll().subscribe( data => {
      result = data
    })

    expect(result).toEqual(testData);
    expect(httpServiceSpy.get).toHaveBeenCalledWith('maker');
  });

  it('add method', () => {
    const testData: Maker = { id: '', name: 'Test Maker1' };
    const resultData: Maker = { id: 'testid1', name: 'Test Maker1' };
    makerService = TestBed.get(MakerService);
    httpServiceSpy = TestBed.get(HttpService);
    httpServiceSpy.post.and.returnValue(of(resultData));
  
    let result: Maker;

    makerService.add(testData).subscribe( data => {
      result = data
    })

    expect(result).toEqual(resultData);
    expect(httpServiceSpy.post).toHaveBeenCalledWith('maker', testData);
  });

});
