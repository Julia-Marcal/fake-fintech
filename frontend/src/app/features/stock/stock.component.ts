import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-stock',
  templateUrl: './stock.component.html',
  styleUrls: ['./stock.component.scss'],
  standalone: true,
  imports: [CommonModule]
})
export class StockComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
