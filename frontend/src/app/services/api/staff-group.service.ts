import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpService } from '../http.service';
import { StaffGroup } from '../models/group/staff-group';

@Injectable({
  providedIn: 'root'
})
export class StaffGroupService {

  delete(data: string[]): Observable<StaffGroup[]> {
    return this.http.delete<StaffGroup>('/staffGroup', data);
  }

  constructor( private http: HttpService ) { }
}
