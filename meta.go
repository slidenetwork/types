package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// Meta @TODO
type Meta map[string]interface{}

// Add @TODO
func (m Meta) Add(key string, val interface{}) {
	m.Set(key, val)
}

// Set @TODO
func (m Meta) Set(key string, val interface{}) {
	m[key] = val
}

// Get returns the value from meta
func (m Meta) Get(key string) interface{} {
	if m == nil {
		return nil
	}
	v, ok := m[key]
	if !ok {
		return nil
	}
	return v
}

// GetBytes @DOC
func (m Meta) GetBytes(key string) ([]byte, error) {
	b, err := toBytes(m.Get(key))
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetString @DOC
func (m Meta) GetString(key string) string {
	v := m.Get(key)
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

// GetStruct handles unmarshaling whatever value is stored in a Meta key into
// the given interface argument
func (m Meta) GetStruct(key string, s interface{}) error {
	v := m.Get(key)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}

func toBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
