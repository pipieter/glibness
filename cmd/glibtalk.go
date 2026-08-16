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

	engine := glibness.NewEngine(dialogues)
	err = engine.Start("main")

	if err != nil {
		fmt.Println(err)
		return
	}

	for engine.Active() {
		response, err := engine.Next()

		if err != nil {
			fmt.Println(err)
			return
		}

		switch response := response.(type) {

		case glibness.SayResponse:
			fmt.Printf("[%s]: %s\n", response.Speaker, response.Sentence)
			// TODO find something better
			time.Sleep(time.Second)

		case glibness.ChoiceResponse:
			for i, choice := range response.Choices {
				fmt.Printf("%d. %s\n", i+1, choice.Name)
			}

			validAnswer := false
			var index int
			for !validAnswer {
				fmt.Printf("> ")
				_, err = fmt.Scanf("%d", &index)
				if err != nil {
					fmt.Printf("[%s] Hmm, I didn't quite catch that. Please select a number.\n", engine.Speaker())
					continue
				}

				err = engine.RespondIndex(index - 1)
				if err != nil {
					fmt.Printf("[%s] Hmm, I don't think that's right. Please select a valid number.\n", engine.Speaker())
					continue
				}

				validAnswer = true
			}

		}

	}
}
