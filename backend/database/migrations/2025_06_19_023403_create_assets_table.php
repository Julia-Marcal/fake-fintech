<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
   /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('assets', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('name');
            $table->string('ticker')->unique();
            $table->string('type');
            $table->decimal('price', 15, 4)->nullable();
            $table->decimal('quantity', 15, 4)->nullable();
            $table->string('currency')->nullable();
            $table->string('sector')->nullable();
            $table->decimal('dividend_yield', 8, 4)->nullable();
            $table->decimal('expense_ratio', 8, 4)->nullable();
            $table->date('maturity_date')->nullable();
            $table->decimal('coupon_rate', 8, 4)->nullable();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('assets');
    }
};
