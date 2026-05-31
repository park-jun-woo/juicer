//ff:func feature=scan type=test topic=joi control=sequence
//ff:what ParsePair pair → Field(name/json/type) 변환 및 nil 가드 테스트
package joi

import "testing"

func TestParsePair(t *testing.T) {
	root, src := parseJoiTS(t, `const o = { name: Joi.string().required() };`)
	pair := firstOfType(root, "pair")
	if pair == nil {
		t.Fatal("no pair")
	}
	f := ParsePair(pair, src)
	if f == nil || f.Name != "name" || f.JSON != "name" || f.Type != "string" {
		t.Fatalf("got %+v", f)
	}
}
