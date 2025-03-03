import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IconDirective } from '@coreui/icons-angular';
import { ContainerComponent, RowComponent, ColComponent, TextColorDirective, CardComponent, CardBodyComponent, FormDirective, InputGroupComponent, InputGroupTextDirective, FormControlDirective, ButtonDirective } from '@coreui/angular';
import { FormBuilder, FormGroup, Validators, AbstractControl, ValidationErrors } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
    imports: [
      CommonModule, ReactiveFormsModule, ContainerComponent, RowComponent, 
      ColComponent, TextColorDirective, CardComponent, CardBodyComponent, 
      FormDirective, InputGroupComponent, InputGroupTextDirective, IconDirective, 
      FormControlDirective, ButtonDirective
    ]
})
export class RegisterComponent implements OnInit {
  registerForm!: FormGroup;
  
  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.registerForm = this.fb.group({
      name: ['', Validators.required],
      lastName: ['', Validators.required],
      age: ['', [Validators.required, Validators.min(18)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(8)]],
      confirmPassword: ['', Validators.required]
    }, { validators: this.passwordMatchValidator });
  }

  passwordMatchValidator(control: AbstractControl): ValidationErrors | null {
    const password = control.get('password')?.value;
    const confirmPassword = control.get('confirmPassword')?.value;
    return password === confirmPassword ? null : { passwordsMismatch: true };
  }

  onSubmit(): void {
    if (this.registerForm.valid) {
      const formData = this.registerForm.value;
      console.log('Form Data:', formData);
    } else {
      console.log('Form is invalid');
    }
  }
}
