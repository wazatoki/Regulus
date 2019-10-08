import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import {MatTableModule} from '@angular/material/table';
import { ProductMasterComponent } from './product-master.component';
import { Component } from '@angular/core';

@Component({selector: 'app-layout-header-sidebar-contents', template: ''})
class LayoutHeaderSidebarContentsComponent {}

@Component({selector: 'app-product-search', template: ''})
class ProductSearchComponent {}

describe('ProductMasterComponent', () => {
  let component: ProductMasterComponent;
  let fixture: ComponentFixture<ProductMasterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ 
        ProductMasterComponent,
        ProductSearchComponent,
        LayoutHeaderSidebarContentsComponent,
      ],
      imports: [
        MatTableModule,
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductMasterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
