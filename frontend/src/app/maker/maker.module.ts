import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatTableModule } from '@angular/material/table';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatBottomSheetModule } from '@angular/material/bottom-sheet';
import { MatListModule } from '@angular/material/list';
import { LayoutModule } from '../layout/layout.module';
import { MakerRoutingModule } from './maker-routing.module';
import { MakerMasterComponent } from './maker-master/maker-master.component';
import { MakerSearchComponent } from './maker-search/maker-search.component';
import { MakerInputFormComponent } from './maker-input-form/maker-input-form.component';


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
    MatBottomSheetModule,
    MatListModule,
    MatTableModule,
    LayoutModule,
  ],
  entryComponents: [
    MakerInputFormComponent,
  ]
})
export class MakerModule { }
