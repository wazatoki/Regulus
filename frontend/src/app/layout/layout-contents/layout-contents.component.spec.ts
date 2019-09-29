import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LayoutContentsComponent } from './layout-contents.component';

describe('LayoutContentsComponent', () => {
  let component: LayoutContentsComponent;
  let fixture: ComponentFixture<LayoutContentsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LayoutContentsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LayoutContentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
