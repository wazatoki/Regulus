import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ComplexSearchConditionRoutingModule } from './complex-search-condition-routing.module';
import { ComplexSearchConditionMasterComponent } from '../complex-search-condition/complex-search-condition-master/complex-search-condition-master.component';


@NgModule({
  declarations: [ComplexSearchConditionMasterComponent],
  imports: [
    CommonModule,
    ComplexSearchConditionRoutingModule
  ]
})
export class ComplexSearchConditionModule { }
