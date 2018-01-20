package ccii

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		obj     interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "simple",
			obj:     12,
			want:    "12",
			wantErr: false,
		},
		{
			name:    "array",
			obj:     []int{1, 2, 3},
			want:    `[1,2,3]`,
			wantErr: false,
		},
		{
			name: "map trivial",
			obj: map[int][]string{
				1: {"1", "ab"},
			},
			want:    "{1[1,ab]}",
			wantErr: false,
		},
		{
			name: "struct",
			obj: struct {
				A int
				B struct {
					C string
					D int
				}
			}{
				A: 1,
				B: struct {
					C string
					D int
				}{
					C: "ab",
					D: 2,
				},
			},
			want:    "{A=1,B{C=ab,D=2}}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}

			// Сейчас проверяем, что
			dest := reflect.New(reflect.TypeOf(tt.obj))
			if err := Unmarshal(tt.want, dest.Interface()); err != nil {
				t.Fatal(err)
			}
			require.Equal(t, tt.obj, dest.Elem().Interface())
		})
	}
}

func TestMarshalNonTrivialMap(t *testing.T) {
	source := map[string]string{
		"a": "а",
		"b": "б",
		"c": "ц",
	}
	var dest map[string]string
	inter, err := Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := Unmarshal(inter, &dest); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, source, dest)
}

func TestMarshalEmbeddedStruct(t *testing.T) {
	type (
		T1 struct {
			A string
			B int
		}
		T2 struct {
			T1
			C T1
			D string
		}
	)

	t1 := T1{
		A: "field",
		B: 12,
	}
	source := T2{
		T1: t1,
		C:  t1,
		D:  "pole",
	}
	inter, err := Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, `{A=field,B=12,C{A=field,B=12},D=pole}`, inter)

	var dest T2
	if err := Unmarshal(inter, &dest); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, source, dest)
}
