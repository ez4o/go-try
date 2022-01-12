package try

import (
	"reflect"
)

func (e *Exception) Catch(f any) *Exception {
	if e == nil {
		return nil
	}

	fVal := reflect.ValueOf(f)
	fType := fVal.Type()

	if fType.NumIn() == 1 {
		firstParamType := fType.In(0)
		errorType := reflect.TypeOf(e.err)

		if errorType.String() == "*errors.errorString" {
			var err error
			errorType = reflect.TypeOf(&err).Elem()
		}

		if firstParamType.AssignableTo(errorType) {
			fVal.Call([]reflect.Value{reflect.ValueOf(e.err)})
		}
	} else if fType.NumIn() == 2 {
		firstParamType := fType.In(0)
		errorType := reflect.TypeOf(e.err)

		secondParamType := fType.In(1)
		stackTraceType := reflect.TypeOf(e.stackTrace)

		if errorType.String() == "*errors.errorString" {
			var err error
			errorType = reflect.TypeOf(&err).Elem()
		}

		if firstParamType.AssignableTo(errorType) && secondParamType.ConvertibleTo(stackTraceType) {
			fVal.Call([]reflect.Value{reflect.ValueOf(e.err), reflect.ValueOf(e.stackTrace)})
		}
	}

	return e
}
