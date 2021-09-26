import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material';

import { SelectSearchConditionComponent } from './select-search-condition.component';

describe('SelectSearchConditionComponent', () => {
  let component: SelectSearchConditionComponent;
  let fixture: ComponentFixture<SelectSearchConditionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SelectSearchConditionComponent ],
      imports: [ MatButtonModule ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SelectSearchConditionComponent);
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
