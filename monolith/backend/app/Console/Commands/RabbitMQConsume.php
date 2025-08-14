<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Services\RabbitMQService;

class RabbitMQConsume extends Command
{
    protected $signature = 'rabbitmq:consume';
    protected $description = 'Consume messages from RabbitMQ';

    public function handle()
    {
        $rabbit = new RabbitMQService();
        $rabbit->consume('test_queue', function ($msg) {
            $this->info("Received: " . $msg->body);
        });
    }
}
