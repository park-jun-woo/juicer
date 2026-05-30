//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyPathExtractor_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyPathExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/u/{id}")
	sIdx := structIndex{}
	cache := map[string][]scanner.Field{}
	applyPathExtractor(ep, extractorInfo{kind: "path", typeName: "i64"}, sIdx, cache)
	if ep.Request == nil {
		t.Fatal("expected request")
	}
}
