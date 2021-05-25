<?php
$arr = [];
function fbnq($n)
{
    global $arr;
    if ($n <= 0) {
        $arr[$n] = 0;
    } else if ($n == 1 || $n == 2) {
        $arr[$n] = 1;
    } else {
        $arr[$n] = sprintf("%.0f", $arr[$n - 1] + $arr[$n - 2]);
    }
    return $arr[$n];
}

function printFbnq($n)
{
    for ($i = 0; $i < $n; $i++) {
        $result = fbnq($i);
        $str = sprintf("数列第%d位：%s\n", $i + 1, number_format($result, 0, '', ''));
        echo $str;
    }
}

$start = microtime(true);
printFbnq(1000);
$end = microtime(true);
var_dump('运行时间为：' . ($end - $start));
