<?php

namespace App\Services;

use App\Http\Controllers\WalletsController;
use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\AnonymousResourceCollection;

class WalletsService
{
    protected WalletsController $controller;

    public function __construct(WalletsController $controller)
    {
        $this->controller = $controller;
    }

    public function getAllWallets(Request $request): AnonymousResourceCollection
    {
        $pageSize = $request->input(key: 'pageSize', default: 10);
        $page = $request->input(key: 'page', default: 1);

        $walletResource = $this->controller->getWallets(page: $page, perPage: $pageSize);

        return $walletResource;
    }

    public function getWalletById(string $id)
    {
        $walletResource = $this->controller->getWalletById($id);

        if (is_null(value: $walletResource->resource)) {
            return response()->json(data: ['msg' => 'Wallet not found', 'error' => true], status: 404);
        }

        return $walletResource;
    }
}
