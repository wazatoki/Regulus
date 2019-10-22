import { Injectable } from '@angular/core';
import { HttpService } from '../http.service';
import { Maker} from '../models/maker';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MakerService {

  constructor( private http: HttpService ) { }

  get(id: string): Observable<Maker> {
    return this.http.get<Maker>( 'maker', { id: id } );
  }
}
