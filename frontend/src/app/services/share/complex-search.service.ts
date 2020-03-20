import { Injectable } from '@angular/core';
import { Subject }    from 'rxjs';

import { ConditionData } from '../models/search/condition-data';
import { SaveData } from '../models/search/save-data';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchService {

  private complexSearchOrderedSouce = new Subject<ConditionData>();
  private complexSearchSaveOrderedSouce = new Subject<SaveData>();

  complexSearchOrdered$ = this.complexSearchOrderedSouce.asObservable();
  complexSearchSaveOrdered$ = this.complexSearchSaveOrderedSouce.asObservable();

  orderComplexSearch(data: ConditionData){
    this.complexSearchOrderedSouce.next(data);
  }

  orderComplexSearchSave(data: SaveData){
    this.complexSearchSaveOrderedSouce.next(data);
  }
  
  constructor() { }
}
