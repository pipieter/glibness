package test

import (
	"glibness/glibness"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestChoiceTree(t *testing.T) {
	var err error
	var response glibness.StateResponse

	assert := assert.New(t)

	engine := glibness.NewEngine()
	err = engine.ParseString(`
		dialogue test {
			set speaker "Test"
			say "This has a choice message!"
			say "Choices can be made!"

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
	assert.Len(engine.Dialogues, 1)

	err = engine.Start("test")
	assert.Nil(err)

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.SayResponse{})

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.SayResponse{})

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.ChoiceResponse{})

	if choice, ok := response.(glibness.ChoiceResponse); ok {
		assert.ElementsMatch([]string{"A", "B", "C"}, choice.Choices)
	} else {
		assert.Fail("Response should be a choice response")
	}

	err = engine.Respond(0)
	assert.Nil(err)

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.SayResponse{})

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.SayResponse{})

	response, err = engine.Next()
	assert.Nil(err)
	assert.IsType(response, glibness.FinishedResponse{})
}

func TestInvalidChoiceTree(t *testing.T) {
	var response glibness.StateResponse
	var err error

	assert := assert.New(t)
	engine := glibness.NewEngine()
	err = engine.ParseString(`
		dialogue test {
			set speaker "Test"
			say "This has a choice message, which will be answered with a wrong index!"

			choose {
				choice "A" {
					say "A.1"
				}
				choice "B" {
					say "B.1"
				}
 			}
		}
	`)

	assert.Nil(err)
	assert.Len(engine.Dialogues, 1)

	dialogue := engine.Dialogues[0]
	assert.Equal(dialogue.Name, "test")

	err = engine.Start("test")
	assert.Nil(err)

	response, err = engine.Next()
	assert.IsType(response, glibness.SayResponse{})
	assert.Nil(err)

	response, err = engine.Next()
	assert.IsType(response, glibness.ChoiceResponse{})
	assert.Nil(err)

	// Make a wrong choice here, on purpose
	err = engine.Respond(999)
	assert.NotNil(err)

	// Make a right choice here
	err = engine.Respond(1)
	assert.Nil(err)

	// Continue with the script
	response, err = engine.Next()
	assert.IsType(response, glibness.SayResponse{})
	assert.Nil(err)

	response, err = engine.Next()
	assert.IsType(response, glibness.FinishedResponse{})
	assert.Nil(err)
}
