//ff:func feature=scan type=test control=sequence topic=actix
//ff:what joinTypeArgTokens — 꺾쇠 제외 토큰 이어붙이기를 검증
package actix

import (
	"strings"
	"testing"
)

func TestJoinTypeArgTokens(t *testing.T) {
	src := []byte(`fn f(x: HashMap<String, i32>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	ta := firstTypeArguments(root)
	if ta == nil {
		t.Fatal("no type_arguments")
	}
	got := joinTypeArgTokens(ta, src)
	// Angle brackets stripped, inner type names preserved.
	if strings.ContainsAny(got, "<>") {
		t.Errorf("expected no angle brackets, got %q", got)
	}
	if !strings.Contains(got, "String") || !strings.Contains(got, "i32") {
		t.Errorf("expected both type names, got %q", got)
	}
}
