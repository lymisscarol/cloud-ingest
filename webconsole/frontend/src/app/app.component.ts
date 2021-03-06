import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params, NavigationExtras } from '@angular/router';
import { AuthService } from './auth/auth.service';
import { MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  isProjectSelected: boolean;
  authServiceInitialized: boolean;
  loggedIn: boolean;
  userName: string;
  gcsProjectId: string;

  constructor(private authService: AuthService,
              private router: Router, private route: ActivatedRoute,
              private snackBar: MatSnackBar) { }

  ngOnInit() {
    this.route.queryParams.subscribe((params: Params) => {
      // TODO(b/65209049): Validate the selected project is a valid project,
      // and may be provide a list of the current signed in user projects.
      // TODO(b/65209107): Auto populate the recent used project if a project
      // is not selected.
      const projectId = params['project'];
      this.isProjectSelected = (projectId != null && projectId !== '');
    });

    this.authService.init();
    this.authService.loadSignInStatus().then((loggedIn) => {
      this.authServiceInitialized = true;
      this.loggedIn = loggedIn;
      if (this.loggedIn) {
        this.userName = this.authService.getCurrentUser().Name;
      }
    }).catch((error) => {
      this.snackBar.open(`There was an error loading your sign in status: ${error.error}`, 'Dismiss');
    });
  }

  signOut() {
    this.authService.signOut().then(() => {
      this.loggedIn = false;
      this.router.navigate(['/'], { queryParamsHandling: 'merge' });
    }).catch((error) => {
      this.snackBar.open(`There was an error signing out of your account: ${error.error}`, 'Dismiss');
    });
  }

  signIn() {
    this.authService.signIn().then(() => {
      this.loggedIn = true;
      this.userName = this.authService.getCurrentUser().Name;
    }).catch((error) => {
      this.snackBar.open(`There was an error signing into your account: ${error.error}`, 'Dismiss');
    });
  }
}
