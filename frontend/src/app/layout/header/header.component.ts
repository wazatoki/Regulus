import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/services/api/login.service';
import { MatDialog } from '@angular/material/dialog';
import { AlertDialogComponent } from 'src/app/layout/dialog/alert-dialog/alert-dialog.component';
import { TRUE } from 'src/app/services/models/enum/boolean';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  isLoginable: boolean;
  isLogoutable: boolean;

  onLogout(event: any): void {
    event.preventDefault();

    const dialogref = this.dialog.open(AlertDialogComponent, {
      data: {
        title: '確認',
        contents: 'ログアウトしますか？',
      }
    });

    dialogref.afterClosed().subscribe(result => {
      if (result === TRUE) {
        this.loginService.logout();
        this.router.navigate(['/']);
      }
    });
  }

  constructor(
    private loginService: LoginService,
    private dialog: MatDialog,
    private router: Router,
  ) {
    this.loginService.currentUserToken.subscribe(token => {
      if (token === '') {
        this.isLoginable = true;
        this.isLogoutable = false;
      } else {
        this.isLoginable = false;
        this.isLogoutable = true;
      }
    });
  }

  ngOnInit() {
  }

}
