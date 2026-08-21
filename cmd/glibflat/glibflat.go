package main

import (
	"flag"
	"fmt"
	"glibness/glibness"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: ./glibtalk [filename]")
		return
	}

	path := args[0]
	engine := glibness.NewEngine()
	err := engine.ParseFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dialogue := range engine.Dialogues {
		fmt.Printf("%s:\n", dialogue.Name)

		for _, statement := range dialogue.Statements {
			indents := 4
			if _, ok := statement.(glibness.LabelStatement); ok {
				indents = 2
			}
			fmt.Printf("%s%s\n", strings.Repeat(" ", indents), statement.String())
		}
	}
}
