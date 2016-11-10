// duplate.go prints the test of each line that appears more than once in the standard input , preceded by ist count
package test_package

import (
	"bufio"
	"fmt"
	"os"
)

func Test_duplate() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
