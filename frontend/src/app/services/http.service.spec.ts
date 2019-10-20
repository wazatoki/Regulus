import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
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

  fit('should be created', () => {
    const service: HttpService = TestBed.get(HttpService);
    expect(service).toBeTruthy();
  });

  fit('should be get correct hostname', () => {
    const service: HttpService = TestBed.get(HttpService);
    expect(service.BASE_URL).toBe('http://localhost:9876/');
  });

  fit('can test HttpClient.get', () => {
    const testData: Data = { name: 'Test Data' };

    httpClient.get<Data>('/data')
      .subscribe(data =>
        expect(data).toEqual(testData)
      );

    const req = httpTestingController.expectOne('/data');
    expect(req.request.method).toEqual('GET');
    req.flush(testData);
  });
});
