//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOExtends_PlainExtends 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOExtends_PlainExtends(t *testing.T) {
	src := []byte(`class Child extends Parent {}`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]

	got := resolveDTOExtends(cls, src, "x.ts", map[string]string{}, "", map[string][]scanner.Field{})
	if len(got) != 0 {
		t.Fatalf("expected empty, got %+v", got)
	}
}
