import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakerComplexSearchComponent } from './maker-complex-search.component';

describe('MakerComplexSearchComponent', () => {
  let component: MakerComplexSearchComponent;
  let fixture: ComponentFixture<MakerComplexSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakerComplexSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerComplexSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
