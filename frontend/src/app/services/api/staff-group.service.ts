import { HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpService } from '../http.service';
import { StaffGroup } from '../models/group/staff-group';
import { ConditionData } from '../models/search/condition-data';

@Injectable({
  providedIn: 'root'
})
export class StaffGroupService {

  findByCondition(condition: ConditionData): Observable<StaffGroup[] | HttpErrorResponse> {
    const data: Map<string, string> = new Map();
    data.set('condition', JSON.stringify(condition));
    return this.http.get<StaffGroup[]>('/staffGroup', data);
  }

  add(data: StaffGroup): Observable<StaffGroup | HttpErrorResponse> {
    return this.http.post<StaffGroup>( '/staffGroup', data );
  }

  update(data: StaffGroup): Observable<StaffGroup | HttpErrorResponse> {
    return this.http.put<StaffGroup>( '/staffGroup', data );
  }

  delete(data: string[]): Observable<StaffGroup[] | HttpErrorResponse> {
    return this.http.delete<StaffGroup>('/staffGroup', data);
  }

  constructor( private http: HttpService ) { }
}
