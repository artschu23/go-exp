package indeterminate

import "encoding/json"

var Identifier = "indeterminate"

type Field[T comparable] struct {
	Value           T
	IsIndeterminate bool
}

func (field Field[Types]) MarshalJSON() ([]byte, error) {
	if field.IsIndeterminate {
		return json.Marshal(Identifier)
	}

	return json.Marshal(field.Value)
}

func (field *Field[Types]) UnmarshalJSON(payload []byte) error {
	return json.Unmarshal(payload, &field.Value)
}
