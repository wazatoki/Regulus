import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LayoutHeaderContentsComponent } from './layout-header-contents.component';
import { Component } from '@angular/core';

@Component({selector: 'app-header', template: ''})
class HeaderComponent {}

describe('LayoutHeaderContentsComponent', () => {
  let component: LayoutHeaderContentsComponent;
  let fixture: ComponentFixture<LayoutHeaderContentsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        LayoutHeaderContentsComponent,
        HeaderComponent,
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LayoutHeaderContentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
