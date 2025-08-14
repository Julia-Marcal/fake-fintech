<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use App\Models\Wallet;
use App\Models\Assets;
use App\Models\WalletAssets; // Note: Changed from WalletAssets to WalletAsset

class WalletAssetsSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        // Get wallets and assets
        $mainPortfolio = Wallet::where('name', 'Main Portfolio')->first();
        $retirementFund = Wallet::where('name', 'Retirement Fund')->first();
        $tradingAccount = Wallet::where('name', 'Trading Account')->first();

        $aapl = Assets::where('ticker', 'AAPL')->first();
        $msft = Assets::where('ticker', 'MSFT')->first();
        $spy = Assets::where('ticker', 'SPY')->first();
        $ust = Assets::where('ticker', 'UST2030')->first();
        $btc = Assets::where('ticker', 'BTC')->first();

        // Assign assets to wallets
        WalletAssets::create([
            'wallet_id' => $mainPortfolio->id,
            'asset_id' => $aapl->id,
        ]);

        WalletAssets::create([
            'wallet_id' => $mainPortfolio->id,
            'asset_id' => $msft->id,
        ]);

        WalletAssets::create([
            'wallet_id' => $retirementFund->id,
            'asset_id' => $spy->id,
        ]);

        WalletAssets::create([
            'wallet_id' => $retirementFund->id,
            'asset_id' => $ust->id,
        ]);

        WalletAssets::create([
            'wallet_id' => $tradingAccount->id,
            'asset_id' => $btc->id,
        ]);

        WalletAssets::create([
            'wallet_id' => $tradingAccount->id,
            'asset_id' => $aapl->id,
        ]);
    }
}