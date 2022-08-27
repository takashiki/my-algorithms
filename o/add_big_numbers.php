<?php

// $a = '123';
// $b = '12323234938048';
// $a = '6932828238829';
// $b = '7238384282733';

$a = '999';
$b = '11';
echo addBigNum($a, $b);

function addBigNum(string $a, string $b): string
{
    $aLen = strlen($a);
    $bLen = strlen($b);
    $carry = 0;
    $res = '';
    for ($i = 0; $i < $aLen || $i < $bLen; $i++) {
        // iterate $a from right to left
        $ar = 0;
        if ($i < $aLen) {
            $ar = $a[$aLen - $i - 1];
        }
        // iterate $b from right to left
        $br = 0;
        if ($i < $bLen) {
            $br = $b[$bLen - $i - 1];
        }

        // result in current bit is sum of $a and $b and $carry
        $t = $ar + $br + $carry;

        // if $t >= 10, carry is 1, and result of this bit is $t-10
        if ($t >= 10) {
            $carry = 1;
            $t -= 10;
        } else {
            // if $t < 10, carray is 0
            $carry = 0;
        }

        $res = $t . $res;
    }

    // after all bit are iterated, if carry=1 we need to add 1 in front of result
    if ($carry > 0) {
        $res = $carry . $res;
    }

    return $res;
}
