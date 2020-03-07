import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatCardModule } from '@angular/material/card';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, FormBuilder, FormControl, FormArray, FormGroup } from '@angular/forms';
import { ComplexSearchComponent } from './complex-search.component';
import { DebugElement, Component } from '@angular/core';

describe('ComplexSearchComponent', () => {
  let component: TestHostComponent;
  let fixture: ComponentFixture<TestHostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ComplexSearchComponent,
        TestHostComponent,
      ],
      imports: [
        FormsModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatCheckboxModule,
        MatCardModule,
        DragDropModule,
        NoopAnimationsModule,
      ],
      providers: [
        {
          provide: FormBuilder,
          useValue: new FormBuilder()
        },
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TestHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  fit('should create', () => {
    expect(component).toBeTruthy();
  });
});

@Component({
  template: `
  <app-complex-search
    [displayItemList]="displayItemList">
  </app-complex-search>`
})
class TestHostComponent {

  displayItemList = [
    {
      entityName: 'aaa',
      fieldName: 'AAA',
      viewValue: 'aaa-AAA',
      fieldType: 'number',
    },
    {
      entityName: 'bbb',
      fieldName: 'BBB',
      viewValue: 'bbb-BBB',
      fieldType: 'string',
    },
  ]
}
