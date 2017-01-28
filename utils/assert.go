package utils

import (
	"reflect"
	"testing"
)

type Assert struct {
	t *testing.T
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

func errorSingle(t *testing.T, msg string, obj interface{}) {
	t.Errorf("%s: %v", msg, obj)
}

func errorCompare(t *testing.T, msg string, expected, actual interface{}) {
	t.Errorf("%s:\n\texpected: %v\n\tactual: %v", msg, expected, actual)
}

func (a *Assert) Equal(expected, actual interface{}, msg string) {
	if !ObjectsAreEqual(expected, actual) {
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

func NewAssert(t *testing.T) *Assert {
	return &Assert{t}
}