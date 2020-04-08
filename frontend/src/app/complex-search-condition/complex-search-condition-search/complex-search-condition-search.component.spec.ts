import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplexSearchConditionSearchComponent } from './complex-search-condition-search.component';
import { ComplexSearchConditionService } from 'src/app/services/api/complex-search-condition.service';
import { LayoutModule } from 'src/app/layout/layout.module';
import { MatButtonModule } from '@angular/material/button';
import { FlexLayoutModule } from '@angular/flex-layout';

describe('ComplexSearchConditionSearchComponent', () => {
  let component: ComplexSearchConditionSearchComponent;
  let fixture: ComponentFixture<ComplexSearchConditionSearchComponent>;

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('ComplexSearchConditionService', ['findByCondition']);

    TestBed.configureTestingModule({
      declarations: [ ComplexSearchConditionSearchComponent ],
      imports: [
        LayoutModule,
        FlexLayoutModule,
      ],
      providers: [
        { provide: ComplexSearchConditionService, useValue: spy },
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplexSearchConditionSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
