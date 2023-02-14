// THIS FILE IS AUTOGENERATED, DO NOT EDIT
// Generated by mify via OpenAPI Generator

// vim: set ft=go:
package openapi

type TodoNoteAllOf struct {
	CreatedAt string `json:"created_at"`

	Id int64 `json:"id"`

	UpdatedAt string `json:"updated_at"`
}

// AssertTodoNoteAllOfRequired checks if the required fields are not zero-ed
func AssertTodoNoteAllOfRequired(obj TodoNoteAllOf) error {
	elements := map[string]interface{}{
		"created_at": obj.CreatedAt,
		"id":         obj.Id,
		"updated_at": obj.UpdatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTodoNoteAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TodoNoteAllOf (e.g. [][]TodoNoteAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTodoNoteAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTodoNoteAllOf, ok := obj.(TodoNoteAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTodoNoteAllOfRequired(aTodoNoteAllOf)
	})
}
