import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ConfigService } from '../../../shared/services/config/config.service';

export interface Wallet {
    name: string;
    description: string;
    balance: string;
    currency: string;
}

interface ApiResponse {
    user_id: string;
    wallets: Array<Wallet>;
}

@Injectable({
    providedIn: 'root'
})
export class WalletService {

    private readonly url: string;
    private readonly token: string;

    constructor(private http: HttpClient, private configService: ConfigService) {
        this.url = this.configService.apiBaseUrl;
        this.token = this.configService.apiToken;
    }

    getUserWallets(id: string): Observable<Wallet[]> {
        return this.http.get<ApiResponse>(`${this.url}/users/${id}/wallets`, {
            headers: {
                Authorization: `Bearer ${this.token}`
            }
        }).pipe(
            map((response: ApiResponse) => {
                return response.wallets
            })
        );
    }
}