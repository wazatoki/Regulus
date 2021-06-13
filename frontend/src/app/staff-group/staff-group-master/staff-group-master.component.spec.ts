import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { StaffGroupMasterComponent } from './staff-group-master.component';

describe('StaffGroupMasterComponent', () => {
  let component: StaffGroupMasterComponent;
  let fixture: ComponentFixture<StaffGroupMasterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StaffGroupMasterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StaffGroupMasterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
