//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestScan_MacroRoutes — proc-macro 라우트 스캔 테스트
package actix

import "testing"

func TestScan_MacroRoutes(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", macroRoutesSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 4 {
		t.Fatalf("expected 4 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" {
		t.Errorf("ep0 method: want GET, got %s", ep0.Method)
	}
	if ep0.Path != "/users/{id}" {
		t.Errorf("ep0 path: want /users/{id}, got %s", ep0.Path)
	}
	if ep0.Handler != "get_user" {
		t.Errorf("ep0 handler: want get_user, got %s", ep0.Handler)
	}
	if ep0.Request == nil || len(ep0.Request.PathParams) != 1 {
		t.Fatalf("ep0 expected 1 path param")
	}
	if ep0.Request.PathParams[0].Name != "id" {
		t.Errorf("ep0 param name: want id, got %s", ep0.Request.PathParams[0].Name)
	}
	if ep0.Request.PathParams[0].Type != "integer" {
		t.Errorf("ep0 param type: want integer, got %s", ep0.Request.PathParams[0].Type)
	}

	ep1 := result.Endpoints[1]
	if ep1.Method != "POST" {
		t.Errorf("ep1 method: want POST, got %s", ep1.Method)
	}
	if ep1.Path != "/users" {
		t.Errorf("ep1 path: want /users, got %s", ep1.Path)
	}
	if ep1.Request == nil || ep1.Request.Body == nil {
		t.Fatalf("ep1 expected body")
	}
	if ep1.Request.Body.TypeName != "CreateUserRequest" {
		t.Errorf("ep1 body type: want CreateUserRequest, got %s", ep1.Request.Body.TypeName)
	}
	if len(ep1.Request.Body.Fields) != 3 {
		t.Fatalf("ep1 body fields: want 3, got %d", len(ep1.Request.Body.Fields))
	}

	ep2 := result.Endpoints[2]
	if ep2.Method != "PUT" {
		t.Errorf("ep2 method: want PUT, got %s", ep2.Method)
	}

	ep3 := result.Endpoints[3]
	if ep3.Method != "DELETE" {
		t.Errorf("ep3 method: want DELETE, got %s", ep3.Method)
	}
}
