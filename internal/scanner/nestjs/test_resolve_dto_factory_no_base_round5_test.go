//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOFactory_NoBase_Round5 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOFactory_NoBase_Round5(t *testing.T) {

	b := []byte("class C extends PartialType('x') {}")
	root, _ := parseTypeScript(b)
	calls := findAllByType(root, "call_expression")
	got := resolveDTOFactory(calls[0], b, "/proj/c.dto.ts", map[string]string{}, "/proj", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil when base unresolved, got %+v", got)
	}
}
