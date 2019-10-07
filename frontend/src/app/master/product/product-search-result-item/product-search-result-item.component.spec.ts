import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductSearchResultItemComponent } from './product-search-result-item.component';

describe('ProductSearchResultItemComponent', () => {
  let component: ProductSearchResultItemComponent;
  let fixture: ComponentFixture<ProductSearchResultItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductSearchResultItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductSearchResultItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
