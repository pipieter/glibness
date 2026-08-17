package main

/*
#include <stdint.h>
typedef uintptr_t Ptr;
*/
import "C"

import (
	"fmt"
	"glibness/glibness"
	"reflect"
	"runtime/cgo"
	"unsafe"
)

// Make sure this is the same as glib_response_code in glibness.h
const (
	GLIB_RESPONSE_NONE = iota
	GLIB_RESPONSE_ERROR
	GLIB_RESPONSE_SAY
	GLIB_RESPONSE_CHOICE
	GLIB_RESPONSE_INTERNAL_CHANGE
	GLIB_RESPONSE_FINISHED
)

// Wrapper struct around the glibness.Engine that keeps track of the last error.
// This way we don't need to store the error in the actual glibness.Engine
type CEngine struct {
	Engine *glibness.Engine
	Error  error
}

func (engine *CEngine) SetError(err error) C.char {
	engine.Error = err

	if err != nil {
		return C.char(0)
	} else {
		return C.char(1)
	}
}

//export glib_new_engine
func glib_new_engine() C.Ptr {
	engine := new(CEngine)
	engine.Engine = glibness.NewEngine()
	engine.Error = nil

	handle := cgo.NewHandle(engine)
	return C.Ptr(handle)
}

//export glib_free_engine
func glib_free_engine(engine C.Ptr) {
	handle := cgo.Handle(engine)
	handle.Delete()
}

//export glib_load_string
func glib_load_string(cEngine C.Ptr, cScript *C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)
	script := C.GoString(cScript)

	err := engine.Engine.ParseString(script)
	return engine.SetError(err)
}

//export glib_load_file
func glib_load_file(cEngine C.Ptr, cPath *C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)
	path := C.GoString(cPath)

	err := engine.Engine.ParseFile(path)
	return engine.SetError(err)
}

//export glib_start
func glib_start(cEngine C.Ptr, cDialogue *C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)
	dialogue := C.GoString(cDialogue)

	err := engine.Engine.Start(dialogue)
	return engine.SetError(err)
}

func getResponseCode(response glibness.StateResponse) (int, error) {
	switch response.(type) {

	case glibness.SayResponse:
		return GLIB_RESPONSE_SAY, nil

	case glibness.ChoiceResponse:
		return GLIB_RESPONSE_CHOICE, nil

	case glibness.InternalChangeResponse:
		return GLIB_RESPONSE_INTERNAL_CHANGE, nil

	case glibness.FinishedResponse:
		return GLIB_RESPONSE_FINISHED, nil
	}

	return GLIB_RESPONSE_ERROR, fmt.Errorf("Unsupported response type: '%s'", reflect.TypeOf(response))
}

//export glib_next
func glib_next(cEngine C.Ptr) C.int {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	response, err := engine.Engine.Next()
	if err == nil {
		code, err := getResponseCode(response)
		engine.SetError(err)
		return C.int(code)
	} else {
		engine.SetError(err)
		return C.int(GLIB_RESPONSE_ERROR)
	}
}

//export glib_choose
func glib_choose(cEngine C.Ptr, index C.int) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	err := engine.Engine.Respond(int(index))
	return engine.SetError(err)
}

func copyStringToCBuffer(str string, buffer *C.char) C.size_t {
	length := len(str)

	src := unsafe.StringData(str)

	dstBytes := unsafe.Slice((*byte)(unsafe.Pointer(buffer)), length)
	srcBytes := unsafe.Slice((*byte)(src), length)
	copy(dstBytes, srcBytes)

	// NUL-terminate so C strlen is valid
	*(*byte)(unsafe.Add(unsafe.Pointer(buffer), length)) = 0

	return C.size_t(length)

}

//export glib_get_sentence
func glib_get_sentence(cEngine C.Ptr, buffer *C.char) C.size_t {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	sentence := engine.Engine.Sentence()
	return copyStringToCBuffer(sentence, buffer)
}

//export glib_get_speaker
func glib_get_speaker(cEngine C.Ptr, buffer *C.char) C.size_t {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	speaker := engine.Engine.Speaker()
	return copyStringToCBuffer(speaker, buffer)
}

//export glib_get_error
func glib_get_error(cEngine C.Ptr, buffer *C.char) C.size_t {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	if engine.Error == nil {
		return C.size_t(0)
	}
	return copyStringToCBuffer(engine.Error.Error(), buffer)
}

//export glib_is_active
func glib_is_active(cEngine C.Ptr) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	if engine.Engine.State.Active {
		return C.char(1)
	} else {
		return C.char(0)
	}
}

func (engine CEngine) getVariable(key string) (glibness.Value, error) {
	value, ok := engine.Engine.Variables[key]

	if !ok {
		return nil, fmt.Errorf("Could not find variable %s", key)
	} else {
		return value, nil
	}
}

//export glib_get_str
func glib_get_str(cEngine C.Ptr, cKey *C.char, buffer *C.char) C.size_t {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	value, err := engine.getVariable(key)

	if err != nil {
		engine.SetError(err)
		return C.size_t(0)
	}

	if value, ok := value.(glibness.StringValue); ok {
		return copyStringToCBuffer(value.Value, buffer)
	} else {
		engine.SetError(fmt.Errorf("Requested a string for %s, but received %s instead.", key, value.Type()))
		return C.size_t(0)
	}
}

//export glib_get_int
func glib_get_int(cEngine C.Ptr, cKey *C.char, buffer *C.int) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	value, err := engine.getVariable(key)

	if err != nil {
		return engine.SetError(err)
	}

	if value, ok := value.(glibness.IntegerValue); ok {
		*buffer = C.int(value.Value)
		return C.char(1)
	} else {
		return engine.SetError(fmt.Errorf("Requested an integer for %s, but received %s instead.", key, value.Type()))
	}
}

//export glib_get_bool
func glib_get_bool(cEngine C.Ptr, cKey *C.char, buffer *C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	value, err := engine.getVariable(key)

	if err != nil {
		return engine.SetError(err)
	}

	if value, ok := value.(glibness.BooleanValue); ok {
		if value.Value {
			*buffer = C.char(1)
		} else {
			*buffer = C.char(0)
		}
		return C.char(1)
	} else {
		return engine.SetError(fmt.Errorf("Requested a boolean for %s, but received %s instead.", key, value.Type()))
	}
}

//export glib_set_str
func glib_set_str(cEngine C.Ptr, cKey *C.char, cValue *C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	value := C.GoString(cValue)

	err := engine.Engine.Set(key, glibness.MakeStringValue(value))
	return engine.SetError(err)
}

//export glib_set_int
func glib_set_int(cEngine C.Ptr, cKey *C.char, cValue C.int) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	value := int(cValue)

	err := engine.Engine.Set(key, glibness.MakeIntValue(value))
	return engine.SetError(err)
}

//export glib_set_bool
func glib_set_bool(cEngine C.Ptr, cKey *C.char, cValue C.char) C.char {
	handle := cgo.Handle(cEngine)
	engine, _ := handle.Value().(*CEngine)

	key := C.GoString(cKey)
	var value bool
	if cValue == 0 {
		value = false
	} else {
		value = true
	}

	err := engine.Engine.Set(key, glibness.MakeBooleanValue(value))
	return engine.SetError(err)
}

func main() {}
