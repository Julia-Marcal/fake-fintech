import { Component } from '@angular/core';
import { ToastService } from '../../../shared/services/toast/toast.service';
import { LoaderService } from '../../../shared/services/loader/loader.service';
import { ReactiveFormsModule } from '@angular/forms';
import { AuthService } from '../../../shared/services/auth/auth.service';
import { WalletService } from './wallet.service';
import { Router } from '@angular/router';
import { WidgetsDropdownComponent } from '../../../app/base/widgets/widgets-dropdown/widgets-dropdown.component';

import { AfterContentInit, ChangeDetectionStrategy, ChangeDetectorRef, inject } from '@angular/core';
import { WidgetsBrandComponent } from '../../../app/base/widgets/widgets-brand/widgets-brand.component';
import { IconDirective } from '@coreui/icons-angular';
import { WidgetsEComponent } from '../../../app/base/widgets/widgets-e/widgets-e.component';
import {
  TextColorDirective,
  CardBodyComponent,
  CardComponent,
  CardGroupComponent,
  CardHeaderComponent,
  ColComponent,
  ProgressComponent,
  TemplateIdDirective,
  WidgetStatBComponent,
  WidgetStatCComponent,
  WidgetStatFComponent,
  WidgetStatAComponent,
  RowComponent
} from '@coreui/angular';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss'],
  standalone: true,
  imports: [
    TextColorDirective,
    CardComponent,
    CardHeaderComponent,
    CardBodyComponent,
    RowComponent,
    ReactiveFormsModule,
    WidgetsDropdownComponent,

    WidgetsBrandComponent,
    IconDirective,
    WidgetsEComponent,
    WidgetStatBComponent,
    ProgressComponent,
    WidgetStatFComponent,
    TemplateIdDirective,
    CardGroupComponent,
    WidgetStatCComponent,
    WidgetStatAComponent,
    ColComponent
  ],
  changeDetection: ChangeDetectionStrategy.Default
})
export class WalletComponent implements AfterContentInit {
  wallets: any[] = [];
  user: any = null;
  private changeDetectorRef = inject(ChangeDetectorRef);

  constructor(private authService: AuthService, private walletService: WalletService, private router: Router, private toastService: ToastService, private loaderService: LoaderService) {
    this.wallets = [];
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

    this.walletService.getUserWallets(this.user.sub).subscribe({
      next: (wallets) => {
        this.wallets = wallets;
        console.log(this.wallets);
        
      },
      error: (err) => {
        console.error('Error fetching user wallets:', err);
      }
    });
  }

  ngAfterContentInit(): void {
    this.changeDetectorRef.detectChanges();
  }

}