<?php

namespace Tests\Unit;

use Tests\TestCase;

class UserTest extends TestCase
{
    /**
     * A basic unit test example.
     */
    public function test_example(): void
    {
        $this->assertTrue(true);
    }

    public function test_user_can_be_created_with_fillable_attributes(): void
    {
        $user = new \App\Models\User([
            'name' => 'John',
            'last_name' => 'Doe',
            'age' => 30,
            'email' => 'john@example.com',
            'password' => 'secret',
        ]);
        $this->assertEquals('John', $user->name);
        $this->assertEquals('Doe', $user->last_name);
        $this->assertEquals(30, $user->age);
        $this->assertEquals('john@example.com', $user->email);
    }

    public function test_hidden_attributes_are_not_visible_in_array(): void
    {
        $user = new \App\Models\User([
            'password' => 'secret',
            'remember_token' => 'token',
        ]);
        $array = $user->toArray();
        $this->assertArrayNotHasKey('password', $array);
        $this->assertArrayNotHasKey('remember_token', $array);
    }

    public function test_jwt_identifier_and_claims(): void
    {
        $user = new \App\Models\User();
        $user->id = 123;
        $this->assertEquals(123, $user->getJWTIdentifier());
        $this->assertEquals([], $user->getJWTCustomClaims());
    }
}
