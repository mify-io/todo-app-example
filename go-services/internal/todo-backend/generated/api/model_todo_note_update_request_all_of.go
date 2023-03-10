// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi

type TodoNoteUpdateRequestAllOf struct {
	IsCompleted bool `json:"is_completed,omitempty"`
}

// AssertTodoNoteUpdateRequestAllOfRequired checks if the required fields are not zero-ed
func AssertTodoNoteUpdateRequestAllOfRequired(obj TodoNoteUpdateRequestAllOf) error {
	return nil
}

// AssertRecurseTodoNoteUpdateRequestAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TodoNoteUpdateRequestAllOf (e.g. [][]TodoNoteUpdateRequestAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTodoNoteUpdateRequestAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTodoNoteUpdateRequestAllOf, ok := obj.(TodoNoteUpdateRequestAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTodoNoteUpdateRequestAllOfRequired(aTodoNoteUpdateRequestAllOf)
	})
}
