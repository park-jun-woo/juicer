//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOExtends_NoHeritage 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOExtends_NoHeritage(t *testing.T) {
	src := []byte(`class Dto { name: string; }`)
	root, _ := parseTypeScript(src)
	cls := findAllByType(root, "class_declaration")[0]
	if got := resolveDTOExtends(cls, src, "x.ts", map[string]string{}, "", map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
