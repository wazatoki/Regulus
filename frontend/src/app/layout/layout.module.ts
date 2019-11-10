import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout'
import { HeaderComponent } from './header/header.component';
import { LayoutHeaderSidebarContentsComponent } from './layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { LayoutHeaderContentsComponent } from './layout-header-contents/layout-header-contents.component';
import { LayoutContentsComponent } from './layout-contents/layout-contents.component';
import { SearchComponent } from './search/search.component';
import { MatFormFieldModule } from '@angular/material/form-field'; 
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { CancelComponent } from './form/buttons/cancel/cancel.component';
import { AutofocusDirective } from '../directive/autofocus.directive';
import { SubmitComponent } from './form/buttons/submit/submit.component';
import { ClearComponent } from './form/buttons/clear/clear.component';

@NgModule({
  declarations: [
    HeaderComponent,
    LayoutContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutHeaderSidebarContentsComponent,
    AutofocusDirective,
    SearchComponent,
    CancelComponent,
    SubmitComponent,
    ClearComponent,
  ],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    FormsModule,
  ],
  exports: [
    LayoutContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutHeaderSidebarContentsComponent,
    SearchComponent,
    CancelComponent,
    SubmitComponent,
    ClearComponent,
  ]
})
export class LayoutModule { }
