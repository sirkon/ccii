package ccii

import (
	"bytes"
	"fmt"
	"reflect"
)

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

// Marshal обратное преобразование
func Marshal(obj interface{}) (string, error) {
	m := &marshaler{}
	if err := m.marshal(obj, ""); err != nil {
		return "", err
	}
	return m.String(), nil
}

type marshaler struct {
	bytes.Buffer
}

func (m *marshaler) marshal(obj interface{}, scalarPrefix string) error {
	v := reflect.ValueOf(obj)
	switch v.Type().Kind() {
	case reflect.Ptr:
		if err := m.marshal(v.Elem().Interface(), scalarPrefix); err != nil {
			return err
		}
	case reflect.Struct:
		m.WriteByte('{')
		if err := m.writeStruct(v); err != nil {
			return err
		}
		m.WriteByte('}')
	case reflect.Map:

		m.WriteByte('{')
		notFirst := false
		for _, key := range v.MapKeys() {
			if isZero(v.MapIndex(key)) {
				continue
			}
			if notFirst {
				m.WriteByte(',')
			}
			notFirst = true
			if err := m.marshal(key.Interface(), ""); err != nil {
				return err
			}
			if err := m.marshal(v.MapIndex(key).Interface(), "="); err != nil {
				return err
			}
		}
		m.WriteByte('}')
	case reflect.Slice:
		m.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if err := m.marshal(v.Index(i).Interface(), ""); err != nil {
				return err
			}
			if i < v.Len()-1 {
				m.WriteByte(',')
			}
		}
		m.WriteByte(']')
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		m.WriteString(scalarPrefix)
		m.WriteString(fmt.Sprintf("%d", obj))
	case reflect.Float32, reflect.Float64:
		m.WriteString(scalarPrefix)
		m.WriteString(fmt.Sprintf("%f", obj))
	case reflect.String:
		m.WriteString(scalarPrefix)
		m.WriteString(Encode(obj.(string)))
	case reflect.Bool:
		m.WriteString(scalarPrefix)
		if obj.(bool) {
			m.WriteByte('t')
		} else {
			m.WriteByte('f')
		}
	default:
		return fmt.Errorf("type %T is not supported for marshalling", obj)
	}
	return nil
}

func (m *marshaler) writeStruct(v reflect.Value) error {
	notFirst := false
	for i := 0; i < v.NumField(); i++ {
		ftype := v.Type().Field(i)
		fval := v.Field(i)
		if isZero(fval) {
			continue
		}
		if notFirst {
			m.WriteByte(',')
		}
		notFirst = true
		if ftype.Anonymous {
			if err := m.writeStruct(fval); err != nil {
				return err
			}
		} else {
			if err := m.marshal(ftype.Name, ""); err != nil {
				return err
			}
			if err := m.marshal(fval.Interface(), "="); err != nil {
				return err
			}

		}
	}
	return nil
}
