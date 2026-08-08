package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "用法: levenshtein <串1> <串2>")
		os.Exit(2)
	}
	fmt.Println(Distance(os.Args[1], os.Args[2]))
}
