import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDialogModule } from '@angular/material/dialog';
import { MatListModule } from '@angular/material/list';
import { LayoutModule } from '../layout/layout.module';
import { MakerRoutingModule } from './maker-routing.module';
import { MakerMasterComponent } from './maker-master/maker-master.component';
import { MakerSearchComponent } from './maker-search/maker-search.component';
import { MakerInputFormComponent } from './maker-input-form/maker-input-form.component';
import { NoticeDialogComponent } from '../layout/dialog/notice-dialog/notice-dialog.component';


@NgModule({
  declarations: [
    MakerMasterComponent,
    MakerSearchComponent,
    MakerInputFormComponent,
  ],
  imports: [
    CommonModule,
    MakerRoutingModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatGridListModule,
    MatDialogModule,
    MatListModule,
    MatTableModule,
    MatCheckboxModule,
    MatPaginatorModule,
    LayoutModule,
  ],
  entryComponents: [
    MakerInputFormComponent,
    NoticeDialogComponent,
  ]
})
export class MakerModule { }
