package sjson

import (
	"encoding/json"
	"fmt"
)

func ExampleUnmarshal() {
	type Bytes struct {
		A []byte
	}

	in := `{
	"a": "not-base64-encoded-string"
}`

	var v Bytes
	// json.Unmarshal:
	if err := json.Unmarshal([]byte(in), &v); err != nil {
		fmt.Printf("json: error: %+v\n", err)
	} else {
		fmt.Printf("json: res: %+v\n", string(v.A))
	}

	// sjson.Unmarshal:
	if err := Unmarshal([]byte(in), &v); err != nil {
		fmt.Printf("sjson: error: %+v\n", err)
	} else {
		fmt.Printf("sjson: res: %+v\n", string(v.A))
	}
	// Output:
	// json: error: illegal base64 data at input byte 3
	// sjson: res: not-base64-encoded-string
}

func ExampleMarshal() {
	in := struct{
		A []byte
	}{A: []byte("an string")}

	// json.Marshal:
	b, _ := json.Marshal(in)
	fmt.Printf("json.Marshal: res: %+v\n", string(b))

	// sjson.Marshal:
	b, _ = Marshal(in);
	fmt.Printf("sjson.Unmarshal: res: %+v\n", string(b))

	// Output:
	// json.Marshal: res: {"A":"YW4gc3RyaW5n"}
	// sjson.Unmarshal: res: {"A":"an string"}
}