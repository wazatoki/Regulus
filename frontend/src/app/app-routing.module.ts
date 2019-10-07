import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {WelcomComponent} from './welcom/welcom.component';
import {ProductMasterComponent} from "./master/product/product-master/product-master.component";


const routes: Routes = [
  { path: '', redirectTo: '/index', pathMatch: 'full' },
  { path: 'index', component: WelcomComponent },
  { path: 'product-master', component: ProductMasterComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
