import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable } from 'rxjs';

import { LoginService } from '../api/login.service';

@Injectable()
export class JwtInterceptor implements HttpInterceptor {

    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        // add authorization header with jwt token if available
        const currentUser = this.loginService.currentUserValue;
        const currentUserToken = this.loginService.currentUserTokenValue;
        if (currentUser && currentUserToken) {
            request = request.clone({
                setHeaders: {
                  Authorization: `Bearer ${currentUserToken}`,
                }
            });
        }

        return next.handle(request);
    }

    constructor(private loginService: LoginService) { }
}
