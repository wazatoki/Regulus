import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OperatorUsableConditionsComponent } from './operator-usable-conditions.component';
import { createTestArray as createTeatSaveDataArray} from '../../../services/models/search/save-data.spec';
import { DebugElement } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatListModule } from '@angular/material/list';

describe('OperatorUsableConditionsComponent', () => {
  let component: OperatorUsableConditionsComponent;
  let fixture: ComponentFixture<OperatorUsableConditionsComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OperatorUsableConditionsComponent ],
      imports: [
        MatGridListModule,
        MatListModule,
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OperatorUsableConditionsComponent);
    component = fixture.componentInstance;
    component.usableConditions = createTeatSaveDataArray();
    dbElement = fixture.debugElement;
    element = dbElement.nativeElement;
    spyOn(component.selectedCondition, 'emit');
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain condition name', () => {
    expect(element.textContent).toContain('saveName1');
    expect(element.textContent).toContain('saveName2');
  });

  it('should event emit button clicked', () => {
    const buttonList = element.querySelectorAll('.mat-action-list button');
    const button = buttonList[0];
    button.dispatchEvent(new Event('click'));
    fixture.detectChanges();
    expect(component.selectedCondition.emit).toHaveBeenCalledWith(component.conditions[0]);
  });
});
