//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what resolveEnumTypeNames 이미 해석된 enum 건너뛰기 테스트
package nestjs

import "testing"

func TestResolveEnumTypeNames_SkipsAlreadyResolved(t *testing.T) {
	src := []byte(`enum TaskStatus { OPEN = 'open' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := []dtoField{
		{name: "status", enumTypeName: "TaskStatus", enum: []string{"preset"}},
	}
	resolveEnumTypeNames(fields, root, src, "/fake/dto.ts", nil, "")
	if len(fields[0].enum) != 1 || fields[0].enum[0] != "preset" {
		t.Fatalf("expected preset enum to be preserved, got %v", fields[0].enum)
	}
}
