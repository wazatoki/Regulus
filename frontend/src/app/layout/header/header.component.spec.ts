import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HeaderComponent } from './header.component';
import { LoginService } from 'src/app/services/api/login.service';
import { BehaviorSubject } from 'rxjs';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  let spy:jasmine.SpyObj<LoginService>;

  beforeEach(async(() => {
    
    const spy = jasmine.createSpyObj('LoginService', ['currentUser', 'currentUserToken', 'currentUserValue', 'currentUserTokenValue']);

    TestBed.configureTestingModule({
      declarations: [ HeaderComponent ],
      imports: [ RouterTestingModule ],
      providers: [
        { provide: LoginService, useValue: spy },
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    spy = TestBed.get(LoginService);
    spy.currentUserToken= new BehaviorSubject<string>('').asObservable()

    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
