import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FlexLayoutModule } from '@angular/flex-layout'
import {MatMenuModule} from '@angular/material/menu';
import {MatButtonModule} from '@angular/material/button'; 

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { WelcomComponent } from './welcom/welcom.component';
import { LayoutHeaderSidebarContentsComponent } from './layout/layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { LayoutHeaderContentsComponent } from './layout/layout-header-contents/layout-header-contents.component';
import { LayoutContentsComponent } from './layout/layout-contents/layout-contents.component';
import { ProductMasterComponent } from './master/product/product-master/product-master.component';
import { ProductSearchComponent } from './master/product/product-search/product-search.component';
import { ProductSearchResultComponent } from './master/product/product-search-result/product-search-result.component';
import { ProductSearchResultItemComponent } from './master/product/product-search-result-item/product-search-result-item.component';
import { ProductInputFormComponent } from './master/product/product-input-form/product-input-form.component';
import { ProductMasterFileUploadComponent } from './master/product/product-master-file-upload/product-master-file-upload.component';
import { ProductMasterDetailComponent } from './master/product/product-master-detail/product-master-detail.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    WelcomComponent,
    LayoutHeaderSidebarContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutContentsComponent,
    ProductMasterComponent,
    ProductSearchComponent,
    ProductSearchResultComponent,
    ProductSearchResultItemComponent,
    ProductInputFormComponent,
    ProductMasterFileUploadComponent,
    ProductMasterDetailComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    MatButtonModule,
    MatMenuModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
