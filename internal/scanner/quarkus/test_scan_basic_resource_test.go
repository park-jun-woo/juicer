//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestScan_BasicResource -- 기본 JAX-RS 리소스 스캔 테스트
package quarkus

import "testing"

func TestScan_BasicResource(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/UserResource.java", basicUserResourceSource)
	writeFile(t, dir, "src/main/java/com/example/dto/CreateUserRequest.java", basicCreateUserRequestSource)
	writeFile(t, dir, "src/main/java/com/example/dto/UserResponse.java", basicUserResponseSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" {
		t.Errorf("ep0 method: want GET, got %s", ep0.Method)
	}
	if ep0.Path != "/api/users" {
		t.Errorf("ep0 path: want /api/users, got %s", ep0.Path)
	}
	if ep0.Handler != "listUsers" {
		t.Errorf("ep0 handler: want listUsers, got %s", ep0.Handler)
	}
	if len(ep0.Responses) == 0 {
		t.Fatal("ep0 expected response")
	}
	if ep0.Responses[0].Body != "array" {
		t.Errorf("ep0 response body: want array, got %s", ep0.Responses[0].Body)
	}

	ep1 := result.Endpoints[1]
	if ep1.Path != "/api/users/{id}" {
		t.Errorf("ep1 path: want /api/users/{id}, got %s", ep1.Path)
	}
	if ep1.Request == nil || len(ep1.Request.PathParams) != 1 {
		t.Fatalf("ep1 expected 1 path param")
	}
	if ep1.Request.PathParams[0].Name != "id" {
		t.Errorf("ep1 param name: want id, got %s", ep1.Request.PathParams[0].Name)
	}

	ep2 := result.Endpoints[2]
	if ep2.Method != "POST" {
		t.Errorf("ep2 method: want POST, got %s", ep2.Method)
	}
	if ep2.Request == nil || ep2.Request.Body == nil {
		t.Fatalf("ep2 expected body")
	}
	if ep2.Request.Body.TypeName != "CreateUserRequest" {
		t.Errorf("ep2 body type: want CreateUserRequest, got %s", ep2.Request.Body.TypeName)
	}
	if len(ep2.Request.Body.Fields) != 3 {
		t.Fatalf("ep2 body fields: want 3, got %d", len(ep2.Request.Body.Fields))
	}
}
