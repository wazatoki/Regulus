import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from 'src/app/layout/layout.module';
import { StaffGroupService } from 'src/app/services/api/staff-group.service';

import { StaffGroupSearchComponent } from './staff-group-search.component';

describe('StaffGroupSearchComponent', () => {
  let component: StaffGroupSearchComponent;
  let fixture: ComponentFixture<StaffGroupSearchComponent>;
  let staffGroupServiceSpy: jasmine.SpyObj<StaffGroupService>

  beforeEach(async(() => {

    const staffGroupServiceSpy = jasmine.createSpyObj('StaffGroupService', ['findByCondition']);

    TestBed.configureTestingModule({
      declarations: [ StaffGroupSearchComponent ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: StaffGroupService, useValue: staffGroupServiceSpy },
      ],
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
