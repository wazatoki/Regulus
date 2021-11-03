import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpService } from '../http.service';
import { SaveData } from '../models/search/save-data';
import { ConditionData } from '../models/search/condition-data';
import { Category } from '../models/search/category';
import { ComplexSearchItems } from '../models/search/complex-search-items';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchConditionService {

  updateFavoriteConditions(data: string[]): Observable<string[] | HttpErrorResponse> {
    return this.http.put<string[]>('/complexSearchCondition/updateFavoriteConditions', data);
  }

  findAllCategories(): Observable<Category[] | HttpErrorResponse> {
    return this.http.get<Category[]>('/complexSearchCondition/dataInputFormItems');
  }

  findComplexSearchItems(): Observable<ComplexSearchItems | HttpErrorResponse> {
    return this.http.get<ComplexSearchItems>('/complexSearchCondition/complexSearchItems');
  }

  findByCondition(condition: ConditionData): Observable<SaveData[] | HttpErrorResponse> {
    const data: Map<string, string> = new Map();
    data.set('condition', JSON.stringify(condition));
    return this.http.get<SaveData[]>('/complexSearchCondition', data);
  }

  // add(data: SaveData): Observable<SaveData> {
  //   return this.http.post<SaveData>( '/complexSearchCondition', data );
  // }

  // update(data: SaveData): Observable<SaveData> {
  //   return this.http.put<SaveData>( '/complexSearchCondition', data );
  // }

  delete(data: string[]): Observable<SaveData[] | HttpErrorResponse> {
    return this.http.delete<SaveData>('/complexSearchCondition', data);
  }

  constructor( private http: HttpService ) { }
}
