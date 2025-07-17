import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';
import { CommonModule } from '@angular/common';
import { StockService } from './stock.service';
import { ChangeDetectionStrategy } from '@angular/core';
import { WidgetsDropdownComponent } from '../../../app/base/widgets/widgets-dropdown/widgets-dropdown.component';
import { WidgetsBrandComponent } from '../../../app/base/widgets/widgets-brand/widgets-brand.component';
import { IconDirective } from '@coreui/icons-angular';
import { WidgetsEComponent } from '../../../app/base/widgets/widgets-e/widgets-e.component';
import { ThemeService } from '../../../shared/services/themes/themes.service';

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
    // TODO: Implement logic to delete a stock, perhaps with a confirmation dialog
    console.log('Excluindo stock:', stock);
    // Example: this.stockService.deleteStock(stock.id).subscribe(...)
  }
}