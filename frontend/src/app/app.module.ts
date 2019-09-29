import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FlexLayoutModule } from '@angular/flex-layout'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { WelcomComponent } from './welcom/welcom.component';
import { LayoutHeaderSidebarContentsComponent } from './layout/layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { LayoutHeaderContentsComponent } from './layout/layout-header-contents/layout-header-contents.component';
import { LayoutContentsComponent } from './layout/layout-contents/layout-contents.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    WelcomComponent,
    LayoutHeaderSidebarContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutContentsComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
