//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyJSONExtractor_Round5 테스트
package actix

import "testing"

func TestApplyJSONExtractor_Round5(t *testing.T) {
	ep, sIdx, cache := extractorTestSetup()
	applyJSONExtractor(ep, extractorInfo{kind: "json", typeName: "CreateReq"}, sIdx, cache)
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "CreateReq" {
		t.Fatalf("body: %+v", ep.Request.Body)
	}
}
