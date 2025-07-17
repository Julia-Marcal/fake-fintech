import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { ConfigService } from '../../../shared/services/config/config.service';

interface Wallet {
    id: string;
    name: string;
    description: string;
    balance: string;
    currency: string;
    created_at: string;
    updated_at: string;
    assets: Array<{
        id: string;
        name: string;
        ticker: string;
        sector: string;
        type: string;
        price: string;
        quantity: string;
        currency: string;
        dividend_yield: string;
        expense_ratio: string | null;
        coupon_rate: string | null;
        maturity_date: string | null;
        total_value?: number;
        created_at: string;
        updated_at: string;
    }>;
}

interface ApiResponse {
    data: Wallet;
}

@Injectable({
    providedIn: 'root'
})
export class StockService {
    private readonly url: string;
    private readonly token: string;

    constructor(private http: HttpClient, private configService: ConfigService) {
        this.url = this.configService.apiBaseUrl;
        this.token = this.configService.apiToken;
    }

    getStocks(id: string): Observable<Wallet> {
        return this.http.get<ApiResponse>(`${this.url}/wallets/${id}`, {
            headers: {
                Authorization: `Bearer ${this.token}`
            }
        }).pipe(
            map(response => response.data)
        );
    }
}