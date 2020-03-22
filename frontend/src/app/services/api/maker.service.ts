import { Injectable } from '@angular/core';
import { HttpService } from '../http.service';
import { Maker } from '../models/maker/maker';
import { ConditionData } from '../models/search/condition-data';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MakerService {

  constructor( private http: HttpService ) { }

  findById(id: string): Observable<Maker> {
    const data: Map<string, string> = new Map();
    data.set('id', id);
    return this.http.get<Maker>( 'maker', data );
  }

  findByCondition(condition: ConditionData): Observable<Maker[]> {
    const data: Map<string, string> = new Map();
    data.set('condition',JSON.stringify(condition));
    return this.http.get<Maker[]>('maker', data);
  }

  findAll(): Observable<Maker[]> {
    return this.http.get<Maker[]>( 'maker' );
  }

  add(data: Maker): Observable<Maker> {
    return this.http.post<Maker>( 'maker', data );
  }

  update(data: Maker): Observable<Maker> {
    return this.http.put<Maker>( 'maker', data );
  }

  delete(data: string[]): Observable<Maker[]> {
    return this.http.delete<Maker>('maker/delete', data);
  }
}
