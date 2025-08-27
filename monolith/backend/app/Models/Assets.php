<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Str;

class Assets extends Model
{
     use HasFactory;

    protected $table = 'assets';

    protected $fillable = [
        'name',
        'ticker',
        'type',
        'price',
        'quantity',
        'currency',
        'sector',
        'dividend_yield',
        'expense_ratio',
        'maturity_date',
        'coupon_rate',
    ];

    protected $casts = [
        'price' => 'decimal:4',
        'quantity' => 'decimal:4',
        'dividend_yield' => 'decimal:4',
        'expense_ratio' => 'decimal:4',
        'coupon_rate' => 'decimal:4',
        'maturity_date' => 'date',
        'created_at' => 'datetime',
        'updated_at' => 'datetime',
    ];

    protected static function boot()
    {
        parent::boot();

        static::creating(function ($model) {
            $model->id = Str::uuid()->toString();
        });
    }

    public function walletAssets()
    {
        return $this->hasMany(WalletAsset::class);
    }
}
