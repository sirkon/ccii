package ccii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "void encoding",
			input: "123 321",
			want:  "123\\20321",
		},
		{
			name:  "simple encoding",
			input: "={}[]\\,:$",
			want:  `\3d\7b\7d\5b\5d\5c\2c\3a\24`,
		},
		{
			name:  "common",
			input: "12=13",
			want:  `12\3d13`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.input); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			} else {
				back, err := Decode(tt.want)
				if err != nil {
					t.Errorf("Can't decode `%s` back to `%s`: %s", tt.want, tt.input, err)
				}
				require.Equal(t, tt.input, back)
			}
		})
	}
}
