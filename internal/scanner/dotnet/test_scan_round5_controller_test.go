//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestScan_Round5Controller 테스트
package dotnet

import "testing"

func TestScan_Round5Controller(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Models/Models.cs", round5RecordSource)
	writeFile(t, dir, "Controllers/WidgetsController.cs", round5ControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	create := result.Endpoints[0]
	if create.Path != "/api/v1/widgets" {
		t.Errorf("create path: want /api/v1/widgets, got %s", create.Path)
	}
	if len(create.Roles) == 0 {
		t.Errorf("expected authorize roles propagated, got %+v", create.Roles)
	}
	if create.Request == nil || create.Request.Body == nil {
		t.Fatalf("create body missing: %+v", create.Request)
	}

	if len(create.Request.Body.Fields) == 0 {
		t.Errorf("record fields not resolved: %+v", create.Request.Body)
	}

	list := result.Endpoints[1]
	if list.Path != "/api/v1/widgets/{id}" {
		t.Errorf("list path: want /api/v1/widgets/{id}, got %s", list.Path)
	}
	if len(list.Responses) == 0 || list.Responses[0].Body != "array" {
		t.Errorf("list should return array: %+v", list.Responses)
	}

	remove := result.Endpoints[2]
	if len(remove.Responses) == 0 {
		t.Fatalf("remove should have StatusCode response: %+v", remove.Responses)
	}
	found503 := false
	for _, r := range remove.Responses {
		if r.Status != "503" {
			continue
		}
		found503 = true
		if r.TypeName != "ItemResponse" {
			t.Errorf("503 type: want ItemResponse, got %s", r.TypeName)
		}
	}
	if !found503 {
		t.Errorf("expected 503 StatusCode response: %+v", remove.Responses)
	}
}
