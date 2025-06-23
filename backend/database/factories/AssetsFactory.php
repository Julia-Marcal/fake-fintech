<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Assets>
 */
class AssetsFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'name' => $this->faker->company,
            'ticker' => strtoupper($this->faker->lexify('???')),
            'type' => $this->faker->randomElement(['stock', 'bond', 'etf']),
            'price' => $this->faker->randomFloat(2, 10, 1000),
            'quantity' => $this->faker->randomFloat(4, 1, 100),
            'currency' => $this->faker->currencyCode,
            'sector' => $this->faker->word,
            'dividend_yield' => $this->faker->randomFloat(4, 0, 0.1),
            'expense_ratio' => $this->faker->randomFloat(4, 0, 0.05),
            'maturity_date' => $this->faker->optional()->date(),
            'coupon_rate' => $this->faker->optional()->randomFloat(4, 0, 0.1),
        ];
    }
}
