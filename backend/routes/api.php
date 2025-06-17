<?php

require __DIR__ . '/UsersRoutes.php';
require __DIR__ . '/AuthRoutes.php';

use Illuminate\Support\Facades\RateLimiter;
use Illuminate\Cache\RateLimiting\Limit;
use Illuminate\Http\Request;

RateLimiter::for('api', function (Request $request) {
    return Limit::perMinute(60)->by($request->ip());
});
