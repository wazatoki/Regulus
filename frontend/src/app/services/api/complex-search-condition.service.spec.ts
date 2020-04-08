import { TestBed } from '@angular/core/testing';

import { ComplexSearchConditionService } from './complex-search-condition.service';
import { HttpTestingController, HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpService } from '../http.service';

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
      })

      httpTestingController = TestBed.get(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  

  it('should be created', () => {
    const service: ComplexSearchConditionService = TestBed.get(ComplexSearchConditionService);
    expect(service).toBeTruthy();
  });
});
