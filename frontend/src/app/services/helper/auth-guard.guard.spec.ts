import { TestBed, async, inject } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { HttpClient, HttpParams, HttpErrorResponse, HttpHeaders } from '@angular/common/http';

import { AuthGuardGuard } from './auth-guard.guard';

describe('AuthGuardGuard', () => {
  beforeEach(() => {

    const spy = jasmine.createSpyObj('HttpClient', ['get', 'post', 'put', 'delete']);

    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      providers: [
        AuthGuardGuard,
        { provide: HttpClient, useValue: spy },
      ]
    });
  });

  it('should ...', inject([AuthGuardGuard], (guard: AuthGuardGuard) => {
    expect(guard).toBeTruthy();
  }));
});
