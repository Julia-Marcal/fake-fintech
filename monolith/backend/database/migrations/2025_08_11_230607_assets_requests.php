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
            $table->uuid(column: 'id')->primary();
            $table->foreignUuid(column: 'assets_id')->constrained(table: 'assets')->onDelete(action: 'cascade');
            $table->string(column: 'link');
            $table->json(column: 'payload')->nullable();
            $table->timestamps();

            $table->foreign(columns: 'assets_id')->references(columns: 'id')->on(table: 'assets')->onDelete(action: 'cascade');

            $table->unique(columns: ['assets_id']);
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