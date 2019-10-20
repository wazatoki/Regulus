import { Injectable } from '@angular/core';
import { HttpService } from '../http.service';
import { Maker} from '../model/maker';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MakerService {

  get(id: string): Observable<Maker> {

    return this.http.get<Maker>( 'maker', {'id': id} );
  
  }

  constructor( private http: HttpService ) { }
}
