import { Component } from '@angular/core';
import { NgStyle } from '@angular/common';
import { ToastService } from '../../../shared/services/toast/toast.service';
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
  isLoading = false;
  email: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router, private toastService: ToastService) { }

  onSubmit() {
    if(this.isLoading) return
    this.isLoading = true
    this.authService.login(this.email, this.password).subscribe({
      next: (response) => {
        this.toastService.showToast({
          title: 'Sucesso',
          message: 'Você conseguiu logar!',
          duration: 3000, 
          position: 'top-end'
        });

        this.router.navigate(['/home']).then(() => {
          this.isLoading = false; 
        }).catch(() => {
          this.isLoading = false; 
        });
      },
      error: (error) => {
        console.error('Login failed', error);
        this.isLoading = false; 
      }
    });
  }

  onRegister() {
    if(this.isLoading) return
    this.isLoading = true

    this.router.navigate(['/register']).then(() => {
      this.isLoading = false; 
    }).catch(() => {
      this.isLoading = false; 
    });
  }
}
