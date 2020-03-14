import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { SubmitComponent } from './submit.component';
import { MatButtonModule } from '@angular/material/button';

describe('SubmitComponent', () => {
  let component: SubmitComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<SubmitComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SubmitComponent ],
      imports: [ MatButtonModule ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SubmitComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement;
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('input buttonLabelType as noting', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    fixture.detectChanges();
    expect(button.textContent).toContain('送信');
  });

  it('input buttonLabelType as save', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    component.buttonLabelType = 'save';
    fixture.detectChanges();
    expect(button.textContent).toContain('保存');
  });

  it('input buttonLabelType as insert', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    component.buttonLabelType = 'insert';
    fixture.detectChanges();
    expect(button.textContent).toContain('新規作成');
  });

  it('input buttonLabelType as update', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    component.buttonLabelType = 'update';
    fixture.detectChanges();
    expect(button.textContent).toContain('更新');
  });

  it('click button element', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLElement = buttonDe.nativeElement;
    component.clicked.subscribe( (data: string) => {
      expect(data).toEqual('submitClicked');
    });
    button.dispatchEvent(new Event('click'));
    fixture.detectChanges();
  });
});
