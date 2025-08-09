<?php

namespace App\Jobs;

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

class ProcessRabbitMQMessage
{
    private $connection;
    private $channel;

    public function __construct()
    {
        $this->connection = new AMQPStreamConnection(
            env('RABBITMQ_HOST'),
            env('RABBITMQ_PORT'),
            env('RABBITMQ_USER'),
            env('RABBITMQ_PASSWORD'),
            env('RABBITMQ_VHOST')
        );

        $this->channel = $this->connection->channel();
    }

    public function publish(string $queue, string $message)
    {
        $this->channel->queue_declare($queue, false, true, false, false);
        $msg = new AMQPMessage($message);
        $this->channel->basic_publish($msg, '', $queue);
    }

    public function get(string $queue): ?string
    {
        $this->channel->queue_declare($queue, false, true, false, false);
        $message = $this->channel->basic_get($queue);

        if (null === $message) {
            return null;
        }

        // Acknowledge the message so RabbitMQ knows it can be deleted.
        $this->channel->basic_ack($message->getDeliveryTag());

        return $message->getBody();
    }

    public function consume(string $queue, callable $callback)
    {
        $this->channel->queue_declare($queue, false, true, false, false);
        $this->channel->basic_consume($queue, '', false, true, false, false, $callback);

        while ($this->channel->is_consuming()) {
            $this->channel->wait();
        }
    }

    public function __destruct()
    {
        $this->channel->close();
        $this->connection->close();
    }
}
