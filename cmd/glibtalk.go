package main

import (
	"flag"
	"fmt"
	"glibness/glibness"
	"os"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: ./glibtalk [filename]")
		return
	}

	path := args[0]
	file, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	contents := string(file)
	dialogues, err := glibness.Parse(contents)

	if err != nil {
		fmt.Println(err)
		return
	}

	engine := glibness.Engine{Dialogues: dialogues}
	state, err := engine.Start("main")

	if err != nil {
		fmt.Println(err)
		return
	}

	for !state.Finished {
		status, err := state.Next()

		if err != nil {
			fmt.Println(err)
			return
		}

		if status == glibness.DialogueChangeSay {
			fmt.Printf("[%s]: %s\n", state.Speaker, state.Sentence)
			// TODO find something better
			time.Sleep(time.Second)
		}
	}
}
