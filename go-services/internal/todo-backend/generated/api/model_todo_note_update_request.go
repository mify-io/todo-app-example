// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi

type TodoNoteUpdateRequest struct {
	Description string `json:"description,omitempty"`

	Title string `json:"title"`

	IsCompleted bool `json:"is_completed,omitempty"`
}

// AssertTodoNoteUpdateRequestRequired checks if the required fields are not zero-ed
func AssertTodoNoteUpdateRequestRequired(obj TodoNoteUpdateRequest) error {
	elements := map[string]interface{}{
		"title": obj.Title,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTodoNoteUpdateRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TodoNoteUpdateRequest (e.g. [][]TodoNoteUpdateRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTodoNoteUpdateRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTodoNoteUpdateRequest, ok := obj.(TodoNoteUpdateRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTodoNoteUpdateRequestRequired(aTodoNoteUpdateRequest)
	})
}
