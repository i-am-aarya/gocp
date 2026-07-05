package main

import (
	"fastcopy/internal/tui"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) >= 2 {
		src := args[0]
		dst := args[1]

		fmt.Println("src: ", src, " dst: ", dst)

	} else {
		p := tui.NewTui()
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v", err)
			os.Exit(1)
		}
	}
}
