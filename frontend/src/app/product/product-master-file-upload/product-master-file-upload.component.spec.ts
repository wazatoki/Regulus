import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductMasterFileUploadComponent } from './product-master-file-upload.component';

describe('ProductMasterFileUploadComponent', () => {
  let component: ProductMasterFileUploadComponent;
  let fixture: ComponentFixture<ProductMasterFileUploadComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductMasterFileUploadComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductMasterFileUploadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
