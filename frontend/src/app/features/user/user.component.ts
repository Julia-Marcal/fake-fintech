import { Component, OnInit } from '@angular/core';
import { TextColorDirective, CardComponent, CardHeaderComponent, CardBodyComponent, RowComponent } from '@coreui/angular';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { UserService, User } from './user.service';

import { AuthService } from '../../../shared/services/auth/auth.service';

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

  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private authService: AuthService
  ) {
    this.userForm = this.fb.group({
      name: ['', [Validators.required]],
      lastName: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]],
      age: ['', [Validators.pattern('^[0-9]*$')]]
    });
  }

  ngOnInit(): void {
    const user = this.authService.getDecodedToken();

    this.userService.getCurrentUser(user.sub).subscribe({
      next: (user: User) => {
        this.userForm.setValue({
          name: user.name || '',
          lastName: user.last_name || '',
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
      console.log('Form Submitted!', this.userForm.value);
    }
  }
}
