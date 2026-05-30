//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyQueryExtractor_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyQueryExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	sIdx := structIndexFor(t, `struct Filter { page: i64, q: String }`)
	cache := map[string][]scanner.Field{}
	applyQueryExtractor(ep, extractorInfo{kind: "query", typeName: "Filter"}, sIdx, cache)
	if ep.Request == nil || len(ep.Request.Query) == 0 {
		t.Fatalf("expected query params: %+v", ep.Request)
	}
}
