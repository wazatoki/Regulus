import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {WelcomComponent} from './welcom/welcom.component';

const routes: Routes = [
  { path: '', redirectTo: '/index', pathMatch: 'full' },
  { path: 'index', component: WelcomComponent },
  { path: 'product-list', loadChildren: () => import('./product/product.module').then(m => m.ProductModule) },
  { path: 'maker-list', loadChildren: () => import('./maker/maker.module').then(m => m.MakerModule) },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
