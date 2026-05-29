//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_BodyResponse -- bare IActionResult 본문 return 식에서 타입·상태 코드 추출
package dotnet

import "testing"

func TestScan_BodyResponse(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/ItemsController.cs", bodyResponseControllerSource)
	writeFile(t, dir, "Models/Models.cs", bodyResponseModelsSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	get := result.Endpoints[0]
	if len(get.Responses) != 1 {
		t.Fatalf("GET expected 1 response, got %d", len(get.Responses))
	}
	if get.Responses[0].Status != "200" {
		t.Errorf("GET status: want 200, got %s", get.Responses[0].Status)
	}
	if get.Responses[0].TypeName != "ItemResponse" {
		t.Errorf("GET type: want ItemResponse, got %s", get.Responses[0].TypeName)
	}

	post := result.Endpoints[1]
	if len(post.Responses) != 1 {
		t.Fatalf("POST expected 1 response, got %d", len(post.Responses))
	}
	if post.Responses[0].Status != "200" {
		t.Errorf("POST status: want 200 (not 201 default), got %s", post.Responses[0].Status)
	}
	if post.Responses[0].TypeName != "ResponseViewModel<ItemResponse>" {
		t.Errorf("POST type: want ResponseViewModel<ItemResponse>, got %s", post.Responses[0].TypeName)
	}
	if len(post.Responses[0].Fields) == 0 {
		t.Errorf("POST envelope fields not resolved: %+v", post.Responses[0])
	}
}
