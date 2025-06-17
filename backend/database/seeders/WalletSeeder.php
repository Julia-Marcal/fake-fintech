<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Wallet;

class WalletSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Wallet::factory()
            ->count(2)
            ->create([
                'user_id' => 11,
            ]);

        Wallet::factory()
            ->count(10)
            ->create([
                'user_id' => fn () => rand(1, 10),
            ]);
    }
}