//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScan_FormExtractor — Form extractor 스캔 테스트
package actix

import "testing"

func TestScan_FormExtractor(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/auth.rs", formExtractorSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Method != "POST" {
		t.Errorf("method: want POST, got %s", ep.Method)
	}
	if ep.Path != "/login" {
		t.Errorf("path: want /login, got %s", ep.Path)
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.FormFields) != 2 {
		t.Fatalf("expected 2 form fields, got %d", len(ep.Request.FormFields))
	}

	names := map[string]bool{}
	for _, f := range ep.Request.FormFields {
		names[f.Name] = true
	}
	if !names["username"] {
		t.Errorf("expected form field username")
	}
	if !names["password"] {
		t.Errorf("expected form field password")
	}
}
