//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_InterfaceWithControllerPrefix — 컨트롤러 prefix가 인터페이스 prefix를 우선하는지 테스트
package spring

import "testing"

func TestScan_InterfaceWithControllerPrefix(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/api/OwnersApi.java", interfaceApiSource)
	writeFile(t, dir, "src/main/java/com/example/controller/OwnerRestControllerV2.java", interfaceControllerWithPrefixSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Path != "/v2/owners/{ownerId}" {
		t.Errorf("ep0 path: want /v2/owners/{ownerId}, got %s", ep0.Path)
	}

	ep1 := result.Endpoints[1]
	if ep1.Path != "/v2/owners" {
		t.Errorf("ep1 path: want /v2/owners, got %s", ep1.Path)
	}
}
