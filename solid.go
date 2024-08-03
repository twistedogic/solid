package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/twistedogic/solid/internal"
)

const (
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	reset  = "\033[0m"
)

func main() {
	m, err := internal.DefaultModel()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("missing file path argument")
	}
	f, err := internal.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	review, err := internal.ReviewCode(context.Background(), m, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(review)
	if mi, err := internal.MI(context.Background(), m, f); err == nil {
		fmt.Printf("Maintainability Index for %q:\n", f.Name)
		switch {
		case mi < 9.0:
			fmt.Println(red, mi, reset)
		case mi < 20:
			fmt.Println(yellow, mi, reset)
		default:
			fmt.Println(green, mi, reset)
		}
		return
	}
	mi, err := internal.MaintainabilityIndex(context.Background(), m, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mi)
}
