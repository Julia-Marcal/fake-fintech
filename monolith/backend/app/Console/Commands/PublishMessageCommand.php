<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Jobs\ProcessRabbitMQMessage;

class PublishMessageCommand extends Command
{
    /**
     * The name of the console command.
     *
     * @var string
     */
    protected $name = 'rabbitmq:publish';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Publish a message to a RabbitMQ queue';

    /**
     * Execute the console command.
     */
    public function handle(): void
    {
        $queue = $this->argument('queue');
        $message = $this->argument('message');

        $rabbitMQ = new ProcessRabbitMQMessage();
        $rabbitMQ->publish($queue, $message);

        $this->info("Message published to '{$queue}' queue.");
    }

    /**
     * Get the console command arguments.
     *
     * @return array
     */
    protected function getArguments()
    {
        return [
            ['queue', \Symfony\Component\Console\Input\InputArgument::REQUIRED, 'The name of the queue.'],
            ['message', \Symfony\Component\Console\Input\InputArgument::REQUIRED, 'The message to publish.'],
        ];
    }
}