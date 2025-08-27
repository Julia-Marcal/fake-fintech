<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Services\JWTAuthService;

Route::prefix('auth')->middleware('throttle:api')->group(function () {
    Route::post('/register', function (Request $request, JWTAuthService $JWTAuthService) {
        return $JWTAuthService->postUser($request);
    });
    Route::post('/login', function (Request $request, JWTAuthService $JWTAuthService) {
        return $JWTAuthService->login($request);
    });
    Route::post('/logout', function (Request $request, JWTAuthService $JWTAuthService) {
        return $JWTAuthService->logout($request);
    });
    Route::post('/refresh', function (Request $request, JWTAuthService $JWTAuthService) {
        return $JWTAuthService->refresh($request);
    });
});
