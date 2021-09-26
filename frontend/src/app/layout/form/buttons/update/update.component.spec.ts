import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';

import { UpdateComponent } from './update.component';

describe('UpdateComponent', () => {
  let component: UpdateComponent;
  let fixture: ComponentFixture<UpdateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UpdateComponent ],
      imports: [ MatButtonModule ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('set other name', () => {
    component.buttonLabelName = 'testLabel';
    fixture.detectChanges();
    const htmlel: HTMLElement = fixture.nativeElement;
    expect(htmlel.textContent).toContain('testLabel');
  });
});
