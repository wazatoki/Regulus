import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpService } from '../http.service';
import { StaffGroup } from '../models/group/staff-group';
import { ConditionData } from '../models/search/condition-data';

@Injectable({
  providedIn: 'root'
})
export class StaffGroupService {

  findByCondition(condition: ConditionData): Observable<StaffGroup[]> {
    const data: Map<string, string> = new Map();
    data.set('condition',JSON.stringify(condition));
    return this.http.get<StaffGroup[]>('/staffGroup', data);
  }
  
  delete(data: string[]): Observable<StaffGroup[]> {
    return this.http.delete<StaffGroup>('/staffGroup', data);
  }

  constructor( private http: HttpService ) { }
}
