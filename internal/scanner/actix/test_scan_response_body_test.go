//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScan_ResponseBody — HttpResponse::*().json(struct) 본문 스키마 추출 테스트
package actix

import "testing"

func TestScan_ResponseBody(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", responseBodySource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	wantStatus := map[string]string{"create_url": "201", "get_url": "200"}
	for _, ep := range result.Endpoints {
		if len(ep.Responses) != 1 {
			t.Fatalf("%s expected 1 response, got %d", ep.Handler, len(ep.Responses))
		}
		r := ep.Responses[0]
		if r.Status != wantStatus[ep.Handler] {
			t.Errorf("%s status: want %s, got %s", ep.Handler, wantStatus[ep.Handler], r.Status)
		}
		if r.Kind != "json" {
			t.Errorf("%s kind: want json, got %s", ep.Handler, r.Kind)
		}
		if r.TypeName != "CreatedURL" {
			t.Errorf("%s typeName: want CreatedURL, got %s", ep.Handler, r.TypeName)
		}
		if len(r.Fields) != 3 {
			t.Fatalf("%s expected 3 fields, got %d", ep.Handler, len(r.Fields))
		}
		if r.Confidence != "full" {
			t.Errorf("%s confidence: want full, got %s", ep.Handler, r.Confidence)
		}
	}
}
