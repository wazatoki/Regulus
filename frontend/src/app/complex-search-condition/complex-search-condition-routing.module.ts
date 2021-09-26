import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {
  ComplexSearchConditionMasterComponent
} from 'src/app/complex-search-condition/complex-search-condition-master/complex-search-condition-master.component';


const routes: Routes = [{ path: '', component: ComplexSearchConditionMasterComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ComplexSearchConditionRoutingModule { }
