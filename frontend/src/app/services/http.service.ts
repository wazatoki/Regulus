import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  readonly HOST_URL: string = window.location.host

  getHttpParams(data: { [key: string]: string }): HttpParams {

    const params: HttpParams = new HttpParams();
    for (let key in data) {
      params.set(key, data[key])
    }
    return params;
  }

  get<T>(path: string, data: { [key: string]: string } = {}): Observable<T> {

    return this.client.get<T>(`${this.HOST_URL}${path}`, { params: this.getHttpParams(data) })
      .pipe(
        retry(3),
        catchError(this.handleError)
      );
  }

  private handleError(error: HttpErrorResponse) {
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
    return throwError(
      'Something bad happened; please try again later.');
  };

  constructor(private client: HttpClient) { }
}
