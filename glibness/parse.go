package glibness

import (
	"fmt"
	parser "glibness/grammar"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type GlibnessListener struct {
	*parser.BaseGlibnessListener
}

func parseDialogue(node parser.IDialogueContext) (Dialogue, error) {
	name := node.GetName().GetText()
	statements, err := parseStatementBlock(node.StatementBlock(), nil)

	if err != nil {
		return Dialogue{}, err
	}

	return Dialogue{Name: name, Root: statements}, nil
}

func parseStatementBlock(node parser.IStatementBlockContext, parent *StatementBlock) (StatementBlock, error) {
	block := StatementBlock{Statements: make([]Statement, 0), Parent: parent}

	for _, statement := range node.AllStatement() {
		if statement, ok := statement.(*parser.StatementContext); ok {
			statement, err := parseStatement(statement, &block)

			if err != nil {
				return StatementBlock{}, err
			}

			block.Statements = append(block.Statements, statement)
		} else {
			return StatementBlock{}, fmt.Errorf("Unknown statements type: %s", reflect.TypeOf(statement))
		}
	}

	return block, nil
}

func parseStatement(node parser.IStatementContext, parent *StatementBlock) (Statement, error) {
	child := node.GetChildren()[0]

	switch statement := child.(type) {

	case parser.ISetStatementContext:
		return parseSetStatement(statement)

	case parser.ISayStatementContext:
		return parseSayStatement(statement)

	case parser.IChooseStatementContext:
		return parseChooseStatement(statement, parent)
	}

	return nil, fmt.Errorf("Unsupported statement type: %s", reflect.TypeOf(node))
}

func parseSetStatement(node parser.ISetStatementContext) (SetStatement, error) {
	variable := node.GetVariable().GetText()
	value := parseValue(node.GetVal())
	return SetStatement{Variable: variable, Value: value}, nil
}

func parseSayStatement(node parser.ISayStatementContext) (SayStatement, error) {
	value := parseValue(node.GetVal())
	return SayStatement{Sentence: value}, nil
}

func parseChooseStatement(node parser.IChooseStatementContext, parent *StatementBlock) (ChooseStatement, error) {
	choices := make([]Choice, 0)

	for _, choice := range node.ChoiceBlock().AllChoice() {
		choice, err := parseChoice(choice, parent)

		if err != nil {
			return ChooseStatement{}, err
		}

		choices = append(choices, choice)
	}

	return ChooseStatement{Choices: choices}, nil
}

func parseChoice(node parser.IChoiceContext, parent *StatementBlock) (Choice, error) {
	name := trimStringQuotes(node.GetName().GetText())
	block, err := parseStatementBlock(node.GetBlock(), parent)

	if err != nil {
		return Choice{}, err
	}

	return Choice{Name: name, Block: block}, nil

}

func parseValue(node parser.IValueContext) string {
	// For now only strings are supported...
	value := node.GetText()
	return trimStringQuotes(value)
}

func trimStringQuotes(value string) string {
	value = strings.TrimSuffix(value, "\"")
	value = strings.TrimPrefix(value, "\"")
	return value
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
