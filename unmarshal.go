package ccii

import (
	"github.com/sirkon/ccii/parser"

	"fmt"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Unmarshal преобразует компрессированное текстовое представление в объект
func Unmarshal(data string, dest interface{}) (err error) {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("pointer expected, got %T", dest)
	}
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()
	input := antlr.NewInputStream(data)
	lexer := parser.NewCCIILexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCCIIParser(stream)
	p.RemoveErrorListeners()
	tree := p.Main()

	v := Visitor{
		obj: dest,
	}
	v.Visit(tree)

	return nil
}
