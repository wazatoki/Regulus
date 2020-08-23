import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { HttpClient, HttpParams, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
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

  private handleError(error: HttpErrorResponse) {
    console.log(error)
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    // return an observable with a user-facing error message
    if (error.status === 401){
      return throwError(
        'login incorrect.\nID or password is mismatch.');  
    }
    return throwError(
      'Something bad happened; please try again later.');
  };

  login(data: { id: string, password: string }): Observable<boolean> {

    return this.client.post<any>('http://' + window.location.host + '/login', data).pipe(map(data => {
      // store user details and jwt token in local storage to keep user logged in between page refreshes
      if (data.jwtToken) {
        this.currentUserTokenSubject.next(data.jwtToken);
        this.currentUserSubject.next(data.staff);
        this.token = data.jwtToken;
        this.user = data.staff
        return true
      }
      return false
    }),
    catchError(this.handleError)
    );
  }

  logout() {
    // remove user from local storage to log user out
    this.token = '';
    this.user = null;
    this.currentUserSubject.next(null);
    this.currentUserTokenSubject.next(null);
  }

  constructor(
    private client: HttpClient) {

    this.user = null;
    this.token = '';
    this.currentUserSubject = new BehaviorSubject<Staff>(this.user);
    this.currentUser = this.currentUserSubject.asObservable();
    this.currentUserTokenSubject = new BehaviorSubject<string>(this.token);
    this.currentUserToken = this.currentUserTokenSubject.asObservable();
  }
}
