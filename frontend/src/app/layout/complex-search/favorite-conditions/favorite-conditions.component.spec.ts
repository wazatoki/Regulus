import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule, MatListModule } from '@angular/material';

import { FavoriteConditionsComponent } from './favorite-conditions.component';
import { createTestArray as createTeatSaveDataArray} from '../../../services/models/search/save-data.spec';
import { DebugElement } from '@angular/core';

describe('FavoriteConditionsComponent', () => {
  let component: FavoriteConditionsComponent;
  let fixture: ComponentFixture<FavoriteConditionsComponent>;
  let dbElement: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FavoriteConditionsComponent ],
      imports: [
        MatButtonModule,
        MatListModule,
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FavoriteConditionsComponent);
    component = fixture.componentInstance;
    component.favoriteConditions = createTeatSaveDataArray();
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
    expect(component.selectedCondition.emit).toHaveBeenCalledWith(component.favoriteConditions[0].conditionData);
  });


});
