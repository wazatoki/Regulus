import { TestBed, inject } from '@angular/core/testing';
import { HttpErrorInterceptor } from './http-error-interceptor';
import { LoginService } from '../api/login.service'

describe('HttpErrorInterceptor', () => {

  beforeEach( () => {

    const spy = jasmine.createSpyObj('LoginService', ['login', 'logout']);

    TestBed.configureTestingModule({
      providers: [
        HttpErrorInterceptor,
        { provide: LoginService, useValue: spy },
      ]
    });

  });

  it('should create an instance', () => {
    const service: HttpErrorInterceptor = TestBed.get(HttpErrorInterceptor);
    expect(service).toBeTruthy();
  });
});
