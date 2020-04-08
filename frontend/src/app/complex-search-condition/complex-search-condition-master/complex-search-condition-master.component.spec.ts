import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionMasterComponent } from './complex-search-condition-master.component';

describe('ComplexSearchConditionMasterComponent', () => {
  let component: ComplexSearchConditionMasterComponent;
  let fixture: ComponentFixture<ComplexSearchConditionMasterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ComplexSearchConditionMasterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchConditionMasterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
