<?php
require_once __DIR__ . '/vendor/autoload.php';

$klein = new \Klein\Klein();

$klein->respond('GET', '/', function () {
  return 'Hello World! Hello Klein!';
});

$klein->dispatch();
