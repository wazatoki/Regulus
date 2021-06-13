import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { StaffGroupSearchComponent } from './staff-group-search.component';

describe('StaffGroupSearchComponent', () => {
  let component: StaffGroupSearchComponent;
  let fixture: ComponentFixture<StaffGroupSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StaffGroupSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StaffGroupSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
