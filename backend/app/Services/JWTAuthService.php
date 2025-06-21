<?php

namespace App\Services;

use App\Http\Controllers\JwtAuthController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Http\JsonResponse;

class JWTAuthService
{
    protected JwtAuthController $controller;

    public function __construct(JwtAuthController $controller)
    {
        $this->controller = $controller;
    }

    public function postUser(Request $request): JsonResponse
    {
        $validator = Validator::make($request->all(), [
            'name' => ['required', 'string', 'max:255'],
            'last_name' => ['required', 'string', 'max:255'],
            'age' => ['required', 'integer'],
            'email' => ['required', 'string', 'max:255', 'email', 'unique:users'],
            'password' => ['required', 'string', 'min:6', 'confirmed'],
        ]);

        if ($validator->fails()) {
            return response()->json([
                'msg' => 'Validation failed',
                'errors' => $validator->errors()->toJson(),
                'error' => true
            ], 422);
        }

        $validated = $validator->validated();

        $user = $this->controller->register($validated);

        if (!$user) {
            return response()->json(['error' => 'User creation failed.'], 422);
        }

        return response()->json(
            ['msg' => 'User created successfully.', 'error' => false],
            200
        );
    }

    public function login(Request $request): JsonResponse
    {
        $validator = Validator::make($request->all(), [
            'email' => ['required', 'email'],
            'password' => ['required'],
        ]);

        if ($validator->fails()) {
            return response()->json([
                'msg' => 'Validation failed',
                'errors' => $validator->errors(),
                'error' => true
            ], 422);
        }

        $validated = $validator->validated();

        $response = $this->controller->login($validated);

        return $response;
    }

    public function logout(): JsonResponse
    {
        $response = $this->controller->logout();

        return $response;
    }

    public function refresh(): JsonResponse
    {

        $response = $this->controller->refresh();

        return $response;
    }
}
