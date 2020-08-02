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
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDialogModule } from '@angular/material/dialog';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search/complex-search-condition-search.component';
import { ComplexSearchConditionInputFormDialogComponent } from './complex-search-condition-input-form-dialog/complex-search-condition-input-form-dialog.component';


@NgModule({
  declarations: [
    ComplexSearchConditionMasterComponent,
    ComplexSearchConditionSearchComponent,
    ComplexSearchConditionInputFormDialogComponent
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
    MatCardModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatRadioModule,
    MatGridListModule,
    MatDialogModule,
    DragDropModule
  ],
  entryComponents: [
    NoticeDialogComponent,
    ComplexSearchDialogComponent,
    ComplexSearchConditionInputFormDialogComponent,
  ]
})
export class ComplexSearchConditionModule { }
