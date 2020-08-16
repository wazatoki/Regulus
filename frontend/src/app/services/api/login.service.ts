import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpService } from '../http.service';
import { Staff } from '../models/staff/staff'

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private user: Staff;
  private token: string;
  private currentUserSubject: BehaviorSubject<Staff>;
  private currentUserTokenSubject: BehaviorSubject<string>;

  public currentUser: Observable<Staff>;
  public currentUserToken: Observable<string>;

  public get currentUserValue(): Staff {
    return this.currentUserSubject.value;
  }

  public get currentUserTokenValue(): string {
    return this.currentUserTokenSubject.value;
  }

  login(data: { userID: string, password: string }): Observable<boolean> {

    return this.http.post<any>('/login', data).pipe(map(data => {
      // store user details and jwt token in local storage to keep user logged in between page refreshes
      if (data.jwtToken) {
        this.currentUserTokenSubject.next(data.jwtToken);
        this.currentUserSubject.next(data.staff);
        this.token = data.jwtToken;
        this.user = data.staff
        return true
      }
      return false
    }));
  }

  logout() {
    // remove user from local storage to log user out
    this.token = '';
    this.user = null;
    this.currentUserSubject.next(null);
    this.currentUserTokenSubject.next(null);
  }

  constructor(
    private http: HttpService) {

    this.user = null;
    this.token = '';
    this.currentUserSubject = new BehaviorSubject<Staff>(this.user);
    this.currentUser = this.currentUserSubject.asObservable();
    this.currentUserTokenSubject = new BehaviorSubject<string>(this.token);
    this.currentUserToken = this.currentUserTokenSubject.asObservable();
  }
}
