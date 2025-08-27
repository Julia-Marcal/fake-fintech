import { Component } from '@angular/core';
import { NgStyle } from '@angular/common';
import { ToastService } from '../../../shared/services/toast/toast.service';
import { LoaderService } from '../../../shared/services/loader/loader.service';
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
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss'],
  standalone: true,
  imports: [
    ContainerComponent, RowComponent, ColComponent, CardGroupComponent, TextColorDirective,
    CardComponent, CardBodyComponent, FormDirective, InputGroupComponent, InputGroupTextDirective,
    IconDirective, FormControlDirective, ButtonDirective, NgStyle, FormsModule
  ],

})
export class WalletComponent {
  isLoading = false;
  email: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router, private toastService: ToastService, private loaderService: LoaderService) { }

}
