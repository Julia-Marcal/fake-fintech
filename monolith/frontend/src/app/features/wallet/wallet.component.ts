import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ToastService } from '../../../shared/services/toast/toast.service';
import { LoaderService } from '../../../shared/services/loader/loader.service';
import { ReactiveFormsModule } from '@angular/forms';
import { AuthService } from '../../../shared/services/auth/auth.service';
import { WalletService } from './wallet.service';
import { Router } from '@angular/router';

import { ChangeDetectionStrategy } from '@angular/core';
import {
  TextColorDirective,
  CardBodyComponent,
  CardComponent,
  CardHeaderComponent,
  ColComponent,
  RowComponent,
  DropdownComponent,
  DropdownToggleDirective,
  DropdownMenuDirective,
  DropdownItemDirective,
  ButtonDirective
} from '@coreui/angular';
import { ThemeService } from '../../../shared/services/themes/themes.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    TextColorDirective,
    CardComponent,
    CardHeaderComponent,
    CardBodyComponent,
    RowComponent,
    ReactiveFormsModule,
    ColComponent,
    DropdownComponent,
    DropdownToggleDirective,
    DropdownMenuDirective,
    DropdownItemDirective,
    ButtonDirective
  ],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class WalletComponent implements OnInit {
  isLoading = false;
  wallets: any[] = [];
  user: any = null;
  isDarkTheme$: Observable<boolean>;

  constructor(
    private authService: AuthService,
    private walletService: WalletService,
    private router: Router,
    private toastService: ToastService,
    private loaderService: LoaderService,
    private themeService: ThemeService,
    private cdr: ChangeDetectorRef
  ) {
    this.wallets = [];
    this.isDarkTheme$ = this.themeService.isDarkTheme$;
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
        this.cdr.markForCheck();
      },
      error: (err) => {
        console.error('Error fetching user wallets:', err);
        this.cdr.markForCheck();
      }
    });
  }

  visualizarWallet(wallet: any): void {
    if (this.isLoading) return;

    this.isLoading = true;
    this.loaderService.setLoading(true);

    console.log(wallet);


    this.router.navigate([`/wallet/${wallet.id}/stocks`]).then(() => {
      this.loaderService.setLoading(false);
    }).catch(() => {
      this.isLoading = false;
      this.loaderService.setLoading(false);
    });
  }

  excluirWallet(wallet: any): void {
    if (this.isLoading) return;

    this.isLoading = true;
    this.loaderService.setLoading(true);
  }
}