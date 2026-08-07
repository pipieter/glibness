package glibness

import (
	"fmt"
	parser "glibness/grammar"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

type GlibnessListener struct {
	*parser.BaseGlibnessListener
}

func parseDialogue(node parser.IDialogueContext) (Dialogue, error) {
	name := node.GetName().GetText()
	statements, err := parseStatements(node.Statements())

	if err != nil {
		return Dialogue{}, err
	}

	return Dialogue{name: name, statements: statements}, nil
}

func parseStatements(node parser.IStatementsContext) ([]Statement, error) {
	statements := make([]Statement, 0)

	for _, statement := range node.GetChildren() {
		if statement, ok := statement.(*parser.StatementContext); ok {
			statement, err := parseStatement(statement)

			if err != nil {
				return nil, err
			}

			statements = append(statements, statement)
		} else {
			return nil, fmt.Errorf("Unknown statements type: %s", reflect.TypeOf(statement))
		}
	}

	return statements, nil
}

func parseStatement(node parser.IStatementContext) (Statement, error) {
	statement := node.GetChildren()[0]

	if set, ok := statement.(parser.ISetStatementContext); ok {
		set, err := parseSetStatement(set)

		if err != nil {
			return nil, err
		}

		return set, nil

	} else if say, ok := statement.(parser.ISayStatementContext); ok {
		say, err := parseSayStatement(say)

		if err != nil {
			return nil, err
		}

		return say, nil
	}

	return nil, fmt.Errorf("Unsupported statement type: %s", reflect.TypeOf(node))
}

func parseSetStatement(node parser.ISetStatementContext) (SetStatement, error) {
	variable := node.GetVariable().GetText()
	value := node.GetVal().GetText()
	return SetStatement{variable: variable, value: value}, nil
}

func parseSayStatement(node parser.ISayStatementContext) (SayStatement, error) {
	value := node.GetVal().GetText()
	return SayStatement{sentence: value}, nil
}

func Parse(input string) ([]Dialogue, error) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewGlibnessLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGlibnessParser(stream)

	tree := p.Program()
	dialogues := make([]Dialogue, 0)

	for _, child := range tree.GetChildren() {
		if dialogue, ok := child.(parser.IDialogueContext); ok {
			dialogue, err := parseDialogue(dialogue)

			if err != nil {
				return nil, err
			}

			dialogues = append(dialogues, dialogue)

		}

		// TODO: handle other types
	}

	return dialogues, nil
}
