//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_ConstantDefaultValue — 상수 참조 defaultValue 해석 테스트
package spring

import "testing"

func TestScan_ConstantDefaultValue(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/controller/PostController.java", constDefaultControllerSource)
	writeFile(t, dir, "src/main/java/com/example/constants/AppConstants.java", constAppConstantsSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil || len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ep.Request.Query))
	}
	page := ep.Request.Query[0]
	if page.Name != "page" {
		t.Errorf("param[0] name: want page, got %s", page.Name)
	}
	if page.Default != "0" {
		t.Errorf("param[0] default: want 0, got %s", page.Default)
	}
	size := ep.Request.Query[1]
	if size.Name != "size" {
		t.Errorf("param[1] name: want size, got %s", size.Name)
	}
	if size.Default != "30" {
		t.Errorf("param[1] default: want 30, got %s", size.Default)
	}
}
