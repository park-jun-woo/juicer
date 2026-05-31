//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what appendAstField struct 필드 → scanner.Field(미exported 스킵) 테스트
package fiber

import "testing"

func TestAppendAstField(t *testing.T) {
	fields := fiberFields(t, "package p\ntype T struct {\n Name string `json:\"name\"`\n priv int\n}")
	out := appendAstField(nil, fields[0])
	if len(out) != 1 || out[0].Name != "Name" || out[0].JSON != "name" {
		t.Errorf("exported: %+v", out)
	}
	out = appendAstField(nil, fields[1])
	if len(out) != 0 {
		t.Errorf("unexported skipped: %+v", out)
	}
}
