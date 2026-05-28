//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestScan_CrossPackageParent — 타 패키지 부모 클래스 상속 필드 해석 테스트
package spring

import "testing"

func TestScan_CrossPackageParent(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/controller/AlbumController.java", crossPkgControllerSource)
	writeFile(t, dir, "src/main/java/com/example/dto/AlbumRequest.java", crossPkgChildDtoSource)
	writeFile(t, dir, "src/main/java/com/example/audit/UserDateAuditPayload.java", crossPkgParentDtoSource)

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
	if len(ep.Request.Body.Fields) != 3 {
		t.Fatalf("expected 3 fields (1 own + 2 parent), got %d", len(ep.Request.Body.Fields))
	}
	names := map[string]bool{}
	for _, f := range ep.Request.Body.Fields {
		names[f.Name] = true
	}
	if !names["title"] {
		t.Errorf("expected field 'title'")
	}
	if !names["userId"] {
		t.Errorf("expected parent field 'userId'")
	}
	if !names["updatedAt"] {
		t.Errorf("expected parent field 'updatedAt'")
	}
}
