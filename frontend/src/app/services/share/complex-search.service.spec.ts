import { TestBed } from '@angular/core/testing';

import { ComplexSearchService } from './complex-search.service';
import { HttpClient } from '@angular/common/http';

describe('ComplexSearchService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    providers: [
      { provide: HttpClient, useValue: {}}
    ]
  }));

  it('should be created', () => {
    const service: ComplexSearchService = TestBed.get(ComplexSearchService);
    expect(service).toBeTruthy();
  });
});
