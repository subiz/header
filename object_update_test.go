package header

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestObjectUpdateRoundTrip(t *testing.T) {
	values := []any{"Tiêu đề mới", float64(0), false, nil, map[string]any{"name": "An"}, []any{"a", float64(2)}}
	update := &ObjectUpdate{ObjectType: "conversation", ObjectId: "cs_demo", Unsets: []string{"analysis", "fields.0"}}
	for _, value := range values {
		v, err := structpb.NewValue(value)
		if err != nil {
			t.Fatal(err)
		}
		update.Sets = append(update.Sets, &ObjectFieldSet{Path: `members.#(id=="us_demo").name`, Value: v})
	}
	event := &Event{Type: "conversation_updated_2", Data: &Data{ObjectUpdate: update}}
	for _, codec := range []struct {
		name      string
		marshal   func(*Event) ([]byte, error)
		unmarshal func([]byte, *Event) error
	}{
		{"protobuf", func(e *Event) ([]byte, error) { return proto.Marshal(e) }, func(b []byte, e *Event) error { return proto.Unmarshal(b, e) }},
		{"protojson", func(e *Event) ([]byte, error) { return (protojson.MarshalOptions{UseProtoNames: true}).Marshal(e) }, func(b []byte, e *Event) error { return protojson.Unmarshal(b, e) }},
		{"json", func(e *Event) ([]byte, error) { return json.Marshal(e) }, func(b []byte, e *Event) error { return json.Unmarshal(b, e) }},
	} {
		t.Run(codec.name, func(t *testing.T) {
			b, err := codec.marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			var got Event
			if err := codec.unmarshal(b, &got); err != nil {
				t.Fatal(err)
			}
			if !proto.Equal(event, &got) {
				t.Fatalf("object update did not survive %s round trip", codec.name)
			}
		})
	}
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	var wire struct {
		Data struct {
			Update struct {
				Type   string                       `json:"object_type"`
				ID     string                       `json:"object_id"`
				Sets   []map[string]json.RawMessage `json:"sets"`
				Unsets []string                     `json:"unsets"`
			} `json:"object_update"`
		} `json:"data"`
	}
	if err := json.Unmarshal(b, &wire); err != nil {
		t.Fatal(err)
	}
	if wire.Data.Update.Type != "conversation" || wire.Data.Update.ID != "cs_demo" || len(wire.Data.Update.Unsets) != 2 || len(wire.Data.Update.Sets) != len(values) {
		t.Fatalf("unexpected JSON shape: %s", b)
	}
	for i, expected := range []string{`"Tiêu đề mới"`, "0", "false", "null", `{"name":"An"}`, `["a",2]`} {
		if string(wire.Data.Update.Sets[i]["value"]) != expected {
			t.Fatalf("value %d = %s, want %s", i, wire.Data.Update.Sets[i]["value"], expected)
		}
	}
}

func TestObjectFieldSetNullAndMissingValue(t *testing.T) {
	for _, tc := range []struct {
		wire    string
		present bool
	}{
		{`{"path":"analysis","value":null}`, true},
		{`{"path":"analysis"}`, false},
	} {
		var set ObjectFieldSet
		if err := json.Unmarshal([]byte(tc.wire), &set); err != nil {
			t.Fatal(err)
		}
		if (set.Value != nil) != tc.present {
			t.Fatalf("lost value presence in %s", tc.wire)
		}
		b, err := json.Marshal(&set)
		if err != nil || string(b) != tc.wire {
			t.Fatalf("got %s (%v), want %s", b, err, tc.wire)
		}
	}
}
