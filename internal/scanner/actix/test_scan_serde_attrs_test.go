//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScan_SerdeAttrs — serde 어트리뷰트 적용 테스트
package actix

import "testing"

func TestScan_SerdeAttrs(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/handler.rs", serdeAttrsSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body")
	}

	fields := ep.Request.Body.Fields
	// serde(skip) should exclude internal_id, leaving 2 fields
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields (skip removed 1), got %d", len(fields))
	}

	// Check rename
	nameField := fields[0]
	if nameField.JSON != "userName" {
		t.Errorf("expected JSON name 'userName', got %s", nameField.JSON)
	}

	// Check default makes nullable
	bioField := fields[1]
	if !bioField.Nullable {
		t.Errorf("expected bio field to be nullable (Option + serde default)")
	}
}
