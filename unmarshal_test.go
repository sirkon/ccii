package ccii

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalMap(t *testing.T) {
	var dest map[string]int
	if err := Unmarshal(`{A=1,B=2}`, &dest); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, map[string]int{
		"A": 1,
		"B": 2,
	}, dest)
}

func TestUnmarshalStruct(t *testing.T) {
	type SimpleType struct {
		A int
		B string
	}
	type harderType struct {
		SimpleType
		C string
		D []int
	}
	tests := []struct {
		name    string
		data    string
		dest    interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name: "simple",
			data: "{A=1,B=a}",
			dest: SimpleType{},
			want: SimpleType{
				A: 1,
				B: "a",
			},
			wantErr: false,
		},
		{
			name: "harder",
			data: "{A=1,B=a,C=12,D[1,2,3]}",
			dest: harderType{},
			want: harderType{
				SimpleType: SimpleType{
					A: 1,
					B: "a",
				},
				C: "12",
				D: []int{1, 2, 3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dest := reflect.New(reflect.TypeOf(tt.dest)).Interface()
			if err := Unmarshal(tt.data, dest); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil && tt.wantErr {
				return
			}
			require.Equal(t, tt.want, reflect.ValueOf(dest).Elem().Interface(), "section "+tt.name)
		})
	}
}

func TestUnmarshalArray(t *testing.T) {
	var arrint []int
	var arrstring []string
	tests := []struct {
		name    string
		input   string
		dest    interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "strings",
			input:   "[a,bc,def,ghij,klmno]",
			dest:    arrstring,
			want:    []string{"a", "bc", "def", "ghij", "klmno"},
			wantErr: false,
		},
		{
			name:    "ints",
			input:   `[1,2,3,4,5,7]`,
			dest:    arrint,
			want:    []int{1, 2, 3, 4, 5, 7},
			wantErr: false,
		},
		{
			name:    "invalid structs",
			input:   `[{},{}]`,
			dest:    arrstring,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dest := reflect.New(reflect.TypeOf(tt.dest))
			if err := Unmarshal(tt.input, dest.Interface()); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil && tt.wantErr {
				return
			}
			require.Equal(t, tt.want, dest.Elem().Interface(), "section "+tt.name)
		})
	}
}
