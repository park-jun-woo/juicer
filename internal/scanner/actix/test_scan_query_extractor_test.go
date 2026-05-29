//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScan_QueryExtractor — Query extractor 스캔 테스트
package actix

import "testing"

func TestScan_QueryExtractor(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/handlers.rs", queryExtractorSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Method != "GET" {
		t.Errorf("method: want GET, got %s", ep.Method)
	}
	if ep.Path != "/items" {
		t.Errorf("path: want /items, got %s", ep.Path)
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.Query) != 3 {
		t.Fatalf("expected 3 query params, got %d", len(ep.Request.Query))
	}

	names := map[string]bool{}
	for _, q := range ep.Request.Query {
		names[q.Name] = true
	}
	for _, want := range []string{"page", "limit", "search"} {
		if !names[want] {
			t.Errorf("expected query param %s", want)
		}
	}
}
