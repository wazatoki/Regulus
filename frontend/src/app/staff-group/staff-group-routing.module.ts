import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StaffGroupMasterComponent } from './staff-group-master/staff-group-master.component';

const routes: Routes = [{ path: '', component: StaffGroupMasterComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StaffGroupRoutingModule { }
