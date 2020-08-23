import { Component, OnInit, ElementRef } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { first } from 'rxjs/operators';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { LoginService } from '../../services/api/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  htmlEl: HTMLElement;
  isSubmitFocus: boolean = false;

  returnUrl: string;
  error = '';

  onSubmit() {
    if (this.loginForm.invalid) {
      return;
    }

    this.loginService.login(this.loginForm.value).pipe(first()).subscribe(
      data => {
        if(data){
          this.router.navigate([this.returnUrl])
        }else{
          this.error = 'IDまたはパスワードが違います。'
        }
        
      },
      error => {
        this.error = error
      }
    );
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

  onSubmitFocusout() {
    this.isSubmitFocus = false;
  }

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private loginService: LoginService,
    private el: ElementRef) {

    if (this.loginService.currentUserValue) {
      this.router.navigate(['/']);
    }
  }

  ngOnInit() {
    this.loginForm = new FormGroup({
      id: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    });

    // get return url from route parameters or default to '/'
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';

    this.htmlEl = this.el.nativeElement
  }

}
