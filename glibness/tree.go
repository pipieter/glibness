package glibness

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
}

type Dialogue struct {
	Node
	Name string
	Root StatementBlock
}

type StatementBlock struct {
	Parent     *StatementBlock
	Statements []Statement
}

type SayStatement struct {
	Statement
	Sentence string
}

type SetStatement struct {
	Statement
	Variable string
	Value    string
}

type ChooseStatement struct {
	Statement
	Choices []Choice
}

type Choice struct {
	Name  string
	Block StatementBlock
}

func (s SayStatement) String() string {
	return fmt.Sprintf("say '%s'", s.Sentence)
}

func (s SetStatement) String() string {
	return fmt.Sprintf("set %s '%s'", s.Variable, s.Value)
}

func (s ChooseStatement) String() string {
	choices := make([]string, 0)
	for _, choice := range s.Choices {
		choices = append(choices, choice.Name)
	}
	joined := strings.Join(choices, "\", \"")
	return fmt.Sprintf("choose { \"%s\" }", joined)
}
