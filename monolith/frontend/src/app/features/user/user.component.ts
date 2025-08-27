import { Component, OnInit } from '@angular/core';
import { TextColorDirective, CardComponent, CardHeaderComponent, CardBodyComponent, RowComponent } from '@coreui/angular';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { UserService, User } from './user.service';
import { Router } from '@angular/router';
import { AuthService } from '../../../shared/services/auth/auth.service';
import { ToastService } from '../../../shared/services/toast/toast.service';

@Component({
  selector: 'app-user',
  standalone: true,
  imports: [
    TextColorDirective,
    CardComponent,
    CardHeaderComponent,
    CardBodyComponent,
    RowComponent,
    ReactiveFormsModule
  ],
  templateUrl: './user.component.html',
  styleUrl: './user.component.scss'
})
export class UserComponent implements OnInit {
  userForm: FormGroup;
  user: any = null;


  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private authService: AuthService,
    private router: Router,
    private toastService: ToastService,
  ) {
    this.userForm = this.fb.group({
      name: ['', [Validators.required]],
      last_name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]],
      age: ['', [Validators.pattern('^[0-9]*$')]]
    });
  }

  ngOnInit(): void {
    this.user = this.authService.getDecodedToken();

    if (!this.user || !this.user.sub) {
      this.toastService.showToast({
          title: 'Erro',
          message: 'Erro ao obter os dados do usuário. Você não está logado.',
          duration: 3000,
          position: 'top-end'
        });


      setTimeout(() => {
        this.authService.logout();
        this.router.navigate(['/login']);
      }, 1500);
      return;
    }

    this.userService.getCurrentUser(this.user.sub).subscribe({
      next: (user: User) => {
        this.userForm.setValue({
          name: user.name || '',
          last_name: user.last_name || '',
          email: user.email || '',
          age: user.age || ''
        });
      },
      error: (err) => {
        console.error('Error fetching user data:', err);
      }
    });
  }

  onSubmit(): void {
    if (this.userForm.valid) {
      this.userService.updateUser(this.user.sub, this.userForm.value).subscribe({
        next: (user: User) => {
          this.toastService.showToast({
            title: 'Success',
            message: 'User updated successfully',
            duration: 3000,
            position: 'top-end'
          });
        },
        error: (err) => {
          console.error('Error updating user data:', err);
          this.toastService.showToast({
            title: 'Error',
            message: 'Failed to update user',
            duration: 3000,
            position: 'top-end'
          });
        }
      });
    }
  }
}
