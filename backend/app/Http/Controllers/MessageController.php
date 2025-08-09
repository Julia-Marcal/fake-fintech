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
            'message' => 'required|string',
        ]);

        $rabbit = new ProcessRabbitMQMessage();
        $rabbit->publish(queue: $data['queue'], message: $data['message']);

        return response()->json(['status' => 'Message queued for sending']);
    }
}