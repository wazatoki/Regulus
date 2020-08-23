import { TestBed } from '@angular/core/testing';
import { HttpClient, HttpParams, HttpErrorResponse, HttpHeaders } from '@angular/common/http';

import { LoginService } from './login.service';

describe('LoginService', () => {
  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpClient', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule({
      providers: [
        { provide: HttpClient, useValue: spy },
      ]
    })
    
  });

  it('should be created', () => {
    const service: LoginService = TestBed.get(LoginService);
    expect(service).toBeTruthy();
  });
});
