package glibness

import (
	"fmt"
)

type Engine struct {
	dialogues map[string]Dialogue
}

type DialogueExecution struct {
	engine    *Engine
	speaker   string
	sentence  string
	variables map[string]string
}

func (e *DialogueExecution) execute(s Statement) error {
	if say, ok := s.(SayStatement); ok {
		e.sentence = say.sentence
		return nil
	}

	if set, ok := s.(SetStatement); ok {
		if set.variable == "speaker" {
			e.speaker = set.value
		} else {
			e.variables[set.variable] = set.value
		}
		return nil
	}

	return fmt.Errorf("Unsupported statement: %s", s.String())
}
