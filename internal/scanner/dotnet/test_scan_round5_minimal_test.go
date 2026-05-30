//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestScan_Round5Minimal 테스트
package dotnet

import "testing"

func TestScan_Round5Minimal(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Program.cs", round5MinimalSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	get := result.Endpoints[0]
	if get.Path != "/items/{id}" {
		t.Errorf("path: want /items/{id}, got %s", get.Path)
	}
	if get.Request == nil || len(get.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param, got %+v", get.Request)
	}
	if get.Request.PathParams[0].Name != "id" {
		t.Errorf("pathparam: want id, got %s", get.Request.PathParams[0].Name)
	}
	if len(get.Request.Headers) != 1 || get.Request.Headers[0].Name != "token" {
		t.Errorf("header: want token, got %+v", get.Request.Headers)
	}

	for _, p := range get.Request.PathParams {
		if p.Name == "logger" {
			t.Error("logger should be filtered as DI type")
		}
	}
	post := result.Endpoints[1]
	if post.Request == nil || post.Request.Body == nil {
		t.Fatalf("POST expected body, got %+v", post.Request)
	}
	if post.Request.Body.TypeName != "CreateItemRequest" {
		t.Errorf("POST body type: want CreateItemRequest, got %s", post.Request.Body.TypeName)
	}
	if len(post.Request.Query) != 1 || post.Request.Query[0].Name != "draft" {
		t.Errorf("POST query: want draft, got %+v", post.Request.Query)
	}
	list := result.Endpoints[2]
	_ = list
}
