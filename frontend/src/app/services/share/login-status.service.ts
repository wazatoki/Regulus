import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginStatusService {

  status: boolean
  jwtToken: string

  constructor() {
    this.status = false;
    this.jwtToken = null;
  }
}
