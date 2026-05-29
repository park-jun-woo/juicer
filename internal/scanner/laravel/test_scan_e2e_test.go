//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what E2E 스캔 테스트: 라우트 + apiResource + prefix 그룹 + middleware 그룹 + FormRequest + Resource
package laravel

import "testing"

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()
	writeE2EProject(t, dir)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Endpoints) < 10 {
		logEndpoints(t, result.Endpoints)
		t.Fatalf("expected at least 10 endpoints, got %d", len(result.Endpoints))
	}

	assertE2EEndpoints(t, result.Endpoints)
	assertE2EFormRequest(t, result.Endpoints)
	assertE2EResourceResponse(t, result.Endpoints)
	assertE2EMiddleware(t, result.Endpoints)
}
