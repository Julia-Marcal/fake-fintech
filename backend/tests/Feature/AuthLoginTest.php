<?php

namespace Tests\Feature;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;
use App\Models\User;
use Illuminate\Support\Facades\Hash;

class AuthLoginTest extends TestCase
{
    use RefreshDatabase;

    public function test_login_with_valid_credentials_returns_token()
    {
        User::create([
            'name' => 'Test',
            'last_name' => 'User',
            'age' => 25,
            'email' => 'teste_user@example.com',
            'password' => Hash::make('senha123'),
        ]);

        $response = $this->postJson('api/auth/login', [
            'email' => 'teste_user@example.com',
            'password' => 'senha123',
        ]);

        $response->assertStatus(200);
        $response->assertJsonStructure([
            'token',
        ]);
    }
}
