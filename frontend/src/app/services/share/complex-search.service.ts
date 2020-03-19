import { Injectable } from '@angular/core';
import { Subject }    from 'rxjs';

import { ConditionData } from '../models/search/condition-data';

@Injectable({
  providedIn: 'root'
})
export class ComplexSearchService {

  private complexSearchOrderedSouce = new Subject<ConditionData>();

  complexSearchOrdered$ = this.complexSearchOrderedSouce.asObservable();

  orderComplexSearch(data: ConditionData){
    this.complexSearchOrderedSouce.next(data);
  }
  
  constructor() { }
}
