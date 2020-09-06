import { Injectable } from '@angular/core';
import { Subject, Observable }    from 'rxjs';

import { ConditionData } from '../models/search/condition-data';
import { SaveData } from '../models/search/save-data';
import { HttpService } from '../http.service';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchService {

  private complexSearchOrderedSouce = new Subject<SaveData>();
  private complexSearchSaveOrderedSouce = new Subject<SaveData>();

  complexSearchOrdered$ = this.complexSearchOrderedSouce.asObservable();
  complexSearchSaveOrdered$ = this.complexSearchSaveOrderedSouce.asObservable();

  initSaveDataObj(): SaveData {
    return {
      id: '',
      patternName: '',
      category: null,
      isDisclose: false,
      discloseGroupIDs: [],
      ownerID: '',
      conditionData: this.initConditionDataObj(),
      owner: {
        id: '',
        name: '',
      }
    };
  }

  initConditionDataObj(): ConditionData {
    return {
      searchStrings: [],
      displayItemList: [],
      searchConditionList: [],
      orderConditionList: [],
    };
  }

  orderComplexSearch(data: SaveData){
    this.complexSearchOrderedSouce.next(data);
  }

  addSearchCondition(data: SaveData):Observable<SaveData>{
    return this.http.post<SaveData>( '/complexSearchCondition', data );
  }

  updateSearchCondition(data: SaveData):Observable<SaveData>{
    return this.http.put<SaveData>( '/complexSearchCondition', data );
  }
  
  constructor( private http: HttpService ) { }
}
