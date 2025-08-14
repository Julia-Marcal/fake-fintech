<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

class WalletFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        $currencies = ['USD', 'EUR', 'BRL'];
        return [
            'id' => (string) Str::uuid(),
            'user_id' => $this->faker->numberBetween(1, 11),
            'name' => $this->faker->words(2, true),
            'description' => $this->faker->optional()->sentence(),
            'balance' => $this->faker->randomFloat(2, 0, 10000),
            'currency' => $this->faker->randomElement($currencies),
            'is_active' => $this->faker->boolean(90),
            'created_at' => now(),
            'updated_at' => now(),
        ];
    }
}