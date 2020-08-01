package main

import (
	"fmt"
	"os"

	"github.com/alam0rt/ssm/pkg/ssm"
)

func main() {
	fmt.Printf("%s\n", "hello")
	path := os.Args[1]
	p, err := ssm.GetParameters(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	for _, param := range p {
		test := ssm.Parameter(param).OutputToInput()
	}
	fmt.Print(p)
}
