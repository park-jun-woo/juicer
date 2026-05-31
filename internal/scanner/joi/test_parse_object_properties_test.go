//ff:func feature=scan type=test topic=joi control=sequence
//ff:what ParseObjectProperties object pair 순회 → Field 슬라이스 테스트
package joi

import "testing"

func TestParseObjectProperties(t *testing.T) {
	root, src := parseJoiTS(t, `const o = { name: Joi.string(), age: Joi.number() };`)
	obj := firstOfType(root, "object")
	fields := ParseObjectProperties(obj, src)
	if len(fields) != 2 {
		t.Fatalf("want 2 fields, got %+v", fields)
	}
	if fields[0].Name != "name" || fields[1].Name != "age" {
		t.Errorf("names: %+v", fields)
	}
}
