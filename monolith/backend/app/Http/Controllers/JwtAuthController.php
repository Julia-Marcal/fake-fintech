<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use Tymon\JWTAuth\Facades\JWTAuth;
use Tymon\JWTAuth\Exceptions\JWTException;
use Illuminate\Support\Facades\Auth;

class JwtAuthController extends Controller
{
    public function register(array $body)
    {
        $user = User::create([
            'name' => $body['name'],
            'last_name' => $body['last_name'],
            'age' => $body['age'],
            'email' => $body['email'],
            'password' => Hash::make($body['password']),
        ]);

        $token = JWTAuth::fromUser($user);

        return response()->json(compact('user', 'token'), 201);
    }

    public function login(array $credentials)
    {

        try {
            if (! $token = JWTAuth::attempt($credentials)) {
                return response()->json(['error' => 'Invalid credentials'], 401);
            }

            $user = Auth::user();

            $token = JWTAuth::claims(['role' => $user->role])->fromUser($user);

            return response()->json(compact('token'));
        } catch (JWTException $e) {
            return response()->json(['error' => 'Could not create token'], 500);
        }
    }

    public function logout()
    {
        JWTAuth::invalidate(JWTAuth::getToken());

        return response()->json(['message' => 'Successfully logged out']);
    }

    public function refresh()
    {
        try {
            $token = JWTAuth::getToken();
            if (!$token) {
                return response()->json(['error' => 'Token not provided'], 401);
            }
            $newToken = JWTAuth::refresh($token);
            return response()->json(['token' => $newToken]);
        } catch (JWTException $e) {
            return response()->json(['error' => 'Could not refresh token'], 401);
        }
    }

    public function validateToken($credentials)
    {
        try {
            $token = JWTAuth::attempt($credentials);
            if (!$token) {
                return response()->json(['error' => 'Invalid credentials'], 401);
            } else {
                $user = Auth::user();
                return response()->json(['msg' => 'Token is valid.', 'user' => $user, 'error' => false], 200);
            }
        } catch (JWTException $e) {
            return response()->json(['error' => 'Could not validate token'], 500);
        }
    }
}
