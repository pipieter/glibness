package glibness

import (
	"fmt"
	"strings"
)

type Engine struct {
	Dialogues []Dialogue
}

type DialogueState struct {
	Engine   *Engine
	Dialogue *Dialogue

	Blocks []DialogueBlock

	Speaker   string
	Sentence  string
	Variables map[string]string
	Finished  bool

	CurrentChoices []Choice
}

type DialogueBlock struct {
	Block *StatementBlock
	Index int
}

type StateResponse interface {
}

type SayResponse struct {
	Speaker  string
	Sentence string
}

type InternalChangeResponse struct {
	Change string
}

type ChoiceResponse struct {
	Choices []Choice
}

type FinishedResponse struct {
}

func (engine *Engine) Start(name string) (DialogueState, error) {
	for _, dialogue := range engine.Dialogues {
		if dialogue.Name == name {
			block := DialogueBlock{Block: &dialogue.Root, Index: 0}
			return DialogueState{
				Engine:    engine,
				Dialogue:  &dialogue,
				Blocks:    []DialogueBlock{block},
				Speaker:   "",
				Sentence:  "",
				Variables: make(map[string]string),
				Finished:  false,
			}, nil
		}
	}

	return DialogueState{}, fmt.Errorf("Could not find dialogue with name '%s'", name)
}

func (state *DialogueState) execute(statement Statement) (StateResponse, error) {
	switch statement := statement.(type) {

	case SayStatement:
		state.Sentence = statement.Sentence
		return SayResponse{Speaker: state.Speaker, Sentence: state.Sentence}, nil

	case SetStatement:
		if statement.Variable == "speaker" {
			state.Speaker = statement.Value
		} else {
			state.Variables[statement.Variable] = statement.Value
		}
		return InternalChangeResponse{Change: "set variable"}, nil

	case ChooseStatement:
		state.CurrentChoices = statement.Choices
		return ChoiceResponse{Choices: statement.Choices}, nil
	}

	return FinishedResponse{}, fmt.Errorf("Unsupported statement: %s", statement.String())
}

func (state *DialogueState) CurrentBlock() *DialogueBlock {
	if len(state.Blocks) == 0 {
		return nil
	}
	return &state.Blocks[len(state.Blocks)-1]
}

// Pop the last dialogue block and return the length of the new blocks
func (state *DialogueState) PopBlock() int {
	if len(state.Blocks) == 0 {
		return 0
	}

	state.Blocks = state.Blocks[:len(state.Blocks)-1]
	return len(state.Blocks)
}

func (state *DialogueState) Next() (StateResponse, error) {
	// If currently waiting for a ChooseResponse, repeat it
	if state.CurrentChoices != nil {
		return ChoiceResponse{Choices: state.CurrentChoices}, nil
	}

	if state.Finished {
		return FinishedResponse{}, nil
	}

	block := state.CurrentBlock()

	// Block has finished executing
	for block.Index >= len(block.Block.Statements) {
		remaining := state.PopBlock()
		if remaining == 0 {
			state.Finished = true
			return FinishedResponse{}, nil
		}
		block = state.CurrentBlock()
	}

	statement := block.Block.Statements[block.Index]
	status, err := state.execute(statement)
	block.Index += 1

	return status, err
}

func (state *DialogueState) RespondIndex(index int) error {
	if state.CurrentChoices == nil {
		return fmt.Errorf("The engine state is currently not expecting any choices.")
	}

	if index < 0 || index >= len(state.CurrentChoices) {
		return fmt.Errorf("Invalid index %d for %d choices", index, len(state.CurrentChoices))
	}

	choice := state.CurrentChoices[index]

	state.Blocks = append(state.Blocks, DialogueBlock{Block: &choice.Block, Index: 0})
	state.CurrentChoices = nil

	return nil
}

func (state *DialogueState) Respond(choice string) error {
	for i, possibleChoice := range state.CurrentChoices {
		if possibleChoice.Name == choice {
			return state.RespondIndex(i)
		}
	}

	choices := make([]string, 0)
	for _, choice := range state.CurrentChoices {
		choices = append(choices, choice.Name)
	}
	joined := strings.Join(choices, ", ")
	return fmt.Errorf("Choice %s is not a possible choice of %s.", choice, joined)
}
