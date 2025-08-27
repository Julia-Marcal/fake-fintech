<?php

namespace App\Http\Controllers;

use App\Jobs\ProcessRabbitMQMessage;
use Illuminate\Http\Request;

class MessageController extends Controller
{
    public function sendMessage(Request $request)
    {
        $data = $request->validate([
            'queue' => 'required|string',
            'message' => 'required|array',
        ]);

        $rabbit = new ProcessRabbitMQMessage();
        $rabbit->publish(queue: $data['queue'], message: json_encode($data['message']));

        return response()->json(['status' => 'Message queued for sending']);
    }
    public function consumeSingleMessage(Request $request)
    {
        $data = $request->validate([
            'queue' => 'required|string',
        ]);

        $rabbit = new ProcessRabbitMQMessage();
        $message = $rabbit->get($data['queue']);

        $decoded_message = $this->decodeMessage(message: $message);

        if ($message) {
            return response()->json(data: [
                'status' => 'Message consumed',
                'message' => $decoded_message->priceUsd
            ]);
        }

        return response()->json(data: ['status' => 'No messages in queue']);
    }

    function decodeMessage($message)
    {
        while (!is_object(value: $message)) {
            $message = json_decode(json: $message);
            $message = $message->response ?? $message->data ?? $message;
        }
        return $message;
    }
}