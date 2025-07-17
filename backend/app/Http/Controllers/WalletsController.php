<?php

namespace App\Http\Controllers;

use App\Models\Wallet;
use App\Models\WalletAssets;
use App\Models\Assets;
use App\Http\Resources\WalletsResource;


class WalletsController extends Controller
{

    public function getWallets($page = 1, $perPage = 10)
    {
        $wallets = Wallet::orderBy('name')->paginate($perPage, ['*'], 'page', $page);
        return WalletsResource::collection($wallets);
    }

    public function getWalletById($id)
    {
        $wallet = Wallet::find($id);

        if (!$wallet) {
            return response()->json(['message' => 'Wallet not found'], 404);
        }

        $wallet_assets = WalletAssets::where('wallet_id', $id)->get();
        $assets = Assets::whereIn('id', $wallet_assets->pluck('asset_id'))->get();

        $wallet->assets = $assets;

        return new WalletsResource(resource: $wallet);
    }
}
