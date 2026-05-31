//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what parseField 필드 라인 토큰화(array/nullable/attrs) 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestParseField(t *testing.T) {
	f := parseField("id Int @id @default(autoincrement())")
	if f.name != "id" || f.baseType != "Int" || f.nullable || f.array {
		t.Errorf("scalar: %+v", f)
	}
	if !reflect.DeepEqual(f.attrs, []string{"@id", "@default(autoincrement())"}) {
		t.Errorf("attrs: %v", f.attrs)
	}
	arr := parseField("tags String[]")
	if !arr.array || arr.baseType != "String" {
		t.Errorf("array: %+v", arr)
	}
	opt := parseField("bio String?")
	if !opt.nullable || opt.baseType != "String" {
		t.Errorf("nullable: %+v", opt)
	}
}
