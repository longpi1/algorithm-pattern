<?php
$txt1 = "Hello";
$txt2 = "World";
echo $txt1 . " " . $txt2; // 输出 "Hello World"


echo strlen("Hello world!"); // 输出 12


echo strpos("Hello world!", "world"); // 输出 6


echo str_replace("world", "Dolly", "Hello world!"); // 输出 "Hello Dolly!"


echo strtolower("Hello WORLD!"); // 输出 "hello world!"
echo strtoupper("Hello WORLD!"); // 输出 "HELLO WORLD!"


print_r(explode(" ", "Hello world!")); // 输出 Array ( [0] => Hello [1] => world! )


$array = array('Hello', 'World!');
echo implode(" ", $array); // 输出 "Hello World!"


echo trim(" Hello world! "); // 输出 "Hello world!"
