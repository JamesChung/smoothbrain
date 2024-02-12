package smoothbrain

import (
	"encoding/json"
	"fmt"
)

type SmoothBrain struct {
	smoothMap   map[string]any
	startingMap map[string]any
	marshal     func(any) ([]byte, error)
	unmarshal   func([]byte, any) error
}

func (s *SmoothBrain) Unmarshal(data []byte) error {
	err := s.unmarshal(data, &s.startingMap)
	if err != nil {
		return err
	}
	flattenMap("", s.startingMap, s.smoothMap)
	return nil
}

func (s *SmoothBrain) Marshal() ([]byte, error) {
	data, err := s.marshal(&s.smoothMap)
	if err != nil {
		return nil, err
	}
	return data, nil
}

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
