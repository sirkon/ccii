package ccii

import (
	"bytes"
	"fmt"
	"reflect"
)

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
		for i, key := range v.MapKeys() {
			if err := m.marshal(key.Interface(), ""); err != nil {
				return err
			}
			if err := m.marshal(v.MapIndex(key).Interface(), "="); err != nil {
				return err
			}
			if i < len(v.MapKeys())-1 {
				m.WriteByte(',')
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
	default:
		return fmt.Errorf("type %T is not supported for marshalling", obj)
	}
	return nil
}

func (m *marshaler) writeStruct(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		ftype := v.Type().Field(i)
		fval := v.Field(i)
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
		if i < v.NumField()-1 {
			m.WriteByte(',')
		}
	}
	return nil
}
