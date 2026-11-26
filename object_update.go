package header

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"
)

// UnmarshalJSON preserves an explicit JSON null as a protobuf NullValue.
// encoding/json otherwise turns the *Value into nil and loses the distinction
// between {"value":null} and an omitted value when the update is forwarded.
func (s *ObjectFieldSet) UnmarshalJSON(data []byte) error {
	var wire struct {
		Path  string          `json:"path"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	var value *structpb.Value
	if len(wire.Value) > 0 {
		value = new(structpb.Value)
		if err := value.UnmarshalJSON(wire.Value); err != nil {
			return err
		}
	}
	s.Reset()
	s.Path, s.Value = wire.Path, value
	return nil
}
