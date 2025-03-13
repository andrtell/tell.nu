//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// dep
	deps := read_deps()
	fmt.Printf("%v\n", deps)
}

func read_deps() [][]string {
	file, err := os.Open("dat/dep.txt")
	if err != nil {
		panic(err)	
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	deps := make([][]string, 0)
	for scan.Scan() {
		line := scan.Text()
		part := strings.Split(line, " ")
		deps = append(deps, part)
	}
	return deps
}
