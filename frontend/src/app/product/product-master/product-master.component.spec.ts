import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import {MatTableModule} from '@angular/material/table';
import { ProductMasterComponent } from './product-master.component';
import { LayoutModule } from '../../layout/layout.module';
import { RouterTestingModule } from '@angular/router/testing';
import { Component } from '@angular/core';
import { LoginService } from 'src/app/services/api/login.service';
import { BehaviorSubject } from 'rxjs';

@Component({selector: 'app-product-search', template: ''})
class ProductSearchComponent {}

@Component({selector: 'app-header', template: ''})
class HeaderComponent {}

describe('ProductMasterComponent', () => {
  let component: ProductMasterComponent;
  let productMasterElement: HTMLElement;
  let fixture: ComponentFixture<ProductMasterComponent>;
  let spy: jasmine.SpyObj<LoginService>;

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('LoginService', ['currentUser', 'currentUserToken', 'currentUserValue', 'currentUserTokenValue']);

    TestBed.configureTestingModule({
      declarations: [
        ProductMasterComponent,
        ProductSearchComponent,
        HeaderComponent,
      ],
      imports: [
        MatTableModule,
        LayoutModule,
        RouterTestingModule,
      ],
      providers: [
        { provide: LoginService, useValue: spy },
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    spy = TestBed.get(LoginService);
    spy.currentUserToken = new BehaviorSubject<string>('').asObservable();

    fixture = TestBed.createComponent(ProductMasterComponent);
    component = fixture.componentInstance;
    productMasterElement = fixture.debugElement.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain table header', () => {
    // expect(productMasterElement.textContent).toContain('製造販売業者');
    // expect(productMasterElement.textContent).toContain('一般的名称');
    expect(productMasterElement.textContent).toContain('ＪＡＮコード');
    expect(productMasterElement.textContent).toContain('商品名');
    expect(productMasterElement.textContent).toContain('製品番号');
  });

  it('should explain table body' , () => {
    component.dataSource = [
      {
        name: 'name_aaa',
        jancode: 'jancode_aaa',
        code: 'code_aaa',
      },
      {
        name: 'name_bbb',
        jancode: 'jancode_bbb',
        code: 'code_bbb',
      }
    ];

    fixture.detectChanges();

    expect(productMasterElement.textContent).toContain('name_aaa');
    expect(productMasterElement.textContent).toContain('jancode_aaa');
    expect(productMasterElement.textContent).toContain('code_aaa');
    expect(productMasterElement.textContent).toContain('name_bbb');
    expect(productMasterElement.textContent).toContain('jancode_bbb');
    expect(productMasterElement.textContent).toContain('code_bbb');
  });
});
