import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import {MatMenuModule} from '@angular/material/menu';
import {MatButtonModule} from '@angular/material/button';
import { WelcomComponent } from './welcom.component';
import { Component } from '@angular/core';

@Component({selector: 'app-layout-header-sidebar-contents', template: ''})
class LayoutHeaderSidebarContentsComponent {}

describe('WelcomComponent', () => {
  let component: WelcomComponent;
  let fixture: ComponentFixture<WelcomComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        WelcomComponent,
        LayoutHeaderSidebarContentsComponent,
      ],
      imports: [
        RouterTestingModule,
        MatMenuModule,
        MatButtonModule,
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WelcomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
