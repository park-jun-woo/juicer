//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what astFieldJSONName json 태그명 추출(무태그/"-"/콤마옵션) 테스트
package fiber

import "testing"

func TestAstFieldJSONName(t *testing.T) {
	fields := fiberFields(t, "package p\ntype T struct {\n A string `json:\"name,omitempty\"`\n B int\n C bool `json:\"-\"`\n}")
	if got := astFieldJSONName(fields[0]); got != "name" {
		t.Errorf("tag: %q", got)
	}
	if got := astFieldJSONName(fields[1]); got != "" {
		t.Errorf("no tag: %q", got)
	}
	if got := astFieldJSONName(fields[2]); got != "" {
		t.Errorf("dash tag: %q", got)
	}
}
