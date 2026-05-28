//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestScan_SameFileInnerClass — 같은 파일 내부 클래스 필드 해석 테스트
package spring

import "testing"

func TestScan_SameFileInnerClass(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/controller/AuthController.java", sameFileInnerClassSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected body")
	}
	if len(ep.Request.Body.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(ep.Request.Body.Fields))
	}
	names := map[string]bool{}
	for _, f := range ep.Request.Body.Fields {
		names[f.Name] = true
	}
	if !names["email"] {
		t.Errorf("expected field 'email'")
	}
	if !names["password"] {
		t.Errorf("expected field 'password'")
	}
}
