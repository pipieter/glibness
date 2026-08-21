package glibness

import (
	"fmt"
	"strings"
)

type EngineState struct {
	Dialogue *Dialogue
	Index    int
}

type Engine struct {
	Dialogues []Dialogue
	Variables map[string]Value
	State     EngineState
	counter   int // Internal counter used to guarantee unique label ids
}

type StateResponse any

type SayResponse struct {
	Speaker  string
	Sentence string
}

type InternalResponse struct {
	Change string
}

type ChoiceResponse struct {
	Choices []string
}

type FinishedResponse struct {
}

func NewEngine() *Engine {
	engine := new(Engine)

	engine.Dialogues = make([]Dialogue, 0)
	engine.Variables = make(map[string]Value)
	engine.State.Dialogue = nil
	engine.State.Index = 0
	engine.counter = 0

	return engine
}

func (engine *Engine) IncCounter() int {
	engine.counter += 1
	return engine.counter
}

func (engine Engine) Active() bool {
	return engine.State.Dialogue != nil
}

func (state EngineState) Finished() bool {
	return state.Index >= len(state.Dialogue.Statements)
}

func (state EngineState) Statement() (Statement, error) {
	if state.Dialogue == nil {
		return nil, fmt.Errorf("No dialogue loaded to get statement.")
	}

	if state.Index < 0 {
		return nil, fmt.Errorf("Statement index for %s is less than zero.", state.Dialogue.Name)
	}

	if state.Finished() {
		return nil, fmt.Errorf("Dialogue %s has finished executing statements.", state.Dialogue.Name)
	}

	return state.Dialogue.Statements[state.Index], nil
}

func (engine *Engine) Start(name string) error {
	if engine.Active() {
		return fmt.Errorf("A dialogue state named '%s' is already active.", engine.DialogueName())
	}

	for i, dialogue := range engine.Dialogues {
		if dialogue.Name == name {
			engine.State.Dialogue = &engine.Dialogues[i]
			engine.State.Index = 0

			engine.Variables["speaker"] = MakeStringValue("")
			engine.Variables["sentence"] = MakeStringValue("")
			engine.Variables["dialogue"] = MakeStringValue(name)

			return nil
		}
	}

	return fmt.Errorf("Could not find dialogue with name '%s'", name)
}

func (engine Engine) Speaker() string {
	return engine.InterpolateString(engine.Variables["speaker"].String())
}

func (engine Engine) Sentence() string {
	return engine.InterpolateString(engine.Variables["sentence"].String())
}

func (engine Engine) DialogueName() string {
	return engine.InterpolateString(engine.Variables["dialogue"].String())
}

func (engine *Engine) Set(key string, value Value) error {
	// It's important to evaluate the values here because the set operator
	// sets the variable by value, not by reference. Pointer-esque operators
	// are not supported in Glibness.
	evaluated, err := value.Evaluate(*engine)

	if err != nil {
		return err
	} else {
		engine.Variables[key] = evaluated
		return nil
	}
}

func (state *EngineState) Jump(label string) error {
	for i, statement := range state.Dialogue.Statements {
		if statement, ok := statement.(LabelStatement); ok {
			if statement.Name == label {
				state.Index = i
				return nil
			}
		}
	}
	return fmt.Errorf("Could not find label %s in dialogue %s.", label, state.Dialogue.Name)
}

func (engine *Engine) execute(statement Statement) (StateResponse, error) {
	if engine.State.Dialogue == nil {
		return FinishedResponse{}, fmt.Errorf("Engine is currently not executing a dialogue.")
	}

	switch statement := statement.(type) {

	case SayStatement:
		err := engine.Set("sentence", statement.Sentence)
		return SayResponse{Sentence: engine.Sentence(), Speaker: engine.Speaker()}, err

	case LabelStatement:
		return InternalResponse{Change: "visit statement"}, nil

	case SetStatement:
		err := engine.Set(statement.Variable, statement.Value)
		return InternalResponse{Change: "set variable"}, err

	case ChooseStatement:
		return ChoiceResponse{Choices: statement.Names()}, nil

	case UnconditionalJumpStatement:
		err := engine.State.Jump(statement.Label)
		return InternalResponse{Change: "jump"}, err
	}

	return FinishedResponse{}, fmt.Errorf("Unsupported statement: %s", statement.String())
}

func (engine *Engine) InterpolateString(str string) string {
	for variable := range engine.Variables {
		pattern := fmt.Sprintf("{{%s}}", variable)
		str = strings.ReplaceAll(str, pattern, engine.Variables[variable].String())
	}
	return str
}

func (engine *Engine) Finish() error {
	engine.Variables["speaker"] = MakeStringValue("")
	engine.Variables["sentence"] = MakeStringValue("")
	engine.Variables["dialogue"] = MakeStringValue("")
	engine.State.Dialogue = nil
	engine.State.Index = 0
	return nil
}

func (engine *Engine) Next() (StateResponse, error) {
	if !engine.Active() {
		return FinishedResponse{}, nil
	}

	// A while-loop is used until a non-internal response is returned
	for true {
		// Stop if the dialogue has finished executing
		if engine.State.Finished() {
			err := engine.Finish()
			return FinishedResponse{}, err
		}

		current, err := engine.State.Statement()
		if err != nil {
			return nil, err
		}

		// If currently waiting for a ChooseResponse, we are waiting for input and we thus repeat the choices
		if choose, ok := current.(ChooseStatement); ok {
			return ChoiceResponse{Choices: choose.Names()}, nil
		}

		response, err := engine.execute(current)
		if err != nil {
			return response, err
		}
		engine.State.Index += 1

		// Stop if we encounter a non-internal response
		if _, internal := response.(InternalResponse); !internal {
			return response, err
		}
	}

	return FinishedResponse{}, nil // Should be unreachable
}

func (engine *Engine) Respond(index int) error {
	if !engine.Active() {
		return fmt.Errorf("The engine state is currently not executing any dialogues.")
	}

	current, err := engine.State.Statement()

	if err != nil {
		return err
	}

	if choose, ok := current.(ChooseStatement); ok {
		if index < 0 || index >= len(choose.Choices) {
			return fmt.Errorf("Invalid index %d for %d choices", index, len(choose.Choices))
		}

		return engine.State.Jump(choose.Choices[index].Label)
	}

	return fmt.Errorf("Dialogue is currently not expecting a choice.")
}
