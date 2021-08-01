import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router'
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout';
import { HeaderComponent } from './header/header.component';
import { LayoutHeaderSidebarContentsComponent } from './layout-header-sidebar-contents/layout-header-sidebar-contents.component';
import { LayoutHeaderContentsComponent } from './layout-header-contents/layout-header-contents.component';
import { LayoutContentsComponent } from './layout-contents/layout-contents.component';
import { SearchComponent } from './search/search.component';
import { MatFormFieldModule, MatInputModule, MatSelectModule, MatRadioModule,MatCheckboxModule, MatButtonModule,MatDialogModule, MatListModule, MatGridListModule, MatCardModule } from '@angular/material';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CancelComponent } from './form/buttons/cancel/cancel.component';
import { AutofocusDirective } from '../directive/autofocus.directive';
import { SubmitComponent } from './form/buttons/submit/submit.component';
import { ClearComponent } from './form/buttons/clear/clear.component';
import { AlertDialogComponent } from './dialog/alert-dialog/alert-dialog.component';
import { NoticeDialogComponent } from './dialog/notice-dialog/notice-dialog.component';
import { CreateComponent } from './form/buttons/create/create.component';
import { DeleteComponent } from './form/buttons/delete/delete.component';
import { UpdateComponent } from './form/buttons/update/update.component';
import { ComplexSearchConditionItemComponent } from './complex-search/complex-search-condition-item/complex-search-condition-item.component';
import { ComplexSearchOrderItemComponent } from './complex-search/complex-search-order-item/complex-search-order-item.component';
import { ComplexSearchComponent } from './complex-search/complex-search.component';
import { ComplexSearchDialogComponent } from './dialog/complex-search-dialog/complex-search-dialog/complex-search-dialog.component';
import { FavoriteConditionsComponent } from './complex-search/favorite-conditions/favorite-conditions.component';

@NgModule({
  declarations: [
    HeaderComponent,
    LayoutContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutHeaderSidebarContentsComponent,
    SearchComponent,
    CancelComponent,
    SubmitComponent,
    ClearComponent,
    AlertDialogComponent,
    NoticeDialogComponent,
    CreateComponent,
    DeleteComponent,
    UpdateComponent,
    ComplexSearchConditionItemComponent,
    ComplexSearchOrderItemComponent,
    ComplexSearchComponent,
    ComplexSearchDialogComponent,
    AutofocusDirective,
    FavoriteConditionsComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatRadioModule,
    MatCheckboxModule,
    MatButtonModule,
    MatDialogModule,
    MatListModule,
    MatGridListModule,
    MatCardModule,
    DragDropModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  exports: [
    LayoutContentsComponent,
    LayoutHeaderContentsComponent,
    LayoutHeaderSidebarContentsComponent,
    SearchComponent,
    CancelComponent,
    SubmitComponent,
    ClearComponent,
    AlertDialogComponent,
    NoticeDialogComponent,
    CreateComponent,
    DeleteComponent,
    UpdateComponent,
    ComplexSearchConditionItemComponent,
    ComplexSearchOrderItemComponent,
    AutofocusDirective,
  ],
  entryComponents: [
    AlertDialogComponent,
  ]
})
export class LayoutModule { }
