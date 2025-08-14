<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Services\WalletsService;

Route::prefix('wallets')->middleware('throttle:api')->group(function () {
    Route::middleware(['auth:api'])->group(function () {
        Route::get('/', function (Request $request, WalletsService $WalletsService) {
            return $WalletsService->getAllWallets($request);
        });
        Route::get('/{id}', function (string $id, WalletsService $WalletsService) {
            return $WalletsService->getWalletById($id);
        });
    });
});
