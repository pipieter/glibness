package glibness

import (
	"fmt"
)

type Engine struct {
	Dialogues []Dialogue
}

type DialogueState struct {
	Engine   *Engine
	Dialogue *Dialogue

	Block *StatementBlock
	Index int

	Speaker   string
	Sentence  string
	Variables map[string]string
	Finished  bool
}

const (
	DialogueChangeError    = iota
	DialogueChangeStop     = iota
	DialogueChangeSay      = iota
	DialogueChangeInternal = iota
)

func (engine *Engine) Start(name string) (DialogueState, error) {
	for _, dialogue := range engine.Dialogues {
		if dialogue.Name == name {
			return DialogueState{
				Engine:    engine,
				Dialogue:  &dialogue,
				Block:     &dialogue.Root,
				Index:     0,
				Speaker:   "",
				Sentence:  "",
				Variables: make(map[string]string),
				Finished:  false,
			}, nil
		}
	}

	return DialogueState{}, fmt.Errorf("Could not find dialogue with name '%s'", name)
}

func (state *DialogueState) execute(s Statement) (int, error) {
	if say, ok := s.(SayStatement); ok {
		state.Sentence = say.Sentence
		return DialogueChangeSay, nil
	}

	if set, ok := s.(SetStatement); ok {
		if set.Variable == "speaker" {
			state.Speaker = set.Value
		} else {
			state.Variables[set.Variable] = set.Value
		}
		return DialogueChangeInternal, nil
	}

	return DialogueChangeError, fmt.Errorf("Unsupported statement: %s", s.String())
}

func (state *DialogueState) Next() (int, error) {
	if state.Finished {
		return DialogueChangeStop, nil
	}

	if state.Index >= len(state.Dialogue.Root.Statements) {
		state.Finished = true
		return DialogueChangeStop, nil
	}

	statement := state.Block.Statements[state.Index]
	status, err := state.execute(statement)
	state.Index += 1

	return status, err
}
