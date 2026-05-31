//ff:func feature=scan type=test topic=echo control=sequence
//ff:what appendAstField struct 필드 → scanner.Field(미exported/무이름 스킵) 테스트
package echo

import "testing"

func TestAppendAstField(t *testing.T) {
	st := parseEchoStruct(t, `package p
type T struct {
	Name string `+"`json:\"name\"`"+`
	priv int
}`)
	var fields = appendAstField(nil, st.Fields.List[0])
	if len(fields) != 1 || fields[0].Name != "Name" || fields[0].JSON != "name" {
		t.Errorf("exported: %+v", fields)
	}
	// unexported field skipped
	fields = appendAstField(nil, st.Fields.List[1])
	if len(fields) != 0 {
		t.Errorf("unexported should be skipped: %+v", fields)
	}
}
