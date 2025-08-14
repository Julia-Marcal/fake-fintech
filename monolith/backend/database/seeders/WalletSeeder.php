<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Wallet;
use Illuminate\Support\Str;

class WalletSeeder extends Seeder
{
    public function run(): void
    {
        Wallet::create([
            'id' => Str::uuid()->toString(),
            'user_id' => 11,
            'name' => 'Main Portfolio',
            'description' => 'Primary investment wallet',
            'balance' => 10000.00,
            'currency' => 'USD',
            'is_active' => true,
        ]);

        Wallet::create([
            'id' => Str::uuid()->toString(),
            'user_id' => rand(1, 11),
            'name' => 'Retirement Fund',
            'description' => 'Long-term retirement investments',
            'balance' => 50000.00,
            'currency' => 'USD',
            'is_active' => true,
        ]);

        Wallet::create([
            'id' => Str::uuid()->toString(),
            'user_id' => rand(1, 11),
            'name' => 'Trading Account',
            'description' => 'Active trading wallet',
            'balance' => 15000.00,
            'currency' => 'USD',
            'is_active' => true,
        ]);
    }
}