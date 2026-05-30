//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasSetGlobalPrefix 테스트
package nestjs

import "testing"

func TestHasSetGlobalPrefix(t *testing.T) {
	src := []byte(`app.setGlobalPrefix('api');`)
	root, _ := parseTypeScript(src)
	calls := findAllByType(root, "call_expression")
	if !hasSetGlobalPrefix(calls[0], src) {
		t.Fatal("expected true")
	}
	src2 := []byte(`app.listen(3000);`)
	root2, _ := parseTypeScript(src2)
	calls2 := findAllByType(root2, "call_expression")
	if hasSetGlobalPrefix(calls2[0], src2) {
		t.Fatal("expected false")
	}
}
