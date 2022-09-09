package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var repeat int
	fmt.Fscan(in, &repeat)

	pair_programming(repeat, in, out)
}
