import { Injectable } from '@angular/core';
import { HttpService } from '../http.service';
import { Maker} from '../models/maker';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MakerService {

  constructor( private http: HttpService ) { }

  findById(id: string): Observable<Maker> {
    return this.http.get<Maker>( 'maker', { id: id } );
  }

  findAll(): Observable<Maker[]> {
    return this.http.get<Maker[]>( 'maker' );
  }

  add(data: Maker): Observable<Maker> {
    return this.http.post<Maker>( 'maker', data );
  }
}
