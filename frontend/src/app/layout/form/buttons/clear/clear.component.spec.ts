import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { ClearComponent } from './clear.component';
import { MatButtonModule } from '@angular/material/button';

describe('ClearComponent', () => {
  let component: ClearComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<ClearComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClearComponent ],
      imports: [ MatButtonModule ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClearComponent);
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
      expect(data).toEqual('clearClicked');
    });
    button.dispatchEvent(new Event('click'));
    fixture.detectChanges();
  });
});
