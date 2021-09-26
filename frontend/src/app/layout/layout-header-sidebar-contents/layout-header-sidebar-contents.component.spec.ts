import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LayoutHeaderSidebarContentsComponent } from './layout-header-sidebar-contents.component';
import { Component } from '@angular/core';

@Component({ selector: 'app-header', template: '' })
class HeaderComponent { }

describe('LayoutHeaderSidebarContentsComponent', () => {
  let component: LayoutHeaderSidebarContentsComponent;
  let fixture: ComponentFixture<LayoutHeaderSidebarContentsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        LayoutHeaderSidebarContentsComponent,
        HeaderComponent,
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LayoutHeaderSidebarContentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
