<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use App\Models\Assets;

class AssetsSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
         // Stocks
        Assets::create([
            'name' => 'Apple Inc.',
            'ticker' => 'AAPL',
            'type' => 'stock',
            'price' => 175.32,
            'quantity' => 100,
            'currency' => 'USD',
            'sector' => 'Technology',
            'dividend_yield' => 0.0054,
        ]);

        Assets::create([
            'name' => 'Microsoft Corporation',
            'ticker' => 'MSFT',
            'type' => 'stock',
            'price' => 328.39,
            'quantity' => 50,
            'currency' => 'USD',
            'sector' => 'Technology',
            'dividend_yield' => 0.0075,
        ]);

        // ETFs
        Assets::create([
            'name' => 'SPDR S&P 500 ETF',
            'ticker' => 'SPY',
            'type' => 'etf',
            'price' => 445.78,
            'quantity' => 25,
            'currency' => 'USD',
            'sector' => 'Diversified',
            'dividend_yield' => 0.0123,
            'expense_ratio' => 0.0009,
        ]);

        // Bonds
        Assets::create([
            'name' => 'US Treasury Bond 2030',
            'ticker' => 'UST2030',
            'type' => 'bond',
            'price' => 98.75,
            'quantity' => 1000,
            'currency' => 'USD',
            'sector' => 'Government',
            'coupon_rate' => 0.0375,
            'maturity_date' => '2030-05-15',
        ]);

        // Crypto
        Assets::create([
            'name' => 'Bitcoin',
            'ticker' => 'BTC',
            'type' => 'crypto',
            'price' => 42000.00,
            'quantity' => 0.5,
            'currency' => 'USD',
            'sector' => 'Cryptocurrency',
        ]);
    }
}
