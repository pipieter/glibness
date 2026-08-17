package glibness

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Capture errors instead of printing them
type ErrorCapture struct {
	antlr.DefaultErrorListener
	Errors []string
}

func (capture *ErrorCapture) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	capture.Errors = append(capture.Errors, fmt.Sprintf("%d:%d %s", line, column, msg))
}

func (capture *ErrorCapture) Error() error {
	if capture.Errors == nil {
		return nil
	}
	if len(capture.Errors) == 0 {
		return nil
	}
	return fmt.Errorf("%s", strings.Join(capture.Errors, "; "))
}
