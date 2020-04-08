import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ComplexSearchConditionRoutingModule } from './complex-search-condition-routing.module';
import { ComplexSearchConditionMasterComponent } from '../complex-search-condition/complex-search-condition-master/complex-search-condition-master.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LayoutModule } from '../layout/layout.module';
import { NoticeDialogComponent } from '../layout/dialog/notice-dialog/notice-dialog.component';
import { ComplexSearchDialogComponent } from '../layout/dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatPaginatorModule } from '@angular/material/paginator';
import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search/complex-search-condition-search.component';
import { ComplexSearchConditionInputFormComponent } from './complex-search-condition-input-form/complex-search-condition-input-form.component';


@NgModule({
  declarations: [
    ComplexSearchConditionMasterComponent,
    ComplexSearchConditionSearchComponent,
    ComplexSearchConditionInputFormComponent
  ],
  imports: [
    CommonModule,
    ComplexSearchConditionRoutingModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    LayoutModule,
    MatListModule,
    MatTableModule,
    MatCheckboxModule,
    MatPaginatorModule,
  ],
  entryComponents: [
    NoticeDialogComponent,
    ComplexSearchDialogComponent,
  ]
})
export class ComplexSearchConditionModule { }
