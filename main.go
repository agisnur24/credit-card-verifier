package main

import (
	"bufio"
	"fmt"
	"luhns-algorithm/source"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		card := scanner.Text()
		fmt.Println(source.CardVerifier(card))
	}
	// 4556 7375 8689 9855
}
