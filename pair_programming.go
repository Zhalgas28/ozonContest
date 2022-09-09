package main

import (
	"bufio"
	"fmt"
)

func pair_programming(repeat int, in *bufio.Reader, out *bufio.Writer) {
	for r := 0; r < repeat; r++ {
		// 2 <= devsCnt <= 100
		var devsCnt int
		fmt.Fscan(in, &devsCnt)

		// levels of developers
		levels := make([]int, devsCnt)

		var nums []int

		for d := 0; d < devsCnt; d++ {
			fmt.Fscan(in, &levels[d])
		}

		for i, _ := range levels {
			nums = append(nums, i+1)
		}

		for len(levels) >= 2 {
			idx1, idx2 := findMinDiff(levels)

			fmt.Println(nums[idx1], nums[idx2])

			levels = append(levels[:idx1], levels[idx1+1:]...)
			levels = append(levels[:idx2-1], levels[idx2:]...)

			nums = append(nums[:idx1], nums[idx1+1:]...)
			nums = append(nums[:idx2-1], nums[idx2:]...)
		}
		fmt.Print("\n")
	}
}

func findMinDiff(list []int) (idx1, idx2 int) {
	min := 100 // because max skill level can't be larger than 100
	for j := 1; j < len(list); j++ {
		if list[0] == list[j] {
			idx1, idx2 = 0, j
			return
		}

		if list[0] > list[j] && list[0]-list[j] < min {
			min = list[0] - list[j]
			idx1, idx2 = 0, j
		} else if list[0] < list[j] && list[j]-list[0] < min {
			min = list[j] - list[0]
			idx1, idx2 = 0, j
		}
	}
	return
}
