package glibness

import (
	"fmt"
	"strings"
)

type Dialogue struct {
	Name             string
	Statements       []Statement
	StatementCounter int
}

type Statement interface {
	String() string
}

type SayStatement struct {
	Sentence Value
}

type LabelStatement struct {
	Name string
}

type SetStatement struct {
	Variable string
	Value    Value
}

type ChooseStatement struct {
	Choices []Choice
}

type Choice struct {
	Name  string
	Label string
}

type UnconditionalJumpStatement struct {
	Label string
}

func (s SayStatement) String() string {
	return fmt.Sprintf("say %s", s.Sentence.String())
}

func (s SetStatement) String() string {
	return fmt.Sprintf("set %s '%s'", s.Variable, s.Value)
}

func (s LabelStatement) String() string {
	return fmt.Sprintf("label %s", s.Name)
}

func (s ChooseStatement) String() string {
	choices := make([]string, 0)
	for _, choice := range s.Choices {
		choices = append(choices, fmt.Sprintf("\"%s\" = %s", choice.Name, choice.Label))
	}
	return fmt.Sprintf("choose { %s }", strings.Join(choices, ", "))
}

func (s ChooseStatement) Names() []string {
	names := make([]string, 0)
	for _, choice := range s.Choices {
		names = append(names, choice.Name)
	}
	return names
}

func (s UnconditionalJumpStatement) String() string {
	return fmt.Sprintf("jump %s", s.Label)
}
