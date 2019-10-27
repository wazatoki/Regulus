import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { By } from '@angular/platform-browser';
import { LayoutModule } from '../../layout/layout.module';
import { MakerSearchComponent } from './maker-search.component';
import { MakerService } from '../../services/api/maker.service';
import { MakerCondition } from '../../services/models/maker/maker-condition';
import { SearchComponent } from 'src/app/layout/search/search.component';
import { Maker } from '../../services/models/maker/maker';
import { of } from 'rxjs';

describe('MakerSearchComponent', () => {
  let component: MakerSearchComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<MakerSearchComponent>;
  let makerServiceSpy: jasmine.SpyObj<MakerService>;
  let makerCondition: MakerCondition;


  beforeEach(async(() => {

    const spy = jasmine.createSpyObj('MakerService', ['findByCondition']);

    TestBed.configureTestingModule({
      declarations: [ MakerSearchComponent ],
      imports: [ LayoutModule ],
      providers: [
        { provide: MakerService, useValue: spy },
        MakerCondition,
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakerSearchComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement; 
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should onSearch set keywords to makerCondition.searchStrings', () => {
    makerServiceSpy = TestBed.get(MakerService);
    makerCondition = TestBed.get(MakerCondition);
    const testData: Maker[] = [
      { id: 'testid1', name: 'Test Maker1' },
      { id: 'testid2', name: 'Test Maker2' },
    ];
    const stubValue = of(testData)
    makerServiceSpy.findByCondition.and.returnValue(stubValue);
    component.onSearch('search word');
    expect(makerCondition.searchStrings).toEqual( 'search word' );
    expect(component.makers).toBe(testData);
  });
});
