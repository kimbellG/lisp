package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		str := sc.Text()
		fmt.Printf("Эхо: %s\n", str)
	}
}
