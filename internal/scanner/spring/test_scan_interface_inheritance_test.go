//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_InterfaceInheritance — 인터페이스 상속 엔드포인트 추출 테스트
package spring

import "testing"

func TestScan_InterfaceInheritance(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/api/OwnersApi.java", interfaceApiSource)
	writeFile(t, dir, "src/main/java/com/example/controller/OwnerRestController.java", interfaceControllerSource)

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
	if ep0.Path != "/api/owners/{ownerId}" {
		t.Errorf("ep0 path: want /api/owners/{ownerId}, got %s", ep0.Path)
	}
	if ep0.Handler != "getOwner" {
		t.Errorf("ep0 handler: want getOwner, got %s", ep0.Handler)
	}
	if ep0.Request == nil || len(ep0.Request.PathParams) != 1 {
		t.Fatalf("ep0 expected 1 path param")
	}
	if ep0.Request.PathParams[0].Name != "ownerId" {
		t.Errorf("ep0 param name: want ownerId, got %s", ep0.Request.PathParams[0].Name)
	}

	ep1 := result.Endpoints[1]
	if ep1.Method != "POST" {
		t.Errorf("ep1 method: want POST, got %s", ep1.Method)
	}
	if ep1.Path != "/api/owners" {
		t.Errorf("ep1 path: want /api/owners, got %s", ep1.Path)
	}
	if ep1.Handler != "createOwner" {
		t.Errorf("ep1 handler: want createOwner, got %s", ep1.Handler)
	}
}
