package test

import (
	"glibness/glibness"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestBasicTree(t *testing.T) {
	assert := assert.New(t)

	engine := glibness.NewEngine()
	err := engine.ParseString(`
		dialogue test {
			set speaker "Test"
			say "This is a test message"
			say "This is another test message" 
		}
	`)

	assert.Nil(err)
	assert.Len(engine.Dialogues, 1)

	dialogue := engine.Dialogues[0]
	assert.Equal(dialogue.Name, "test")
	assert.Len(dialogue.Statements, 3)
}
