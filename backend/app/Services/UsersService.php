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

    public function getUserWallets($id): JsonResponse
    {
        $userWallets = $this->controller->getUserWallets($id);

        if($userWallets->isEmpty()) {
            return response()->json(['msg' => 'User not found or has no wallets', 'error' => true], 404);
        }

        return response()->json($userWallets->original);
    }

    public function deleteUser(string $id): JsonResponse
    {
        $user = $this->controller->deleteUser($id);

        return $user;
    }

    public function updateUser(string $id, Request $request): JsonResponse
    {
        $user = $this->controller->updateUser($id, $request->all());

        if(!$user) {
            return response()->json(['msg' => 'User not found', 'error' => true], 404);
        }

        return response()->json($user);
    }

}
