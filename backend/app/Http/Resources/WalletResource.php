<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class WalletResource extends JsonResource
{
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'description' => $this->description,
            'balance' => $this->balance,
            'currency' => $this->currencies,
            'created_at' => $this->created_at,
            'updated_at' => $this->updated_at,
        ];
    }
}
