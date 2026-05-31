//ff:func feature=scan type=test topic=joi control=sequence
//ff:what ParseSchema Joi.object().keys({...}) → Field 슬라이스 및 nil/미지원 격하 테스트
package joi

import "testing"

func TestParseSchema(t *testing.T) {
	// nil node
	if ParseSchema(nil, nil) != nil {
		t.Error("nil node should be nil")
	}
	// valid keys object
	root, src := parseJoiTS(t, `Joi.object().keys({ name: Joi.string().required() })`)
	call := firstOfType(root, "call_expression")
	fields := ParseSchema(call, src)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
	// no keys object -> nil
	root2, src2 := parseJoiTS(t, `Joi.string()`)
	call2 := firstOfType(root2, "call_expression")
	if ParseSchema(call2, src2) != nil {
		t.Error("no keys object should be nil")
	}
}
