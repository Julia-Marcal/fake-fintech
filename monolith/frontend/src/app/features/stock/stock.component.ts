import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';
import { CommonModule } from '@angular/common';
import { StockService } from './stock.service';
import { ChangeDetectionStrategy } from '@angular/core';
import { ThemeService } from '../../../shared/services/themes/themes.service';

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


@Component({
  selector: 'app-stock',
  templateUrl: './stock.component.html',
  styleUrls: ['./stock.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    TextColorDirective,
    CardComponent,
    CardHeaderComponent,
    CardBodyComponent,
    RowComponent,
    ColComponent,
    DropdownComponent,
    DropdownToggleDirective,
    DropdownMenuDirective,
    DropdownItemDirective,
    ButtonDirective
  ],
  changeDetection: ChangeDetectionStrategy.Default
})
export class StockComponent implements OnInit {
  stockId: string | null = null;
  stocks: any = null;
  wallet: any = null;
  isDarkTheme$: Observable<boolean>;

  constructor(
    private route: ActivatedRoute,
    private stockService: StockService,
    private themeService: ThemeService
  ) {
    this.isDarkTheme$ = this.themeService.isDarkTheme$;
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.stockId = params.get('id');
      if (this.stockId) {
        this.stockService.getStocks(this.stockId).subscribe({
          next: (wallet) => {
            this.wallet = wallet;
            this.stocks = wallet.assets;
          },
          error: (err) => {
            console.error('Error fetching stock:', err);
          }
        });
      }
    });
  }

  excluirStock(stock: any): void {
    console.log('Excluindo stock:', stock);
  }
}