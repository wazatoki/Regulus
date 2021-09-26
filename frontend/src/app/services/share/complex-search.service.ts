import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';

import { ConditionData } from '../models/search/condition-data';
import { SaveData } from '../models/search/save-data';
import { HttpService } from '../http.service';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchService {

  private complexSearchOrderedSouce = new Subject<ConditionData>();
  private complexSearchSaveOrderedSouce = new Subject<SaveData>();

  complexSearchOrdered$ = this.complexSearchOrderedSouce.asObservable();
  complexSearchSaveOrdered$ = this.complexSearchSaveOrderedSouce.asObservable();

  initSaveDataObj(): SaveData {
    return {
      id: '',
      patternName: '',
      category: null,
      isDisclose: false,
      discloseGroups: [],
      ownerID: '',
      conditionData: this.initConditionDataObj(),
      owner: {
        id: '',
        name: '',
        operatorUsableConditions: [],
      }
    };
  }

  initConditionDataObj(): ConditionData {
    return {
      searchStrings: [''],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
  }

  /*
  complexSearchOrdered$の購読先にdataを渡す
  */
  orderComplexSearch(data: ConditionData){
    this.complexSearchOrderedSouce.next(data);
  }

  addSearchCondition(data: SaveData): Observable<SaveData | HttpErrorResponse> {
    return this.http.post<SaveData>( '/complexSearchCondition', data );
  }

  updateSearchCondition(data: SaveData): Observable<SaveData | HttpErrorResponse> {
    return this.http.put<SaveData>( '/complexSearchCondition', data );
  }
  constructor( private http: HttpService ) { }
}
