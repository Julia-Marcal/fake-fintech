<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class WalletAssets extends Model
{
    use HasFactory;

    protected $table = 'wallet_assets';

    protected $fillable = [
        'wallet_id',
        'asset_id',
    ];

    /**
     * Get the wallet that owns this asset.
     */
    public function wallet()
    {
        return $this->belongsTo(Wallet::class);
    }

    /**
     * Get the asset that belongs to this wallet.
     */
    public function asset()
    {
        return $this->belongsTo(Asset::class);
    }

    public function walletAssets()
    {
        return $this->hasMany(WalletAsset::class);
    }

    public function assets()
    {
        return $this->belongsToMany(Asset::class, 'wallet_assets');
    }
}