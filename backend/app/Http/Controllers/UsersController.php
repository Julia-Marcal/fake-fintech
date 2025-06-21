<?php

namespace App\Http\Controllers;

use App\Models\User;
use App\Models\Wallet;
use Illuminate\Support\Facades\Redis;
use App\Http\Resources\UserResource;
use App\Http\Resources\WalletsResource;

class UsersController extends Controller
{
    public function getAllUsers($page = 1, $perPage = 10)
    {
        $users = User::orderBy('name')->paginate($perPage, ['*'], 'page', $page);
        return UserResource::collection($users);
    }

    public function getUserById($id)
    {
        $user = Redis::get('user:' . $id);

        if ($user) {
            $userArray = json_decode($user, true);
            $user = new User($userArray);
        } else {
            $user = User::find($id);

            if ($user)
                Redis::set('user:' . $id, json_encode($user));
        }

        if (!$user)
            return response()->json(['msg' => 'User not found', 'error' => true], 404);

        return new UserResource($user);
    }

    public function deleteUser($id)
    {
        $user = User::find($id);

        if (!$user) {
            return response()->json(['msg' => 'User not found', 'error' => true], 404);
        }

        $user->delete();
        return response()->json(['msg' => 'User deleted successfully', 'error' => false], 200);
    }

    public function getUserWallets($id)
    {
        $wallets = Wallet::where('user_id', $id)->get();

        if ($wallets->isEmpty()) {
            return response()->json(['msg' => 'User not found or has no wallets', 'error' => true], 404);
        }

        return response()->json([
            'user_id' => $id,
            'wallets' => WalletsResource::collection($wallets),
        ]);
    }
}
