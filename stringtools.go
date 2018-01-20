package ccii

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const serviceChars = " ={}[]\\,:$\"'`&<>"

// Encode экранирует запрещённые символы в файле
func Encode(input string) string {
	output := bytes.Buffer{}
	for len(input) > 0 {
		pos := strings.IndexAny(input, serviceChars)
		if pos < 0 {
			output.WriteString(input)
			input = ""
			break
		}
		output.WriteString(input[:pos])
		char := input[pos]
		input = input[pos+1:]

		output.WriteByte('\\')
		output.WriteString(strconv.FormatUint(uint64(char), 16))
	}
	return output.String()
}

// Decode совершает обратное к Encode преобразование
func Decode(input string) (string, error) {
	original := input
	output := bytes.Buffer{}
	for len(input) > 0 {
		pos := strings.IndexByte(input, '\\')
		if pos < 0 {
			output.WriteString(input)
			input = ""
			break
		}
		output.WriteString(input[:pos])
		input = input[pos+1:]
		if len(input) < 2 {
			return "", fmt.Errorf("cannot decode `%s`: `%s` is not valid 1-byte hex code", original, input)
		}
		code := input[:2]
		input = input[2:]
		res, err := strconv.ParseUint(code, 16, 8)
		if err != nil {
			return "", fmt.Errorf("cannot decode `%s`: `%s` is not valid 1-byte hex code", original, code)
		}
		output.WriteByte(byte(res))
	}
	return output.String(), nil
}
