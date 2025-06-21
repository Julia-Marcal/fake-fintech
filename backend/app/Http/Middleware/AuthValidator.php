<?php

namespace App\Http\Middleware;

use Closure;
use Tymon\JWTAuth\Facades\JWTAuth;
use Tymon\JWTAuth\Exceptions\JWTException;
use Illuminate\Support\Facades\Log;

class AuthValidator
{
    public function handle($request, Closure $next)
    {
        Log::info('AuthValidator Middleware triggered', [
            'url' => $request->fullUrl(),
            'method' => $request->method(),
            'ip' => $request->ip(),
            'headers' => $request->headers->all(),
        ]);

        try {
            $user = JWTAuth::parseToken()->authenticate();
        } catch (JWTException $e) {
            Log::warning('JWTException in AuthValidator', [
                'message' => $e->getMessage(),
            ]);
            return response()->json(['error' => 'Token not valid'], 401);
        }

        return $next($request);
    }
}
