//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyFormExtractor_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFormExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	sIdx := structIndexFor(t, `struct FormData { title: String }`)
	cache := map[string][]scanner.Field{}
	applyFormExtractor(ep, extractorInfo{kind: "form", typeName: "FormData"}, sIdx, cache)
	if ep.Request == nil || len(ep.Request.FormFields) == 0 {
		t.Fatalf("expected form fields: %+v", ep.Request)
	}
}
