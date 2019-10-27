import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
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
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    LayoutModule,
    MatTableModule,
  ]
})
export class MakerModule { }
