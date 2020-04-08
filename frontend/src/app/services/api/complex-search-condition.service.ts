import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpService } from '../http.service';
import { SaveData } from '../models/search/save-data';
import { ConditionData } from '../models/search/condition-data';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchConditionService {

  findById(id: string): Observable<SaveData> {
    const data: Map<string, string> = new Map();
    data.set('id', id);
    return this.http.get<SaveData>( '/complexSearchCondition', data );
  }

  findByCondition(condition: ConditionData): Observable<SaveData[]> {
    const data: Map<string, string> = new Map();
    data.set('condition',JSON.stringify(condition));
    return this.http.get<SaveData[]>('/complexSearchCondition', data);
  }

  findAll(): Observable<SaveData[]> {
    return this.http.get<SaveData[]>( '/complexSearchCondition' );
  }

  add(data: SaveData): Observable<SaveData> {
    return this.http.post<SaveData>( '/complexSearchCondition', data );
  }

  update(data: SaveData): Observable<SaveData> {
    return this.http.put<SaveData>( '/complexSearchCondition', data );
  }

  delete(data: string[]): Observable<SaveData[]> {
    return this.http.delete<SaveData>('/complexSearchCondition/delete', data);
  }

  constructor( private http: HttpService ) { }
}
