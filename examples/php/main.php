<?php
header ( "Content-Type: application/json;charset=utf-8" );
$url = 'http://localhost:8080/v1/cep/08226021';
$result = file_get_contents($url);
echo $result;
echo "\n\n";
$arr = json_decode($result, true);
print_r($arr);
