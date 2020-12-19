package sjson

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type InnerStruct struct {
	Tag string
}

type TestStruct struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr
	Float32 float32
	Float64 float64

	IntStr int64 `json:",string"`

	PBool    *bool
	PInt     *int
	PInt8    *int8
	PInt16   *int16
	PInt32   *int32
	PInt64   *int64
	PUint    *uint
	PUint8   *uint8
	PUint16  *uint16
	PUint32  *uint32
	PUint64  *uint64
	PUintptr *uintptr
	PFloat32 *float32
	PFloat64 *float64

	String  string
	PString *string

	Map   map[string]InnerStruct
	MapP  map[string]*InnerStruct
	PMap  *map[string]InnerStruct
	PMapP *map[string]*InnerStruct

	EmptyMap map[string]InnerStruct
	NilMap   map[string]InnerStruct

	Slice   []InnerStruct
	SliceP  []*InnerStruct
	PSlice  *[]InnerStruct
	PSliceP *[]*InnerStruct

	EmptySlice []InnerStruct
	NilSlice   []InnerStruct

	StringSlice []string
	ByteSlice   []byte

	InnerStruct   InnerStruct
	PInnerStruct  *InnerStruct
	PPInnerStruct **InnerStruct

	Interface  interface{}
	PInterface *interface{}
}

var allValueIndent = `{
	"Bool": true,
	"Int": 2,
	"Int8": 3,
	"Int16": 4,
	"Int32": 5,
	"Int64": 6,
	"Uint": 7,
	"Uint8": 8,
	"Uint16": 9,
	"Uint32": 10,
	"Uint64": 11,
	"Uintptr": 12,
	"Float32": 14.1,
	"Float64": 15.1,
	"IntStr": "42",
	"PBool": null,
	"PInt": null,
	"PInt8": null,
	"PInt16": null,
	"PInt32": null,
	"PInt64": null,
	"PUint": null,
	"PUint8": null,
	"PUint16": null,
	"PUint32": null,
	"PUint64": null,
	"PUintptr": null,
	"PFloat32": null,
	"PFloat64": null,
	"String": "16",
	"PString": null,
	"Map": {
		"17": {
			"Tag": "tag17"
		},
		"18": {
			"Tag": "tag18"
		}
	},
	"MapP": {
		"19": {
			"Tag": "tag19"
		},
		"20": null
	},
	"PMap": null,
	"PMapP": null,
	"EmptyMap": {},
	"NilMap": null,
	"Slice": [
		{
			"Tag": "tag20"
		},
		{
			"Tag": "tag21"
		}
	],
	"SliceP": [
		{
			"Tag": "tag22"
		},
		null,
		{
			"Tag": "tag23"
		}
	],
	"PSlice": null,
	"PSliceP": null,
	"EmptySlice": [],
	"NilSlice": null,
	"StringSlice": [
		"str24",
		"str25",
		"str26"
	],
	"ByteSlice": "Gxwd",
	"InnerStruct": {
		"Tag": "tag30"
	},
	"PInnerStruct": {
		"Tag": "tag31"
	},
	"PPInnerStruct": null,
	"Interface": 5.2,
	"PInterface": null
}`

var allValueCompact = strings.Map(noSpace, allValueIndent)

func TestMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success",
			args: args{v: TestStruct{
				Bool:    true,
				Int:     2,
				Int8:    3,
				Int16:   4,
				Int32:   5,
				Int64:   6,
				Uint:    7,
				Uint8:   8,
				Uint16:  9,
				Uint32:  10,
				Uint64:  11,
				Uintptr: 12,
				Float32: 14.1,
				Float64: 15.1,
				IntStr:  42,
				String:  "16",
				Map: map[string]InnerStruct{
					"17": {Tag: "tag17"},
					"18": {Tag: "tag18"},
				},
				MapP: map[string]*InnerStruct{
					"19": {Tag: "tag19"},
					"20": nil,
				},
				EmptyMap:     map[string]InnerStruct{},
				Slice:        []InnerStruct{{Tag: "tag20"}, {Tag: "tag21"}},
				SliceP:       []*InnerStruct{{Tag: "tag22"}, nil, {Tag: "tag23"}},
				EmptySlice:   []InnerStruct{},
				StringSlice:  []string{"str24", "str25", "str26"},
				ByteSlice:    []byte("Gxwd"),
				InnerStruct:  InnerStruct{Tag: "tag30"},
				PInnerStruct: &InnerStruct{Tag: "tag31"},
				Interface:    5.2,
			}},
			want:    []byte(allValueCompact),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v,\n wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    TestStruct
		wantErr bool
	}{
		{
			name: "",
			args: args{
				data: []byte(allValueIndent),
			},
			want: TestStruct{
				Bool:    true,
				Int:     2,
				Int8:    3,
				Int16:   4,
				Int32:   5,
				Int64:   6,
				Uint:    7,
				Uint8:   8,
				Uint16:  9,
				Uint32:  10,
				Uint64:  11,
				Uintptr: 12,
				Float32: 14.1,
				Float64: 15.1,
				IntStr:  42,
				String:  "16",
				Map: map[string]InnerStruct{
					"17": {Tag: "tag17"},
					"18": {Tag: "tag18"},
				},
				MapP: map[string]*InnerStruct{
					"19": {Tag: "tag19"},
					"20": nil,
				},
				EmptyMap:     map[string]InnerStruct{},
				Slice:        []InnerStruct{{Tag: "tag20"}, {Tag: "tag21"}},
				SliceP:       []*InnerStruct{{Tag: "tag22"}, nil, {Tag: "tag23"}},
				EmptySlice:   []InnerStruct{},
				StringSlice:  []string{"str24", "str25", "str26"},
				ByteSlice:    []byte("Gxwd"),
				InnerStruct:  InnerStruct{Tag: "tag30"},
				PInnerStruct: &InnerStruct{Tag: "tag31"},
				Interface:    5.2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v TestStruct
			if err := Unmarshal(tt.args.data, &v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v,\n wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(v, tt.want); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func noSpace(c rune) rune {
	if isSpace(byte(c)) { //only used for ascii
		return -1
	}
	return c
}
