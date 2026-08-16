package glibness

import (
	"fmt"
	parser "glibness/grammar"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type GlibnessListener struct {
	*parser.BaseGlibnessListener
}

func (engine *Engine) parseDialogue(node parser.IDialogueContext) (Dialogue, error) {
	name := node.GetName().GetText()
	statements, err := engine.parseStatementBlock(node.StatementBlock(), nil)

	if err != nil {
		return Dialogue{}, err
	}

	return Dialogue{Name: name, Root: statements}, nil
}

func (engine *Engine) parseStatementBlock(node parser.IStatementBlockContext, parent *StatementBlock) (StatementBlock, error) {
	block := StatementBlock{Statements: make([]Statement, 0), Parent: parent}

	for _, statement := range node.AllStatement() {
		if statement, ok := statement.(*parser.StatementContext); ok {
			statement, err := engine.parseStatement(statement, &block)

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

func (engine *Engine) parseStatement(node parser.IStatementContext, parent *StatementBlock) (Statement, error) {
	child := node.GetChildren()[0]

	switch statement := child.(type) {

	case parser.ISetStatementContext:
		return engine.parseSetStatement(statement)

	case parser.ISayStatementContext:
		return engine.parseSayStatement(statement)

	case parser.IChooseStatementContext:
		return engine.parseChooseStatement(statement, parent)
	}

	return nil, fmt.Errorf("Unsupported statement type: %s", reflect.TypeOf(node))
}

func (engine *Engine) parseSetStatement(node parser.ISetStatementContext) (SetStatement, error) {
	variable := node.GetVariable().GetText()
	value, err := engine.parseValue(node.GetVal())
	return SetStatement{Variable: variable, Value: value}, err
}

func (engine *Engine) parseSayStatement(node parser.ISayStatementContext) (SayStatement, error) {
	value, err := engine.parseValue(node.GetVal())
	return SayStatement{Sentence: value}, err
}

func (engine *Engine) parseChooseStatement(node parser.IChooseStatementContext, parent *StatementBlock) (ChooseStatement, error) {
	choices := make([]Choice, 0)

	for _, choice := range node.ChoiceBlock().AllChoice() {
		choice, err := engine.parseChoice(choice, parent)

		if err != nil {
			return ChooseStatement{}, err
		}

		choices = append(choices, choice)
	}

	// Check if any duplicate choices were found
	uniqueChoices := map[string]int{}
	for _, choice := range choices {
		_, ok := uniqueChoices[choice.Name]
		if ok {
			return ChooseStatement{}, fmt.Errorf("Choice %s was found multiple times in dialogue.", choice.Name)
		} else {
			uniqueChoices[choice.Name] = 1
		}
	}

	return ChooseStatement{Choices: choices}, nil
}

func (engine *Engine) parseChoice(node parser.IChoiceContext, parent *StatementBlock) (Choice, error) {
	name := (node.GetName().GetText())
	name = strings.TrimSuffix(name, "\"")
	name = strings.TrimPrefix(name, "\"")

	block, err := engine.parseStatementBlock(node.GetBlock(), parent)

	if err != nil {
		return Choice{}, err
	}

	return Choice{Name: name, Block: block}, nil

}

func (engine *Engine) parseValue(node parser.IValueContext) (Value, error) {
	if node.BOOLEAN() != nil {
		return engine.parseBoolean(node.BOOLEAN()), nil
	}
	if node.STRING() != nil {
		return engine.parseString(node.STRING()), nil
	}
	if node.INTEGER() != nil {
		return engine.parseInteger(node.INTEGER()), nil
	}

	return nil, fmt.Errorf("Unsupported value: '%s'", node.GetText())
}

func (engine *Engine) parseBoolean(node antlr.TerminalNode) BooleanValue {
	if node.GetText() == "true" {
		return MakeBooleanValue(true)
	} else {
		return MakeBooleanValue(false)

	}
}

func (engine *Engine) parseString(node antlr.TerminalNode) StringValue {
	text := node.GetText()
	text = strings.TrimSuffix(text, "\"")
	text = strings.TrimPrefix(text, "\"")
	return MakeStringValue(text)
}

func (engine *Engine) parseInteger(node antlr.TerminalNode) IntegerValue {
	value, _ := strconv.Atoi(node.GetText())
	return MakeIntValue(value)
}

func (engine *Engine) parseVariable(node antlr.TerminalNode) VariableValue {
	return MakeVariableValue(engine, node.GetText())
}

func (engine *Engine) ParseFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return engine.ParseString(string(bytes))
}

func (engine *Engine) ParseString(input string) error {
	is := antlr.NewInputStream(input)
	lexer := parser.NewGlibnessLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGlibnessParser(stream)

	tree := p.Program()

	for _, child := range tree.GetChildren() {
		if dialogue, ok := child.(parser.IDialogueContext); ok {
			dialogue, err := engine.parseDialogue(dialogue)

			if err != nil {
				return err
			}

			engine.Dialogues = append(engine.Dialogues, dialogue)

		}
	}

	return nil
}
