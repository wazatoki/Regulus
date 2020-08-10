import { TestBed } from '@angular/core/testing';

import { HttpService } from '../http.service';
import { LoginService } from './login.service';

describe('LoginService', () => {
  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpService', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule({
      providers: [
        { provide: HttpService, useValue: spy },
      ]
    })
    
  });

  it('should be created', () => {
    const service: LoginService = TestBed.get(LoginService);
    expect(service).toBeTruthy();
  });
});
