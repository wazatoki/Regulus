import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpService } from '../http.service';
import { LoginStatusService } from './login-status.service'

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  login(data: { userID: string, password: string }): Observable<boolean> {

    return new Observable(observer => {

      this.http.post<any>('/login', data).subscribe(data => {
        this.loginStatus.jwtToken = data;
        this.loginStatus.status = true;
        observer.next(this.loginStatus.status)
      }, err => {
        this.loginStatus.status = false;
        this.loginStatus.jwtToken = null;
        observer.next(this.loginStatus.status)
      })

    });

  }

  constructor(
    private http: HttpService,
    private loginStatus: LoginStatusService) { }
}
