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

func TestRegression(t *testing.T) {
	data := `{ProtoInfo{github.com/shilkin/grpc-gateway/third_party/googleapis/google/api/annotations.proto{FPath=/home/emacs/Sources/go/src/github.com/shilkin/grpc-gateway/third_party/googleapis/google/api/annotations.proto},atlas/error_codes.proto{FPath=/home/emacs/Sources/schema/atlas/error_codes.proto,IPath=atlas},bazaar/common/avatar/avatar.proto{FPath=/home/emacs/Sources/schema/bazaar/common/avatar/avatar.proto},google/api/http.proto{FPath=/home/emacs/Sources/grpc-gateway/third_party/googleapis/google/api/http.proto},common/service_status.proto{FPath=/home/emacs/Sources/schema/common/service_status.proto},beef/meta.proto{FPath=/home/emacs/Sources/schema/beef/meta.proto},beef/mail.proto{FPath=/home/emacs/Sources/schema/beef/mail.proto},marker/object.proto{FPath=/home/emacs/Sources/schema/marker/object.proto},common/error.proto{FPath=/home/emacs/Sources/schema/common/error.proto},common/detail.proto{FPath=/home/emacs/Sources/schema/common/detail.proto},marker/permissions.proto{FPath=/home/emacs/Sources/schema/marker/permissions.proto},github.com/gogo/protobuf/gogoproto/gogo.proto{FPath=/home/emacs/Sources/go/src/github.com/gogo/protobuf/gogoproto/gogo.proto},atlas/atlas.proto{FPath=/home/emacs/Sources/schema/atlas/atlas.proto,IPath=atlas},common/module.proto{FPath=/home/emacs/Sources/schema/common/module.proto},marker/tag.proto{FPath=/home/emacs/Sources/schema/marker/tag.proto},common/doctype.proto{FPath=/home/emacs/Sources/schema/common/doctype.proto},google/protobuf/descriptor.proto{FPath=/home/emacs/Sources/protobuf/src/google/protobuf/descriptor.proto},bazaar/common/paging/paging.proto{FPath=/home/emacs/Sources/schema/bazaar/common/paging/paging.proto},bazaar/entities/user/user.proto{FPath=/home/emacs/Sources/schema/bazaar/entities/user/user.proto},bazaar/common/email/email.proto{FPath=/home/emacs/Sources/schema/bazaar/common/email/email.proto},common/timestamp.proto{FPath=/home/emacs/Sources/schema/common/timestamp.proto},github.com/shilkin/grpc-gateway/options/method.proto{FPath=/home/emacs/Sources/go/src/github.com/shilkin/grpc-gateway/options/method.proto},bazaar/entities/region/region.proto{FPath=/home/emacs/Sources/schema/bazaar/entities/region/region.proto}}}`
	type (
		// ProtoFileData содержит следующую информацию касающуюся proto-файла
		// * абсолютный путь к proto-файлу
		// * под каким путём сгенерированный Go-код доступен для импорта
		ProtoFileData struct {
			FPath string // FPath – абсолютный путь к proto-файлу
			IPath string // IPath - путь для импорта
		}

		// PluginParams параметры плагина
		PluginParams struct {
			ProtoInfo map[string]ProtoFileData
		}
	)

	v := PluginParams{}
	if err := Unmarshal(data, &v); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ProtoFileData{
		FPath: "/home/emacs/Sources/protobuf/src/google/protobuf/descriptor.proto",
		IPath: "",
	}, v.ProtoInfo["google/protobuf/descriptor.proto"])
	require.Equal(t, ProtoFileData{
		FPath: "/home/emacs/Sources/schema/common/timestamp.proto",
		IPath: "",
	}, v.ProtoInfo["common/timestamp.proto"])
	require.Equal(t, ProtoFileData{
		FPath: "/home/emacs/Sources/schema/atlas/atlas.proto",
		IPath: "atlas",
	}, v.ProtoInfo["atlas/atlas.proto"])
}
