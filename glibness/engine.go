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
	Index    int

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
		if dialogue.name == name {
			return DialogueState{
				Engine:    engine,
				Dialogue:  &dialogue,
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
		state.Sentence = say.sentence
		return DialogueChangeSay, nil
	}

	if set, ok := s.(SetStatement); ok {
		if set.variable == "speaker" {
			state.Speaker = set.value
		} else {
			state.Variables[set.variable] = set.value
		}
		return DialogueChangeInternal, nil
	}

	return DialogueChangeError, fmt.Errorf("Unsupported statement: %s", s.String())
}

func (state *DialogueState) Next() (int, error) {
	if state.Finished {
		return DialogueChangeStop, nil
	}

	if state.Index >= len(state.Dialogue.statements) {
		state.Finished = true
		return DialogueChangeStop, nil
	}

	statement := state.Dialogue.statements[state.Index]
	status, err := state.execute(statement)
	state.Index += 1

	return status, err
}
