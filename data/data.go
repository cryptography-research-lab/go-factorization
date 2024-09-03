package data

import _ "embed"

// PrimeNumberTable 打包时内嵌的素数表，把前几万个素数都放进来加速因式分解的过程
//
//go:embed prime-number-table.txt
var PrimeNumberTable string
