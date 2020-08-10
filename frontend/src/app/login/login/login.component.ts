import { Component, OnInit, ElementRef } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm = new FormGroup({
    userID: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });
  htmlEl: HTMLElement;
  isSubmitFocus: boolean = false;

  onSubmit() {
    if (this.loginForm.valid){
      console.log(this.loginForm.value)
      
    }
  }

  onEnterUserID(event: any) {
    event.preventDefault()
    const passEl: HTMLElement = this.htmlEl.querySelector('input[type="password"]');
    passEl.focus();
  }

  onEnterPassword(event: any) {
    event.preventDefault()
    this.isSubmitFocus = true;
  }

  onSubmitFocusout(){
    this.isSubmitFocus = false;
  }

  constructor(private el: ElementRef) {
    this.htmlEl = this.el.nativeElement
  }

  ngOnInit() {
  }

}
