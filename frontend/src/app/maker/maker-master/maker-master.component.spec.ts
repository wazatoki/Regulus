import { async, ComponentFixture, TestBed, tick, fakeAsync } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MakerMasterComponent } from './maker-master.component';
import { MakerSearchComponent } from '../maker-search/maker-search.component';
import { LayoutModule } from '../../layout/layout.module';
import { MatDialogModule } from '@angular/material/dialog';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MakerService } from '../../services/api/maker.service';
import { MakerCondition } from '../../services/models/maker/maker-condition'
import { Maker } from '../../services/models/maker/maker';

describe('MakerMasterComponent', () => {
  let component: MakerMasterComponent;
  let dbElement: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<MakerMasterComponent>;

  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('MakerService', ['findByCondition']);

    TestBed.configureTestingModule({
      declarations: [
        MakerMasterComponent,
        MakerSearchComponent,
      ],
      imports: [
        BrowserAnimationsModule,
        LayoutModule,
        MatTableModule,
        MatListModule,
        MatDialogModule,
        MatCheckboxModule,
        MatPaginatorModule,
      ],
      providers: [
        { provide: MakerService, useValue: spy },
        MakerCondition,
      ],
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerMasterComponent);
    component = fixture.componentInstance;
    dbElement = fixture.debugElement;
    element = dbElement.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain table', () => {

    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    expect(element.textContent).toContain(' 製造販売業者 ');

    component.onFetchedMakers(testData);
    fixture.detectChanges();
    expect(element.textContent).toContain('Test Maker1');
    expect(element.textContent).toContain('Test Maker2');
  });

  it('select item as all checked', () => {

    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];

    component.onFetchedMakers(testData);
    fixture.detectChanges();

    fixture.whenStable().then(() => {
      fixture.detectChanges();

      const checkboxList = element.querySelectorAll('.mat-checkbox input')
      const checkbox = checkboxList[0];
      checkbox.dispatchEvent(new Event('click'));
      fixture.detectChanges();
      expect(component.selection.selected.length).toBe(2);

    });
  });
});
