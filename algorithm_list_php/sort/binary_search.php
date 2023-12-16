<?php


function binarySearch($array, $target)
{
    $n = count($array);
    $left = 0;
    $right = $n -1;
    while ($left <= $right) {
        $mid = $left + ($right - $left) / 2;
        if ($array[$mid] == $target) {
            return  $array[$mid];
            }
        if ($array[$mid] > $target) {
            $right = $mid -1;
        } else {
            $left = $mid + 1;
        }
    }
    return -1;
}


$arr = [1,2,3,6,8,9];
$target = 8;
$result = binarySearch($arr,$target);
echo $result, "\n";