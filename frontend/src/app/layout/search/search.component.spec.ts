import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FlexLayoutModule } from '@angular/flex-layout';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {FormsModule} from '@angular/forms';
import { SearchComponent } from './search.component';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

describe('SearchComponent', () => {
  let component: SearchComponent;
  let elementd: DebugElement;
  let element: HTMLElement;
  let fixture: ComponentFixture<SearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SearchComponent ],
      imports: [
        FlexLayoutModule,
        MatButtonModule,
        MatInputModule,
        FormsModule,
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SearchComponent);
    component = fixture.componentInstance;
    elementd = fixture.debugElement; 
    element = elementd.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should explain input.search-strings', () => {
    expect(element.querySelector('input').className).toContain('search-strings');
  })

  it('should explain search button', () => {
    expect(element.querySelector('button').textContent).toContain('検索');
  })

  it('should bind input value to seatchStrings', () => {
    const inputDe: DebugElement = elementd.query(By.css('input.search-strings'));
    const input: HTMLInputElement = inputDe.nativeElement;
    input.value = 'test words';
    input.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    expect(component.searchStrings).toBe('test words');
  })

  it('should emit searchClick event', () => {
    const buttonDe: DebugElement = elementd.query(By.css('button'));
    const button: HTMLInputElement = buttonDe.nativeElement;
    component.searchStrings = 'search text';
    component.searchClick.subscribe( (data: string) => {
      expect(data).toEqual(component.searchStrings);
    });
    button.dispatchEvent(new Event('click'));
    fixture.detectChanges();
  });
});
