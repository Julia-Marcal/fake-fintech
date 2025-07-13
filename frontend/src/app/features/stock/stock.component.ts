import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';
import { CommonModule } from '@angular/common';
import { StockService } from './stock.service';

@Component({
  selector: 'app-stock',
  templateUrl: './stock.component.html',
  styleUrls: ['./stock.component.scss'],
  standalone: true,
  imports: [CommonModule]
})
export class StockComponent implements OnInit {
  stockId$!: Observable<string | null>;
  stockId: string | null = null;
  stock: any = null;

  constructor(private route: ActivatedRoute, private stockService: StockService) { }
  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.stockId = params.get('id');
      if (this.stockId) {
        this.stockService.getStocks(this.stockId).subscribe({
          next: (stock) => {
            console.log(stock);

            this.stock = stock;
            console.log(this.stock);
          },
          error: (err) => {
            console.error('Error fetching stock:', err);
          }
        });
      }
    });
  }

}


