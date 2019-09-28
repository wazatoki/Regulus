import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {WelcomComponent} from './welcom/welcom.component';


const routes: Routes = [
  { path: '', redirectTo: '/index', pathMatch: 'full' },
  { path: 'index', component: WelcomComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
