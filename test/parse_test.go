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

func TestChoiceTree(t *testing.T) {
	assert := assert.New(t)

	dialogues, err := glibness.Parse(`
		dialogue test {
			set speaker "Test"
			say "This has a choice message"
			choose {
				choice "A" {
					say "A.1"
					say "A.1"
				}
				choice "B" {
					say "B.1"
				}
				choice "C" {
					say "C.1"
					say "C.2"
					say "C.3"
				}
 			}
		}
	`)

	assert.Nil(err)
	assert.Len(dialogues, 1)

	dialogue := dialogues[0]
	assert.Equal(dialogue.Name, "test")
	assert.Len(dialogue.Root.Statements, 3)
	assert.IsType(dialogue.Root.Statements[0], glibness.SetStatement{})
	assert.IsType(dialogue.Root.Statements[1], glibness.SayStatement{})
	assert.IsType(dialogue.Root.Statements[2], glibness.ChooseStatement{})

	choose := dialogue.Root.Statements[2]
	if choose, ok := choose.(glibness.ChooseStatement); ok {
		assert.Len(choose.Choices, 3)

		assert.Equal(choose.Choices[0].Name, "A")
		assert.Equal(choose.Choices[1].Name, "B")
		assert.Equal(choose.Choices[2].Name, "C")

		assert.Len(choose.Choices[0].Block.Statements, 2)
		assert.Len(choose.Choices[1].Block.Statements, 1)
		assert.Len(choose.Choices[2].Block.Statements, 3)

		assert.Equal(choose.Choices[0].Block.Parent, &dialogue.Root)
		assert.Equal(choose.Choices[1].Block.Parent, &dialogue.Root)
		assert.Equal(choose.Choices[2].Block.Parent, &dialogue.Root)

	} else {
		assert.Fail("The third statement should be a ChooseStatement")
	}
}
