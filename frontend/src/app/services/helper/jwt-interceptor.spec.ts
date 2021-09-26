import { TestBed, inject } from '@angular/core/testing';
import { LoginService } from '../api/login.service';

import { JwtInterceptor } from './jwt-interceptor';

describe('JwtInterceptor', () => {

  beforeEach(() => {

    const spy = jasmine.createSpyObj('LoginService', ['login', 'logout']);

    TestBed.configureTestingModule({
      providers: [
        JwtInterceptor,
        { provide: LoginService, useValue: spy },
      ]
    });

  });


  it('should create an instance', () => {
    const service: JwtInterceptor = TestBed.get(JwtInterceptor);
    expect(service).toBeTruthy();
  });
});
