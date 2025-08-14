<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use App\Models\AssetsRequests;

class AssetsRequestsSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        AssetsRequests::create([
            'id' => 'a1b2c3d4-e5f6-7890-1234-567890abcdef',
            'assets_id' => 'ef06d611-e2a1-41b1-9321-3fb0e507b7f4',
            'link' => 'https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=BTC&market=USD&apikey=WX618VL5L0J07ON3',
            'created_at' => now(),
            'updated_at' => now(),
        ]);
    }
}
