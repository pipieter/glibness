package glibness

import (
	"fmt"
	"strings"
)

type EngineState struct {
	Active         bool
	Blocks         []DialogueBlock
	CurrentChoices []Choice
}

type Engine struct {
	Dialogues []Dialogue
	Variables map[string]string
	State     EngineState
}

type DialogueBlock struct {
	Block *StatementBlock
	Index int
}

type StateResponse any

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

// TODO change this
func NewEngine(dialogues []Dialogue) Engine {
	return Engine{
		Dialogues: dialogues,
		Variables: make(map[string]string),
		State: EngineState{
			Active:         false,
			Blocks:         nil,
			CurrentChoices: nil,
		},
	}
}

func (engine *Engine) Start(name string) error {
	if engine.State.Active {
		return fmt.Errorf("A dialogue state named '%s' is already active.", engine.DialogueName())
	}

	for _, dialogue := range engine.Dialogues {
		if dialogue.Name == name {
			block := DialogueBlock{Block: &dialogue.Root, Index: 0}

			engine.State.Active = true
			engine.State.Blocks = []DialogueBlock{block}
			engine.State.CurrentChoices = nil

			engine.Variables["speaker"] = ""
			engine.Variables["sentence"] = ""
			engine.Variables["dialogue"] = name

			return nil
		}
	}

	return fmt.Errorf("Could not find dialogue with name '%s'", name)
}

func (engine Engine) Speaker() string {
	return engine.Variables["speaker"]
}

func (engine Engine) Sentence() string {
	return engine.Variables["sentence"]
}

func (engine Engine) DialogueName() string {
	return engine.Variables["dialogue"]
}

func (engine *Engine) execute(statement Statement) (StateResponse, error) {
	if !engine.State.Active {
		return FinishedResponse{}, fmt.Errorf("Engine is currently not executing a dialogue.")
	}

	switch statement := statement.(type) {

	case SayStatement:
		engine.Variables["sentence"] = engine.ResolveString(statement.Sentence)
		return SayResponse{Speaker: engine.Speaker(), Sentence: engine.Sentence()}, nil

	case SetStatement:
		engine.Variables[statement.Variable] = statement.Value
		return InternalChangeResponse{Change: "set variable"}, nil

	case ChooseStatement:
		engine.State.CurrentChoices = statement.Choices
		return ChoiceResponse{Choices: statement.Choices}, nil
	}

	return FinishedResponse{}, fmt.Errorf("Unsupported statement: %s", statement.String())
}

func (engine *Engine) ResolveString(str string) string {
	for variable := range engine.Variables {
		pattern := fmt.Sprintf("{{%s}}", variable)
		str = strings.ReplaceAll(str, pattern, engine.Variables[variable])
	}
	return str
}

func (engine *Engine) Finish() error {
	engine.Variables["speaker"] = ""
	engine.Variables["sentence"] = ""
	engine.Variables["dialogue"] = ""
	engine.State.Active = false
	engine.State.Blocks = nil
	engine.State.CurrentChoices = nil
	return nil
}

func (engine *Engine) CurrentBlock() *DialogueBlock {
	if len(engine.State.Blocks) == 0 {
		return nil
	}
	return &engine.State.Blocks[len(engine.State.Blocks)-1]
}

// Pop the last dialogue block and return the length of the new blocks
func (engine *Engine) PopBlock() int {
	if len(engine.State.Blocks) == 0 {
		return 0
	}

	engine.State.Blocks = engine.State.Blocks[:len(engine.State.Blocks)-1]
	return len(engine.State.Blocks)
}

func (engine *Engine) Next() (StateResponse, error) {
	if !engine.State.Active {
		return FinishedResponse{}, nil
	}

	// If currently waiting for a ChooseResponse, repeat it
	if engine.State.CurrentChoices != nil {
		return ChoiceResponse{Choices: engine.State.CurrentChoices}, nil
	}

	block := engine.CurrentBlock()

	// Block has finished executing
	for block.Index >= len(block.Block.Statements) {
		remaining := engine.PopBlock()
		if remaining == 0 {
			err := engine.Finish()
			return FinishedResponse{}, err
		}
		block = engine.CurrentBlock()
	}

	statement := block.Block.Statements[block.Index]
	status, err := engine.execute(statement)
	block.Index += 1

	return status, err
}

func (engine *Engine) Respond(index int) error {
	if !engine.State.Active {
		return fmt.Errorf("The engine state is currently not executing any dialogues.")
	}

	if engine.State.CurrentChoices == nil {
		return fmt.Errorf("The engine state is currently not expecting any choices.")
	}

	if index < 0 || index >= len(engine.State.CurrentChoices) {
		return fmt.Errorf("Invalid index %d for %d choices", index, len(engine.State.CurrentChoices))
	}

	choice := engine.State.CurrentChoices[index]

	engine.State.Blocks = append(engine.State.Blocks, DialogueBlock{Block: &choice.Block, Index: 0})
	engine.State.CurrentChoices = nil

	return nil
}

func (engine *Engine) Active() bool {
	return engine.State.Active
}
