import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LayoutHeaderContentsComponent } from './layout-header-contents.component';

describe('LayoutHeaderContentsComponent', () => {
  let component: LayoutHeaderContentsComponent;
  let fixture: ComponentFixture<LayoutHeaderContentsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LayoutHeaderContentsComponent ]
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
