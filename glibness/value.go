package glibness

import (
	"fmt"
	"reflect"
	"strconv"
)

type Value interface {
	String() string
	Type() string
	Evaluate(engine Engine) (Value, error)
}

type StringValue struct {
	Value string
}

type BooleanValue struct {
	Value bool
}

type IntegerValue struct {
	Value int
}

type VariableValue struct {
	Engine   *Engine
	Variable string
}

func MakeStringValue(value string) StringValue {
	return StringValue{Value: value}
}

func MakeBooleanValue(value bool) BooleanValue {
	return BooleanValue{Value: value}
}

func MakeIntValue(value int) IntegerValue {
	return IntegerValue{Value: value}
}

func MakeVariableValue(engine *Engine, variable string) VariableValue {
	return VariableValue{Engine: engine, Variable: variable}
}

func (variable StringValue) String() string {
	return variable.Value
}

func (variable BooleanValue) String() string {
	if variable.Value {
		return "true"
	} else {
		return "false"
	}
}

func (variable IntegerValue) String() string {
	return strconv.Itoa(variable.Value)
}

func (variable VariableValue) String() string {
	value := variable.Engine.Variables[variable.Variable].String()
	return variable.Engine.InterpolateString(value)
}

func (variable StringValue) Type() string {
	return "string"
}

func (variable BooleanValue) Type() string {
	return "boolean"
}

func (variable IntegerValue) Type() string {
	return "integer"
}

func (variable VariableValue) Type() string {
	return "variable"
}

func (variable StringValue) Evaluate(engine Engine) (Value, error) {
	value := variable.Value
	value = engine.InterpolateString(value)
	return MakeStringValue(value), nil
}

func (variable BooleanValue) Evaluate(engine Engine) (Value, error) {
	return MakeBooleanValue(variable.Value), nil
}

func (variable IntegerValue) Evaluate(engine Engine) (Value, error) {
	return MakeIntValue(variable.Value), nil
}

func (variable VariableValue) Evaluate(engine Engine) (Value, error) {
	value := variable.Engine.Variables[variable.Variable]

	switch value := value.(type) {

	case VariableValue:
		return value.Evaluate(engine)

	case StringValue:
		return MakeStringValue(value.Value), nil

	case IntegerValue:
		return MakeIntValue(value.Value), nil

	case BooleanValue:
		return MakeBooleanValue(value.Value), nil
	}

	return VariableValue{}, fmt.Errorf("Unsupported Value.Evaluate: '%s'", reflect.TypeOf(value))
}

func Equals(var1 Value, var2 Value) bool {
	// Check if the two types are the same type. If not, return false
	if reflect.TypeOf(var1) != reflect.TypeOf(var2) {
		_, isvar1 := var1.(VariableValue)
		_, isvar2 := var2.(VariableValue)
		if !(isvar1 || isvar2) {
			return false
		}
	}

	// Otherwise, to simplify things, check if they are string equals
	return var1.String() == var2.String()
}
