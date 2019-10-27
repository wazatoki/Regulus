import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakerInputFormComponent } from './maker-input-form.component';

describe('MakerInputFormComponent', () => {
  let component: MakerInputFormComponent;
  let fixture: ComponentFixture<MakerInputFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakerInputFormComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerInputFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
