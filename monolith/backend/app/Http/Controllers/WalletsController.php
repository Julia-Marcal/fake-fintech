<?php

namespace App\Http\Controllers;

use App\Models\Wallet;
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
        return new WalletsResource($wallet);
    }
}
