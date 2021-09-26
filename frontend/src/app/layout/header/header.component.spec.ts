import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HeaderComponent } from './header.component';
import { LoginService } from 'src/app/services/api/login.service';
import { BehaviorSubject } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  const loginServiceSpy: jasmine.SpyObj<LoginService> = jasmine.createSpyObj(
    'LoginService',
    ['currentUser', 'currentUserToken', 'currentUserValue', 'currentUserTokenValue']);
  loginServiceSpy.currentUserToken = new BehaviorSubject<string>('').asObservable();
  const dialogspy = jasmine.createSpyObj('MatDialog', ['open']);

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [HeaderComponent],
      imports: [RouterTestingModule],
      providers: [
        { provide: LoginService, useValue: loginServiceSpy },
        { provide: MatDialog, useValue: dialogspy },
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
