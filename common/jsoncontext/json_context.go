package jsoncontext

import (
	"encoding/json"

	"github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/entity"
)

// JsonContext provides JSON serialization utilities for our types
type JsonContext struct{}

// NewJsonContext creates a new JsonContext
func NewJsonContext() *JsonContext {
	return &JsonContext{}
}

// Marshal marshals an interface to JSON with consistent formatting
func (jc *JsonContext) Marshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// Unmarshal unmarshals JSON data into an interface
func (jc *JsonContext) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Utility functions for specific types
// These assume the types are defined elsewhere in your codebase

func (jc *JsonContext) MarshalStreamSpecs(specs []entity.StreamSpec) ([]byte, error) {
	return jc.Marshal(specs)
}

func (jc *JsonContext) UnmarshalStreamSpecs(data []byte) ([]entity.StreamSpec, error) {
	var specs []entity.StreamSpec
	err := jc.Unmarshal(data, &specs)
	return specs, err
}

func (jc *JsonContext) MarshalMediaSegments(segments []entity.MediaSegment) ([]byte, error) {
	return jc.Marshal(segments)
}

func (jc *JsonContext) UnmarshalMediaSegments(data []byte) ([]entity.MediaSegment, error) {
	var segments []entity.MediaSegment
	err := jc.Unmarshal(data, &segments)
	return segments, err
}

func (jc *JsonContext) MarshalStringDict(dict map[string]string) ([]byte, error) {
	return jc.Marshal(dict)
}

func (jc *JsonContext) UnmarshalStringDict(data []byte) (map[string]string, error) {
	var dict map[string]string
	err := jc.Unmarshal(data, &dict)
	return dict, err
}
