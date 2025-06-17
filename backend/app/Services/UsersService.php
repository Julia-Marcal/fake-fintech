<?php

namespace App\Services;

use App\Http\Controllers\UsersController;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Contracts\Pagination\LengthAwarePaginator;
use Illuminate\Http\JsonResponse;

class UsersService
{
    protected UsersController $controller;

    public function __construct(UsersController $controller)
    {
        $this->controller = $controller;
    }

    public function getAllUsers(Request $request): LengthAwarePaginator
    {
        $pageSize = $request->input('pageSize', 10);

        $users = User::paginate($pageSize);

        return $users;
    }

    public function getUser(string $id)
    {
        $userResource = $this->controller->getUserById($id);

        if (is_null($userResource->resource)) {
            return response()->json(['msg' => 'User not found', 'error' => true], 404);
        }

        return $userResource;
    }

    public function getUserWallets($id)
    {
        $userWallets = $this->controller->getUserWallets($id);

        return $userWallets;
    }

    public function deleteUser(string $id): JsonResponse
    {
        $user = $this->controller->deleteUser($id);

        return $user;
    }

}
