import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { CancelComponent } from './cancel.component';
import { MatButtonModule } from '@angular/material/button';

describe('CancelComponent', () => {
  let component: CancelComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<CancelComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CancelComponent ],
      imports: [ MatButtonModule ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CancelComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement;
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('click button element', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    component.clicked.subscribe( (data: string) => {
      expect(data).toEqual('cancelClicked');
    });
    button.dispatchEvent(new Event('click'));
    fixture.detectChanges();
  });
});
