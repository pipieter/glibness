package test

import (
	"glibness/glibness"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestTrailingWhitespace(t *testing.T) {
	assert := assert.New(t)
	engine := glibness.NewEngine()

	err := engine.ParseString(`
		dialogue test {
			say "This is a message with trailing whitespaces."     
			say "This message also has trailing whitespaces." 		    
		}
	`)
	assert.Nil(err)
}

func TestMissingClosingBracket(t *testing.T) {
	assert := assert.New(t)
	engine := glibness.NewEngine()

	err := engine.ParseString(`
		dialogue test {
			say "This is a dialogue without a closing bracket."
	`)
	assert.NotNil(err)
}
