import { DebugElement } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { By } from '@angular/platform-browser';
import { NoticeDialogComponent } from './notice-dialog.component';

describe('NoticeDialogComponent', () => {
  let component: NoticeDialogComponent;
  let fixture: ComponentFixture<NoticeDialogComponent>;
  let debugElement: DebugElement;

  beforeEach( async( () => {
    const spy = jasmine.createSpyObj('MatDialogRef', ['close']);
    TestBed.configureTestingModule({
      imports: [
        MatDialogModule,
        NoopAnimationsModule,
      ],
      declarations: [ NoticeDialogComponent ],
      providers: [
        {
          provide: MatDialogRef,
          useValue: spy
        },
        {
          provide: MAT_DIALOG_DATA,
          useValue: {}
        },
      ],
    }).compileComponents();
  }));

  beforeEach( () => {
    fixture = TestBed.createComponent(NoticeDialogComponent);
    component = fixture.componentInstance;
    debugElement = fixture.debugElement;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('shows dialog title', () => {
    component.data.title = 'aaa';
    fixture.detectChanges();
    const title = debugElement.query(By.css('h1.dialog-title')).nativeElement.textContent;
    expect(title).toEqual('aaa');
  });

  it('shows dialog contents', () => {
    component.data.contents = 'bbb';
    fixture.detectChanges();
    const contents = debugElement.query(By.css('p.dialog-contents')).nativeElement.textContent;
    expect(contents).toEqual('bbb');
  });

});
