//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_DirectAnnotationNoRegression — 직접 어노테이션 컨트롤러가 인터페이스 로직 영향 없이 동작하는지 테스트
package spring

import "testing"

func TestScan_DirectAnnotationNoRegression(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/controller/PetController.java", directAnnotationControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" {
		t.Errorf("ep0 method: want GET, got %s", ep0.Method)
	}
	if ep0.Path != "/api/pets/{petId}" {
		t.Errorf("ep0 path: want /api/pets/{petId}, got %s", ep0.Path)
	}

	ep1 := result.Endpoints[1]
	if ep1.Method != "POST" {
		t.Errorf("ep1 method: want POST, got %s", ep1.Method)
	}
	if ep1.Path != "/api/pets" {
		t.Errorf("ep1 path: want /api/pets, got %s", ep1.Path)
	}
}
