//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScan_FullE2E — Spring Boot 전체 E2E 스캔 테스트
package spring

import "testing"

func TestScan_FullE2E(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/demo/controller/ItemController.java", e2eControllerSource)
	writeFile(t, dir, "src/main/java/com/example/demo/dto/CreateItemRequest.java", e2eReqDtoSource)
	writeFile(t, dir, "src/main/java/com/example/demo/dto/ItemDto.java", e2eRespDtoSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 4 {
		t.Fatalf("expected 4 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" || ep0.Path != "/api/items" {
		t.Errorf("ep0: want GET /api/items, got %s %s", ep0.Method, ep0.Path)
	}
	if ep0.Request == nil || len(ep0.Request.Query) != 1 {
		t.Fatalf("ep0 expected 1 query param")
	}
	if ep0.Request.Query[0].Name != "page" {
		t.Errorf("ep0 query name: want page, got %s", ep0.Request.Query[0].Name)
	}
	if ep0.Request.Query[0].Default != "0" {
		t.Errorf("ep0 query default: want 0, got %s", ep0.Request.Query[0].Default)
	}

	ep1 := result.Endpoints[1]
	if ep1.Path != "/api/items/{id}" {
		t.Errorf("ep1 path: want /api/items/{id}, got %s", ep1.Path)
	}

	ep2 := result.Endpoints[2]
	if ep2.Method != "POST" {
		t.Errorf("ep2 method: want POST, got %s", ep2.Method)
	}
	if len(ep2.Responses) == 0 || ep2.Responses[0].Status != "201" {
		t.Errorf("ep2 status: want 201")
	}
	if len(ep2.Roles) != 1 || ep2.Roles[0] != "ADMIN" {
		t.Errorf("ep2 roles: want [ADMIN], got %v", ep2.Roles)
	}
	if ep2.Request == nil || ep2.Request.Body == nil {
		t.Fatalf("ep2 expected body")
	}
	if len(ep2.Request.Body.Fields) != 3 {
		t.Fatalf("ep2 body fields: want 3, got %d", len(ep2.Request.Body.Fields))
	}
	nameField := ep2.Request.Body.Fields[0]
	if nameField.Name != "name" {
		t.Errorf("field[0] name: want name, got %s", nameField.Name)
	}
	if nameField.Validate != "required" {
		t.Errorf("field[0] validate: want required, got %s", nameField.Validate)
	}
	if nameField.MinLength == nil || *nameField.MinLength != 1 {
		t.Errorf("field[0] minLength: want 1")
	}
	if nameField.MaxLength == nil || *nameField.MaxLength != 100 {
		t.Errorf("field[0] maxLength: want 100")
	}
	priceField := ep2.Request.Body.Fields[1]
	if priceField.JSON != "item_price" {
		t.Errorf("field[1] json: want item_price, got %s", priceField.JSON)
	}
	if priceField.Minimum == nil || *priceField.Minimum != 0 {
		t.Errorf("field[1] minimum: want 0")
	}

	ep3 := result.Endpoints[3]
	if ep3.Method != "DELETE" {
		t.Errorf("ep3 method: want DELETE, got %s", ep3.Method)
	}
	if len(ep3.Responses) == 0 || ep3.Responses[0].Status != "204" {
		t.Errorf("ep3 status: want 204")
	}
}
