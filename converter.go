package sjson

import (
	"github.com/soichisumi/go-simple-jsonmarshaler/internal/encoding/json"
)

// Unmarshal ...
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal ...
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decoder ...
type Decoder struct {
	json.Decoder
}

// Encoder ...
type Encoder struct {
	json.Encoder
}