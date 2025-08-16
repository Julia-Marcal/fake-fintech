# How to Test RabbitMQ

This tutorial explains how to test RabbitMQ by publishing a message to a queue and running the consumer service.

## 1. Publish a Message to RabbitMQ

Use the following command to publish a message to the `assets-watcher` queue:

```bash
php artisan rabbitmq:publish assets-watcher '{"action": "FETCH_PRICE", "market": "USD", "symbol": "BTC"}'
```

## 2. Run the Consumer Service

Navigate to the consumer service directory:

```bash
cd ../fake-fintech/services/assets-watcher/cmd
```

Then, start the service using Go:

```bash
go run main.go
```

## 3. Verify Message Processing

Check the output of the consumer service to confirm that the message was received and processed.

**Note:** Ensure RabbitMQ is running and properly configured before testing.
