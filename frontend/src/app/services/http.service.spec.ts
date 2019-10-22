import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpClient, HttpParams, HttpErrorResponse } from '@angular/common/http';
import { HttpService } from './http.service';

describe('HttpService', () => {

  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;

  class Data {
    name: string
  }

  beforeEach(
    () => {
      TestBed.configureTestingModule(
        {
          imports: [
            HttpClientTestingModule,
          ],
        }
      )

      httpClient = TestBed.get(HttpClient);
      httpTestingController = TestBed.get(HttpTestingController);
    }
  );

  afterEach(() => {
    // After every test, assert that there are no more pending requests.
    httpTestingController.verify();
  });

  it('should be created', () => {
    const service: HttpService = TestBed.get(HttpService);
    expect(service).toBeTruthy();
  });

  it('should be correct host url', () => {
    const service: HttpService = TestBed.get(HttpService);
    expect(service.HOST_URL).toEqual('localhost:9876');
  });

  it('can test HttpClient.get', () => {
    const testData: Data = { name: 'Test Data' };

    httpClient.get<Data>('/data')
      .subscribe(data => {
        expect(data).toEqual(testData)
      });

    const req = httpTestingController.expectOne('/data');
    expect(req.request.method).toEqual('GET');
    req.flush(testData);
  });

  it('get mothod without params', () => {
    const testData: Data = { name: 'Test Data' };
    const service: HttpService = TestBed.get(HttpService);

    service.get<Data>('/data')
      .subscribe(data => {
        expect(data).toEqual(testData)
      });

    const req = httpTestingController.expectOne('localhost:9876/data');
    expect(req.request.method).toEqual('GET');
    req.flush(testData);
  });

  it('get mothod with params', () => {
    const testData: Data = { name: 'Test Data' };
    const service: HttpService = TestBed.get(HttpService);

    service.get<Data>('/data', { id: 'idstring' })
      .subscribe(data => {
        expect(data).toEqual(testData)
      });

    const req = httpTestingController.expectOne('localhost:9876/data');
    expect(req.request.method).toEqual('GET');
    const params: HttpParams = new HttpParams();
    params.set('id', 'idstring')
    expect(req.request.params.toString).toEqual(params.toString)
    req.flush(testData);
  });

});
