import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { ProductRoutingModule } from './product-routing.module';
import { ProductMasterComponent } from './product-master/product-master.component';
import { LayoutModule } from '../layout/layout.module';
import { ProductInputFormComponent } from './product-input-form/product-input-form.component';
import { ProductMasterDetailComponent } from './product-master-detail/product-master-detail.component';
import { ProductMasterFileUploadComponent } from './product-master-file-upload/product-master-file-upload.component';
import { ProductSearchComponent } from './product-search/product-search.component';
import { ProductSearchResultItemComponent } from './product-search-result-item/product-search-result-item.component';


@NgModule({
  declarations: [
    ProductMasterComponent,
    ProductInputFormComponent,
    ProductMasterDetailComponent,
    ProductMasterFileUploadComponent,
    ProductSearchComponent,
    ProductSearchResultItemComponent,
  ],
  imports: [
    CommonModule,
    LayoutModule,
    ProductRoutingModule,
    MatTableModule,
  ],
  exports: [
    ProductMasterComponent
  ]
})
export class ProductModule { }
