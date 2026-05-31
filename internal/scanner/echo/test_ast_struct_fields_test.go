//ff:func feature=scan type=test topic=echo control=sequence
//ff:what astStructFields struct → exported 필드 슬라이스 변환 테스트
package echo

import "testing"

func TestAstStructFields(t *testing.T) {
	st := parseEchoStruct(t, `package p
type T struct {
	A string
	B int
	c bool
}`)
	fields := astStructFields(st)
	if len(fields) != 2 || fields[0].Name != "A" || fields[1].Name != "B" {
		t.Errorf("got %+v", fields)
	}
}
