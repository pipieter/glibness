package glibness

import (
	"reflect"
	"strconv"
)

type Value interface {
	String() string
}

type StringValue struct {
	value string
}

type BooleanValue struct {
	value bool
}

type IntegerValue struct {
	value int
}

type VariableValue struct {
	Engine   *Engine
	Variable string
}

func MakeStringValue(value string) StringValue {
	return StringValue{value: value}
}

func MakeBooleanValue(value bool) BooleanValue {
	return BooleanValue{value: value}
}

func MakeIntValue(value int) IntegerValue {
	return IntegerValue{value: value}
}

func MakeVariableValue(engine *Engine, variable string) VariableValue {
	return VariableValue{Engine: engine, Variable: variable}
}

func (variable StringValue) String() string {
	return variable.value
}

func (variable BooleanValue) String() string {
	if variable.value {
		return "true"
	} else {
		return "false"
	}
}

func (variable IntegerValue) String() string {
	return strconv.Itoa(variable.value)
}

func (variable VariableValue) String() string {
	value := variable.Engine.Variables[variable.Variable].String()
	return variable.Engine.ResolveString(value)
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
