package ccii

import (
	"strconv"

	"fmt"

	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ccii/parser"
)

// Visitor для разбора структуры
type Visitor struct {
	obj interface{}
}

// NewVisitor конструктор с объектом назначения
func NewVisitor(obj interface{}) (*Visitor, error) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("pointer to a type expected, got %T instead", obj)
	}

	return &Visitor{
		obj: obj,
	}, nil
}

func (v *Visitor) getObject() interface{} {
	return v.obj
}

func (v *Visitor) withObject(obj interface{}) *Visitor {
	return &Visitor{
		obj: obj,
	}
}

// Visit ...
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitChildren ...
func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

// VisitTerminal ...
func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	if err := v.feedScalar(node.GetText()); err != nil {
		panic(err)
	}
	return nil
}

// VisitErrorNode ...
func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node.Accept(v)
}

func (v *Visitor) feedScalar(value string) (err error) {
	value, err = Decode(value)
	if err != nil {
		return err
	}

	switch val := v.obj.(type) {
	case *int:
		intval, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		*val = int(intval)
	case *int8:
		int8val, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return err
		}
		*val = int8(int8val)
	case *int16:
		int16val, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		*val = int16(int16val)
	case *int32:
		int32val, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		*val = int32(int32val)
	case *int64:
		int64val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		*val = int64(int64val)
	case *uint:
		uintval, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		*val = uint(uintval)
	case *uint8:
		uint8val, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		*val = uint8(uint8val)
	case *uint16:
		uint16val, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		*val = uint16(uint16val)
	case *uint32:
		uint32val, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		*val = uint32(uint32val)
	case *uint64:
		uint64val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		*val = uint64(uint64val)
	case *float32:
		float32val, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		*val = float32(float32val)
	case *float64:
		float64val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*val = float64(float64val)
	case *string:
		*val = value

	case **int:
		intval, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		k := int(intval)
		*val = &k
	case **int8:
		int8val, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return err
		}
		k := int8(int8val)
		*val = &k
	case **int16:
		int16val, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		k := int16(int16val)
		*val = &k
	case **int32:
		int32val, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		k := int32(int32val)
		*val = &k
	case **int64:
		int64val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		k := int64(int64val)
		*val = &k
	case **uint:
		uintval, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		k := uint(uintval)
		*val = &k
	case **uint8:
		uint8val, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		k := uint8(uint8val)
		*val = &k
	case **uint16:
		uint16val, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		k := uint16(uint16val)
		*val = &k
	case **uint32:
		uint32val, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		k := uint32(uint32val)
		*val = &k
	case **uint64:
		uint64val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		k := uint64(uint64val)
		*val = &k
	case **float32:
		float32val, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		k := float32(float32val)
		*val = &k
	case **float64:
		float64val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		k := float64(float64val)
		*val = &k
	case **string:
		*val = &value
	default:
		return fmt.Errorf("cannot put a scalar into %T", v.obj)
	}
	return nil
}

// VisitMain ...
func (v *Visitor) VisitMain(ctx *parser.MainContext) interface{} {
	switch {
	case ctx.SCALAR() != nil:
		return ctx.SCALAR().Accept(v)
	case ctx.Array() != nil:
		return ctx.Array().Accept(v)
	case ctx.Object() != nil:
		return ctx.Object().Accept(v)
	}
	return nil
}

type specerr string

// Error ...
func (e specerr) Error() string {
	return string(e)
}

// VisitObject ...
func (v *Visitor) VisitObject(ctx *parser.ObjectContext) interface{} {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(specerr); ok {
				panic(specerr(fmt.Sprintf("cannot unmarshal object into %T", v.obj)))
			}
			panic(r)
		}
	}()
	if ctx.Object_intern() == nil {
		return nil
	}
	valueType := reflect.TypeOf(v.obj).Elem()
	switch valueType.Kind() {
	case reflect.Struct, reflect.Map:
		return ctx.Object_intern().Accept(v)
	case reflect.Ptr:
		newItem := reflect.New(valueType.Elem())
		res := v.withObject(newItem.Int()).VisitObject(ctx)
		reflect.ValueOf(v.obj).Set(newItem)
		return res
	default:
		panic(specerr(fmt.Sprintf("cannot unmarshal object into %T", v.obj)))
	}
}

type acceptor interface {
	Accept(v antlr.ParseTreeVisitor) interface{}
}

func (v *Visitor) set(name string, value acceptor) interface{} {
	switch reflect.TypeOf(v.obj).Elem().Kind() {
	case reflect.Map:
		keyType := reflect.TypeOf(v.obj).Elem().Key()
		valueType := reflect.TypeOf(v.obj).Elem().Elem()
		if reflect.ValueOf(v.obj).Elem().IsNil() {
			newmap := reflect.MakeMap(reflect.MapOf(keyType, valueType))
			reflect.ValueOf(v.obj).Elem().Set(newmap)
		}
		ptrkey := reflect.New(keyType)
		if err := Unmarshal(name, ptrkey.Interface()); err != nil {
			panic(err)
		}

		ptrvalue := reflect.New(valueType)
		vv, err := NewVisitor(ptrvalue.Interface())
		if err != nil {
			panic(err)
		}
		res := value.Accept(vv)
		reflect.ValueOf(v.obj).Elem().SetMapIndex(ptrkey.Elem(), ptrvalue.Elem())
		return res

	case reflect.Struct:
		field := reflect.ValueOf(v.obj).Elem().FieldByName(name).Addr()
		vv, err := NewVisitor(field.Interface())
		if err != nil {
			panic(err)
		}
		res := value.Accept(vv)
		return res

	default:
	}
	return nil
}

// VisitObject_item ...
func (v *Visitor) VisitObject_item(ctx *parser.Object_itemContext) interface{} {
	key := ctx.SCALAR(0).GetText()
	switch {
	case ctx.SCALAR(1) != nil:
		return v.set(key, ctx.SCALAR(1))
	case ctx.Array() != nil:
		return v.set(key, ctx.Array())
	case ctx.Object() != nil:
		return v.set(key, ctx.Object())
	}
	return nil
}

// VisitObject_intern ...
func (v *Visitor) VisitObject_intern(ctx *parser.Object_internContext) interface{} {
	for _, item := range ctx.AllObject_item() {
		item.Accept(v)
	}
	return nil
}

// VisitArray ...
func (v *Visitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	if ctx.Array_intern() != nil {
		return ctx.Array_intern().Accept(v)
	}
	return nil
}

// VisitArray_item ...
func (v *Visitor) VisitArray_item(ctx *parser.Array_itemContext) interface{} {
	switch {
	case ctx.SCALAR() != nil:
		if err := v.feedScalar(ctx.GetText()); err != nil {
			panic(err)
		}
	case ctx.Array() != nil:
		ctx.Array().Accept(v)
	case ctx.Object() != nil:
		ctx.Object().Accept(v)
	}
	return nil
}

// VisitArray_intern ...
func (v *Visitor) VisitArray_intern(ctx *parser.Array_internContext) interface{} {
	if len(ctx.AllArray_item()) == 0 {
		return nil
	}
	objType := reflect.TypeOf(v.obj).Elem()
	if objType.Kind() != reflect.Slice {
		panic(fmt.Errorf("cannot unmarshal array into %T", reflect.ValueOf(v.obj).Elem().Interface()))
	}
	reflect.ValueOf(v.obj).Elem().Set(reflect.Zero(objType))
	for _, item := range ctx.AllArray_item() {
		newItem := reflect.New(reflect.TypeOf(v.obj).Elem().Elem())
		reflect.ValueOf(v.obj).Elem().Set(reflect.Append(reflect.ValueOf(v.obj).Elem(), newItem.Elem()))
		newItem = reflect.ValueOf(v.obj).Elem().Index(reflect.ValueOf(v.obj).Elem().Len() - 1)
		newv := v.withObject(newItem.Addr().Interface())
		item.Accept(newv)
	}
	return nil
}
