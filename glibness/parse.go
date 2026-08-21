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
	statements, err := engine.parseStatementBlock(name, node.StatementBlock())
	return Dialogue{Name: name, Statements: statements}, err
}

func (engine *Engine) parseStatementBlock(dialogue string, node parser.IStatementBlockContext) ([]Statement, error) {
	statements := make([]Statement, 0)

	for _, statement := range node.AllStatement() {
		if statement, ok := statement.(*parser.StatementContext); ok {
			substatements, err := engine.parseStatement(dialogue, statement)

			if err != nil {
				return nil, err
			}

			statements = append(statements, substatements...)
		} else {
			return nil, fmt.Errorf("Unknown statements type: %s", reflect.TypeOf(statement))
		}
	}

	return statements, nil
}

func (engine *Engine) parseStatement(dialogue string, node parser.IStatementContext) ([]Statement, error) {
	child := node.GetChildren()[0]

	switch statement := child.(type) {

	case parser.ISetStatementContext:
		return engine.parseSetStatement(statement)

	case parser.ISayStatementContext:
		return engine.parseSayStatement(statement)

	case parser.IChooseStatementContext:
		return engine.parseChooseStatement(dialogue, statement)
	}

	return nil, fmt.Errorf("Unsupported statement type: %s", reflect.TypeOf(node))
}

func (engine *Engine) parseSetStatement(node parser.ISetStatementContext) ([]Statement, error) {
	variable := node.GetVariable().GetText()
	value, err := engine.parseValue(node.GetVal())

	if err != nil {
		return nil, err
	}

	return []Statement{SetStatement{Variable: variable, Value: value}}, err
}

func (engine *Engine) parseSayStatement(node parser.ISayStatementContext) ([]Statement, error) {
	value, err := engine.parseValue(node.GetVal())
	if err != nil {
		return nil, err
	}

	return []Statement{SayStatement{Sentence: value}}, err
}

func (engine *Engine) parseChooseStatement(dialogue string, node parser.IChooseStatementContext) ([]Statement, error) {
	statements := make([]Statement, 0)
	substatements := make([]Statement, 0)

	chooseId := engine.IncCounter()
	exitLabel := fmt.Sprintf("L%d_choose_exit", chooseId)

	choices := make([]Choice, 0)

	for i, choice := range node.ChoiceBlock().AllChoice() {
		choice, choiceStatements, err := engine.parseChoice(dialogue, chooseId, i, choice)

		if err != nil {
			return nil, err
		}

		choiceStatements = append(choiceStatements, UnconditionalJumpStatement{Label: exitLabel})
		substatements = append(substatements, choiceStatements...)
		choices = append(choices, choice)
	}

	statements = append(statements, ChooseStatement{Choices: choices})
	statements = append(statements, substatements...)
	statements = append(statements, LabelStatement{Name: exitLabel})

	// Check if any duplicate choices were found
	uniqueChoices := map[string]int{}
	for _, choice := range choices {
		_, ok := uniqueChoices[choice.Name]
		if ok {
			return nil, fmt.Errorf("Choice %s was found multiple times in dialogue.", choice.Name)
		} else {
			uniqueChoices[choice.Name] = 1
		}
	}

	return statements, nil
}

func (engine *Engine) parseChoice(dialogue string, chooseId int, index int, node parser.IChoiceContext) (Choice, []Statement, error) {
	name := (node.GetName().GetText())
	name = strings.TrimSuffix(name, "\"")
	name = strings.TrimPrefix(name, "\"")

	label := fmt.Sprintf("L%d_choose_%d_'%s'", chooseId, index, name)
	choice := Choice{Name: name, Label: label}

	statements := []Statement{LabelStatement{Name: label}}
	substatements, err := engine.parseStatementBlock(dialogue, node.GetBlock())

	if err != nil {
		return Choice{}, nil, err
	}

	statements = append(statements, substatements...)
	return choice, statements, nil

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
	if node.VARIABLE() != nil {
		return engine.parseVariable(node.VARIABLE()), nil
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
	parse := parser.NewGlibnessParser(stream)

	// By default, antlr4 prints errors out to the console. In our case,
	// we want to capture errors and return them, so we need to create a
	// special error capture class.
	errorCapture := &ErrorCapture{Errors: make([]string, 0)}
	parse.RemoveErrorListeners()
	parse.AddErrorListener(errorCapture)

	tree := parse.Program()

	for _, child := range tree.AllDialogue() {
		dialogue, err := engine.parseDialogue(child)

		if err != nil {
			return err
		}

		engine.Dialogues = append(engine.Dialogues, dialogue)
	}

	return errorCapture.Error()
}
