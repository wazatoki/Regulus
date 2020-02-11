import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchOrderItemComponent } from './complex-search-order-item.component';

describe('ComplexSearchOrderItemComponent', () => {
  let component: ComplexSearchOrderItemComponent;
  let fixture: ComponentFixture<ComplexSearchOrderItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ComplexSearchOrderItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchOrderItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
