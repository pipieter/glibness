package glibness

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestBasicTree(t *testing.T) {
	assert := assert.New(t)

	dialogues, err := Parse(`
		dialogue test {
			set speaker "Test"
			say "This is a test message"
			say "This is another test message" 
		}
	`)

	assert.Nil(err)
	assert.Len(dialogues, 1)

	dialogue := dialogues[0]
	assert.Equal(dialogue.name, "test")
	assert.Len(dialogue.statements, 3)
	assert.IsType(dialogue.statements[0], SetStatement{})
	assert.IsType(dialogue.statements[1], SayStatement{})
	assert.IsType(dialogue.statements[2], SayStatement{})
}
