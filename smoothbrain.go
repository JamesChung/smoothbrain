package smoothbrain

import (
	"encoding/json"
	"fmt"
)

// SmoothBrain is an instance of a smooth mapping.
type SmoothBrain struct {
	smoothMap   map[string]any
	startingMap map[string]any
	marshal     func(any) ([]byte, error)
	unmarshal   func([]byte, any) error
}

// Unmarshal takes json bytes and will return an error if one occurs.
func (s *SmoothBrain) Unmarshal(data []byte) error {
	err := s.unmarshal(data, &s.startingMap)
	if err != nil {
		return err
	}
	flattenMap("", s.startingMap, s.smoothMap)
	return nil
}

// Marshal returns the smooth json bytes or an error if one occurs.
func (s *SmoothBrain) Marshal() ([]byte, error) {
	data, err := s.marshal(&s.smoothMap)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Get returns the smooth mapping of unmarshalled data.
func (s *SmoothBrain) Get() map[string]any {
	return s.smoothMap
}

// New returns a new instance of a SmoothBrain.
func New() *SmoothBrain {
	return &SmoothBrain{
		marshal:     json.Marshal,
		unmarshal:   json.Unmarshal,
		startingMap: make(map[string]any),
		smoothMap:   make(map[string]any),
	}
}

func handleArray(prefix, key string, arr []any, dest map[string]any) {
	for i, item := range arr {
		switch val := item.(type) {
		case map[string]any:
			// If it's a nested map, recurse
			newPrefix := prefix + key + "." + fmt.Sprintf("%d", i)
			dest[newPrefix] = item
			flattenMap(newPrefix+".", val, dest)
		case []any:
			// If it's an array again, recurse
			handleArray(prefix, key, val, dest)
		}
	}
}

func flattenMap(prefix string, src map[string]any, dest map[string]any) {
	for key, value := range src {
		switch val := value.(type) {
		case map[string]any:
			dest[prefix+key] = value
			// If it's a nested map, recurse
			newPrefix := prefix + key + "."
			flattenMap(newPrefix, val, dest)
		case []any:
			dest[prefix+key] = value
			// If it's an array, we'll handle it based on preference
			handleArray(prefix, key, val, dest)
		default:
			// Otherwise, store directly
			dest[prefix+key] = value
		}
	}
}
