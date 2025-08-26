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

        if ($message) {
            return response()->json([
                'status' => 'Message consumed',
                'message' => $message
            ]);
        }

        return response()->json(['status' => 'No messages in queue']);
    }
}