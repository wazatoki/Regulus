import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { StaffGroupRoutingModule } from './staff-group-routing.module';
import { StaffGroupMasterComponent } from './staff-group-master/staff-group-master.component';
import { LayoutModule } from '../layout/layout.module';
import { StaffGroupSearchComponent } from './staff-group-search/staff-group-search.component';
import { StaffGroupInputFormDialogComponent } from './staff-group-input-form-dialog/staff-group-input-form-dialog.component';
import { NoticeDialogComponent } from '../layout/dialog/notice-dialog/notice-dialog.component';


@NgModule({
  declarations: [
    StaffGroupMasterComponent,
    StaffGroupSearchComponent,
    StaffGroupInputFormDialogComponent],
  imports: [
    CommonModule,
    StaffGroupRoutingModule,
    LayoutModule,
  ],
  entryComponents: [
    NoticeDialogComponent,
  ]
})
export class StaffGroupModule { }
