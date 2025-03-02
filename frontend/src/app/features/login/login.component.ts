import { Component } from '@angular/core';
import { NgStyle } from '@angular/common';
import { IconDirective } from '@coreui/icons-angular';
import {
  ContainerComponent, RowComponent, ColComponent, CardGroupComponent, TextColorDirective,
  CardComponent, CardBodyComponent, FormDirective, InputGroupComponent, InputGroupTextDirective,
  FormControlDirective, ButtonDirective
} from '@coreui/angular';
import { AuthService } from '../../../shared/services/auth/auth.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms'

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
    standalone: true,
    imports: [
      ContainerComponent, RowComponent, ColComponent, CardGroupComponent, TextColorDirective,
      CardComponent, CardBodyComponent, FormDirective, InputGroupComponent, InputGroupTextDirective,
      IconDirective, FormControlDirective, ButtonDirective, NgStyle, FormsModule
    ],
    
})
export class LoginComponent {
  email: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router) {}

  onSubmit() {
    this.authService.login(this.email, this.password).subscribe({
      next: (response) => {
        console.log('Login successful', response);
        this.router.navigate(['/home']);
      },
      error: (error) => {
        console.error('Login failed', error);
      }
    });
  }
}
