<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration {
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create(table: 'assets_requests', callback: function (Blueprint $table): void {
            $table->id();
            $table->foreignId(column: 'wallet_assets_id');
            $table->string(column: 'link');
            $table->json(column: 'payload');
            $table->timestamps();

            $table->foreign(columns: 'wallet_assets_id')->references(columns: 'id')->on(table: 'wallet_assets')->onDelete(action: 'cascade');

            $table->unique(columns: ['wallet_assets_id']);
        });
    }
    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists(table: 'assets_requests');
    }
};