package ccii

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/require"
)

func TestVisitor_feedScalar(t *testing.T) {
	var idest int
	var i8dest int8
	var udest uint
	var fdest float64
	var sdest string
	var invalidTarget struct{}
	psdest := &sdest
	tests := []struct {
		name    string
		input   string
		obj     interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "integer",
			input:   "-12",
			obj:     &idest,
			want:    -12,
			wantErr: false,
		},
		{
			name:    "int8",
			input:   "-12",
			obj:     &i8dest,
			want:    int8(-12),
			wantErr: false,
		},
		{
			name:    "unsigned",
			input:   "12",
			obj:     &udest,
			want:    uint(12),
			wantErr: false,
		},
		{
			name:    "float",
			input:   "4.5",
			obj:     &fdest,
			want:    float64(4.5),
			wantErr: false,
		},
		{
			name:    "string",
			input:   `12\3c34`,
			obj:     &sdest,
			want:    "12<34",
			wantErr: false,
		},
		{
			name:    "pointer to string",
			input:   `12\3c34`,
			obj:     psdest,
			want:    "12<34",
			wantErr: false,
		},
		{
			name:    "structure",
			input:   "4",
			obj:     &invalidTarget,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Visitor{
				obj: tt.obj,
			}
			if err := v.feedScalar(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("Visitor.feedScalar() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil && tt.wantErr {
				return
			}
			require.Equal(t, reflect.Ptr, reflect.TypeOf(tt.obj).Kind())
			deref := reflect.ValueOf(tt.obj).Elem().Interface()
			if reflect.TypeOf(deref).Kind() != reflect.Ptr {
				require.Equal(t, tt.want, deref, tt.name+" section")
			} else {
				require.Equal(t, tt.want, reflect.ValueOf(deref).Elem().Interface(), tt.name+" section")
			}
		})
	}
}
