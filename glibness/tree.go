package glibness

import "fmt"

type Node interface {
	String() string
}

type Statement interface {
	Node
}

type Dialogue struct {
	Node
	name       string
	statements []Statement
}

type SayStatement struct {
	Statement
	sentence string
}

type SetStatement struct {
	Statement
	variable string
	value    string
}

func (s SayStatement) String() string {
	return fmt.Sprintf("say '%s'", s.sentence)
}

func (s SetStatement) String() string {
	return fmt.Sprintf("set %s '%s'", s.variable, s.value)
}
