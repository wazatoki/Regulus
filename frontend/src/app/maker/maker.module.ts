import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { LayoutModule } from '../layout/layout.module';
import { MakerRoutingModule } from './maker-routing.module';
import { MakerMasterComponent } from './maker-master/maker-master.component';
import { MakerSearchComponent } from './maker-search/maker-search.component';


@NgModule({
  declarations: [
    MakerMasterComponent,
    MakerSearchComponent,
  ],
  imports: [
    CommonModule,
    MakerRoutingModule,
    LayoutModule,
    MatTableModule,
  ]
})
export class MakerModule { }
