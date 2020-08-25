import { Component, OnInit } from '@angular/core';
import { LoginService } from 'src/app/services/api/login.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  isLoginable: boolean
  isLogoutable: boolean

  constructor(
    private loginService :LoginService
  ) { 
    this.loginService.currentUserToken.subscribe( token => {
      if (token === '') {
        this.isLoginable = true;
        this.isLogoutable = false;
      }else{
        this.isLoginable = false;
        this.isLogoutable = true;
      }
    });
  }

  ngOnInit() {
  }

}
