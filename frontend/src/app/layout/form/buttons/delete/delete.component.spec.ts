import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';

import { DeleteComponent } from './delete.component';

describe('DeleteComponent', () => {
  let component: DeleteComponent;
  let fixture: ComponentFixture<DeleteComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeleteComponent ],
      imports: [ MatButtonModule ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeleteComponent);
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
