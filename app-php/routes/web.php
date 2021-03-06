<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

use Illuminate\Http\Request;

$router->get('/', function (Request $request) use ($router) {
    return [
        "status" => "running",
        "language" => "php"
    ];
});

$router->get('/hello', function (Request $request) use ($router) {
    return $request->all();
});
