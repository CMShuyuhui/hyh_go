// duplate.go prints the test of each line that appears more than once in the standard input.
// It read from stdin or from a list of named files
package test_package

import (
	"bufio"
	"fmt"
	"os"
)

func Test_duplate2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplate2:  %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for str, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, str)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		counts[input.Text()]++
	}
}
