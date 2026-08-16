package test

import (
	"glibness/glibness"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestVariables(t *testing.T) {
	assert := assert.New(t)
	engine := glibness.NewEngine()

	err := engine.ParseString(`
		dialogue test {
			set speaker "test"
			set my_int  33
			set my_str  "str"
			set my_bool true
			set my_var  my_str

			set my_override true
			set my_override "999"
			set my_override my_int
			set my_override false

			set my_self_reference 4
			set my_self_reference my_self_reference

			set my_copy_1 true
			set my_copy_2 my_copy_1
			set my_copy_1 my_copy_2
			set my_copy_2 my_copy_1
			set my_copy_1 my_copy_2
			set my_copy_2 my_copy_1
			set my_copy_1 my_copy_2
			set my_copy_2 my_copy_1
			set my_copy_1 my_copy_2			
		}
	`)
	assert.Nil(err)

	err = engine.Start("test")
	assert.Nil(err)

	for engine.Active() {
		_, err := engine.Next()
		assert.Nil(err)
	}

	// Ensure speaker gets set to an empty string at the end of the dialogue
	assert.IsType(engine.Variables["speaker"], glibness.StringValue{})
	assert.Equal(engine.Variables["speaker"].String(), "")

	assert.IsType(engine.Variables["my_int"], glibness.IntegerValue{})
	assert.Equal(engine.Variables["my_int"].String(), "33")

	assert.IsType(engine.Variables["my_str"], glibness.StringValue{})
	assert.Equal(engine.Variables["my_str"].String(), "str")

	assert.IsType(engine.Variables["my_bool"], glibness.BooleanValue{})
	assert.Equal(engine.Variables["my_bool"].String(), "true")

	assert.IsType(engine.Variables["my_var"], glibness.StringValue{})
	assert.Equal(engine.Variables["my_var"].String(), "str")

	assert.IsType(engine.Variables["my_override"], glibness.BooleanValue{})
	assert.Equal(engine.Variables["my_override"].String(), "false")

	assert.IsType(engine.Variables["my_self_reference"], glibness.IntegerValue{})
	assert.Equal(engine.Variables["my_self_reference"].String(), "4")

	assert.IsType(engine.Variables["my_copy_1"], glibness.BooleanValue{})
	assert.IsType(engine.Variables["my_copy_2"], glibness.BooleanValue{})
	assert.Equal(engine.Variables["my_copy_1"].String(), "true")
	assert.Equal(engine.Variables["my_copy_2"].String(), "true")

}
