package test

import (
	"glibness/glibness"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestBasicTree(t *testing.T) {
	assert := assert.New(t)

	dialogues, err := glibness.Parse(`
		dialogue test {
			set speaker "Test"
			say "This is a test message"
			say "This is another test message" 
		}
	`)

	assert.Nil(err)
	assert.Len(dialogues, 1)

	dialogue := dialogues[0]
	assert.Equal(dialogue.Name, "test")
	assert.Len(dialogue.Root.Statements, 3)
	assert.IsType(dialogue.Root.Statements[0], glibness.SetStatement{})
	assert.IsType(dialogue.Root.Statements[1], glibness.SayStatement{})
	assert.IsType(dialogue.Root.Statements[2], glibness.SayStatement{})
}
