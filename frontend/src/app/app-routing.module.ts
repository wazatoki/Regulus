import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { WelcomComponent } from './welcom/welcom.component';
import { TestComponent } from './test/test.component';
import { AuthGuardGuard } from './services/helper/auth-guard.guard';

const routes: Routes = [
  { path: '', redirectTo: '/index', pathMatch: 'full' },
  { path: 'index', component: WelcomComponent },
  { path: 'test', component: TestComponent },
  { path: 'login', loadChildren: () => import('./login/login.module').then(m => m.LoginModule) },
  { path: 'complex-search-condition-list', loadChildren: () => import('./complex-search-condition/complex-search-condition.module').then(m => m.ComplexSearchConditionModule), canActivate: [AuthGuardGuard] },
  { path: 'staff-group-list', loadChildren: () => import('./staff-group/staff-group.module').then(m => m.StaffGroupModule) },
  { path: 'product-list', loadChildren: () => import('./product/product.module').then(m => m.ProductModule) },
  { path: 'maker-list', loadChildren: () => import('./maker/maker.module').then(m => m.MakerModule) },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
