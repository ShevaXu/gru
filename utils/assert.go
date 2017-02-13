package utils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Inspired by github.com/stretchr/testify/assert
// & https://github.com/benbjohnson/testing

type Assert struct {
	t testing.TB
}

func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	return reflect.DeepEqual(expected, actual)
}

func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func GetAddress(v interface{}) string {
	return fmt.Sprintf("%p", v)
}

func errorSingle(t testing.TB, msg string, obj interface{}) {
	//t.Errorf("%s: %v", msg, obj)
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("\033[31m\t%s:%d: %s\n\n\t\t%#v\033[39m\n\n", filepath.Base(file), line, msg, obj)
	t.Fail()
}

func errorCompare(t testing.TB, msg string, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("\033[31m\t%s:%d: %s\n\n\t\tgot: %#v\n\033[32m\t\texp: %#v\033[39m\n\n", filepath.Base(file), line, msg, actual, expected)
	t.Fail()
}

func (a *Assert) True(cond bool, msg string) {
	if !cond {
		errorSingle(a.t, msg, nil)
	}
}

func (a *Assert) Equal(expected, actual interface{}, msg string) {
	if !ObjectsAreEqual(expected, actual) {
		errorCompare(a.t, msg, expected, actual)
	}
}

func (a *Assert) NotEqual(expected, actual interface{}, msg string) {
	if ObjectsAreEqual(expected, actual) {
		errorCompare(a.t, msg, expected, actual)
	}
}

func (a *Assert) NoError(err error, msg string) {
	if err != nil {
		errorSingle(a.t, msg, err)
	}
}

func (a *Assert) Nil(obj interface{}, msg string) {
	if !IsNil(obj) {
		errorSingle(a.t, msg, obj)
	}
}

func (a *Assert) NotNil(obj interface{}, msg string) {
	if IsNil(obj) {
		errorSingle(a.t, msg, obj)
	}
}

func NewAssert(t testing.TB) *Assert {
	return &Assert{t}
}
