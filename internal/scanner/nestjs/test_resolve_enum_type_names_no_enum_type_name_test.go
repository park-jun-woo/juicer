//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what resolveEnumTypeNames enumTypeName 없는 필드 테스트
package nestjs

import "testing"

func TestResolveEnumTypeNames_NoEnumTypeName(t *testing.T) {
	src := []byte(`class Dto { name: string; }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := []dtoField{
		{name: "name", tsType: "string"},
	}
	resolveEnumTypeNames(fields, root, src, "/fake/dto.ts", nil, "")
	if len(fields[0].enum) != 0 {
		t.Fatalf("expected no enum, got %v", fields[0].enum)
	}
}
