<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Jobs\ProcessRabbitMQMessage as ProcessRabbitMQMessageService;


class ProcessRabbitMQMessage extends Command
{
    protected $signature = 'rabbitmq:consume';
    protected $description = 'Consume messages from RabbitMQ';

    public function handle()
    {
        $rabbit = new ProcessRabbitMQMessageService();
        $rabbit->consume('test_queue', function ($msg) {
            $this->info("Received: {$msg->body}");
        });
    }
}
