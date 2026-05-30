//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveParentDTOFields_Cached 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveParentDTOFields_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{
		"ParentDto": {{Name: "x", Type: "string"}},
	}
	got := resolveParentDTOFields("ParentDto", "x.ts", map[string]string{}, "", cache)
	if len(got) != 1 || got[0].Name != "x" {
		t.Fatalf("got %+v", got)
	}
}
