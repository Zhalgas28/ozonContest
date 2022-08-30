package main

import (
	"bufio"
	"fmt"
)

func SumToPay(repeat int, in *bufio.Reader, out *bufio.Writer) {
	for i := 0; i < repeat; i++ {
		goods := make(map[int]int)
		var gdsCount int
		fmt.Fscan(in, &gdsCount)
		for j := 0; j < gdsCount; j++ {
			var tmp int
			fmt.Fscan(in, &tmp)
			goods[tmp]++
		}
		var result int
		for price, count := range goods {
			result += (count - (count / 3)) * price
		}
		fmt.Fprintln(out, result)
	}
}
