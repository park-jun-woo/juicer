//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_Inheritance — DTO 상속 필드 해석 테스트
package spring

import "testing"

func TestScan_Inheritance(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/controller/TestController.java", inheritanceControllerSource)
	writeFile(t, dir, "src/main/java/com/example/dto/BaseDto.java", inheritanceParentDtoSource)
	writeFile(t, dir, "src/main/java/com/example/dto/ChildDto.java", inheritanceChildDtoSource)

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
		t.Fatalf("expected 3 fields (2 parent + 1 own), got %d", len(ep.Request.Body.Fields))
	}
}
