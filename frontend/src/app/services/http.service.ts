import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  readonly BASE_URL = window.location.protocol + '//' + window.location.host + '/';

  getHttpParams(data: {[key: string]: string}): HttpParams {

    const params: HttpParams = new HttpParams();
    for (let key in data){
      params.set(key,data[key])
    }
    return params;
  }

  get<T>(path: string, data: {[key: string]: string}): Observable<T>{

    return this.client.get<T>(this.BASE_URL + path, {params: this.getHttpParams(data)});
  }

  constructor(private client: HttpClient) { }
}
