import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout';
import { HeaderComponent } from './header/header.component';
import { LayoutHeaderSidebarContentsComponent } from './layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { LayoutHeaderContentsComponent } from './layout-header-contents/layout-header-contents.component';
import { LayoutContentsComponent } from './layout-contents/layout-contents.component';
import { SearchComponent } from './search/search.component';
import { MatFormFieldModule } from '@angular/material/form-field'; 
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { FormsModule } from '@angular/forms';
import { CancelComponent } from './form/buttons/cancel/cancel.component';
import { AutofocusDirective } from '../directive/autofocus.directive';
import { SubmitComponent } from './form/buttons/submit/submit.component';
import { ClearComponent } from './form/buttons/clear/clear.component';
import { AlertDiarogComponent } from './dialog/alert-diarog/alert-diarog.component';
import { NoticeDialogComponent } from './dialog/notice-dialog/notice-dialog.component';
import { CreateComponent } from './form/buttons/create/create.component';
import { DeleteComponent } from './form/buttons/delete/delete.component';
import { UpdateComponent } from './form/buttons/update/update.component';

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
    AlertDiarogComponent,
    NoticeDialogComponent,
    CreateComponent,
    DeleteComponent,
    UpdateComponent,
  ],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatDialogModule,
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
    AlertDiarogComponent,
    NoticeDialogComponent,
    CreateComponent,
    DeleteComponent,
    UpdateComponent,
  ]
})
export class LayoutModule { }
