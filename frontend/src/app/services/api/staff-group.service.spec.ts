import { TestBed } from '@angular/core/testing';

import { StaffGroupService } from './staff-group.service';

describe('StaffGroupService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: StaffGroupService = TestBed.get(StaffGroupService);
    expect(service).toBeTruthy();
  });
});
