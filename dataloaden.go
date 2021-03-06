package main

import (
	"fmt"
	"os"

	"github.com/cybernik/dataloaden/pkg/generator"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("usage: name keyType valueType")
		fmt.Println(" example:")
		fmt.Println(" dataloaden 'UserLoader int []*github.com/my/package.User user_loader.go'")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if err := generator.Generate(os.Args[1], os.Args[2], os.Args[3], os.Args[4], wd); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
