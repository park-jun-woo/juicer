//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScan_ResponseStatus — HttpResponse 상태 코드 추출 테스트
package actix

import "testing"

func TestScan_ResponseStatus(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", macroRoutesSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	// ep0: GET /users/{id} -> HttpResponse::Ok().json() => 200
	ep0 := result.Endpoints[0]
	if len(ep0.Responses) != 1 {
		t.Fatalf("ep0 expected 1 response, got %d", len(ep0.Responses))
	}
	if ep0.Responses[0].Status != "200" {
		t.Errorf("ep0 response status: want 200, got %s", ep0.Responses[0].Status)
	}
	if ep0.Responses[0].Kind != "json" {
		t.Errorf("ep0 response kind: want json, got %s", ep0.Responses[0].Kind)
	}

	// ep1: POST /users -> HttpResponse::Created().json() => 201
	ep1 := result.Endpoints[1]
	if len(ep1.Responses) != 1 {
		t.Fatalf("ep1 expected 1 response, got %d", len(ep1.Responses))
	}
	if ep1.Responses[0].Status != "201" {
		t.Errorf("ep1 response status: want 201, got %s", ep1.Responses[0].Status)
	}

	// ep3: DELETE /users/{id} -> HttpResponse::NoContent().finish() => 204
	ep3 := result.Endpoints[3]
	if len(ep3.Responses) != 1 {
		t.Fatalf("ep3 expected 1 response, got %d", len(ep3.Responses))
	}
	if ep3.Responses[0].Status != "204" {
		t.Errorf("ep3 response status: want 204, got %s", ep3.Responses[0].Status)
	}
}
