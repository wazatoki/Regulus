import { TestBed } from '@angular/core/testing';

import { ComplexSearchService } from './complex-search.service';

describe('ComplexSearchService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ComplexSearchService = TestBed.get(ComplexSearchService);
    expect(service).toBeTruthy();
  });
});
