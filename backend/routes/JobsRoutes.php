<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MessageController;

Route::prefix('jobs')->middleware('throttle:api')->group(function () {
    Route::post('/send-message', [MessageController::class, 'sendMessage']);
});

Route::get('/clear-cache', function () {
    Artisan::call('cache:clear');
    Artisan::call('config:clear');
    Artisan::call('route:clear');
    Artisan::call('view:clear');

    return response()->json(['message' => 'Cache cleared successfully']);
});
