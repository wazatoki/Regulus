import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { StaffGroupInputFormDialogComponent } from './staff-group-input-form-dialog.component';

describe('StaffGroupInputFormDialogComponent', () => {
  let component: StaffGroupInputFormDialogComponent;
  let fixture: ComponentFixture<StaffGroupInputFormDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StaffGroupInputFormDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StaffGroupInputFormDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
