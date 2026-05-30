//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestApplyExtractors_AllKinds_Round5 테스트
package actix

import "testing"

func TestApplyExtractors_AllKinds_Round5(t *testing.T) {
	for _, kind := range []string{"json", "query", "form", "path"} {
		ep, sIdx, cache := extractorTestSetup()
		ext := extractorInfo{kind: kind, typeName: "i64", rawType: "web::" + kind}
		applyExtractor(ep, ext, sIdx, cache)
	}

	ep, sIdx, cache := extractorTestSetup()
	applyExtractors(ep, []extractorInfo{
		{kind: "json", typeName: "CreateReq"},
		{kind: "query", typeName: "Filter"},
	}, sIdx, cache)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected json body: %+v", ep.Request)
	}
}
