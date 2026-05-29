//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what 어트리뷰트 없는 복합타입 파라미터가 암묵 body로 바인딩되는지 검증한다
package dotnet

import "testing"

func TestScan_ImplicitBody(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/UsersController.cs", implicitBodyControllerSource)
	writeFile(t, dir, "Models/UserCreateViewModel.cs", implicitBodyDtoSource)
	writeFile(t, dir, "Models/UserDto.cs", basicUserDtoSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "POST" {
		t.Errorf("ep0 method: want POST, got %s", ep0.Method)
	}
	if ep0.Request == nil || ep0.Request.Body == nil {
		t.Fatal("ep0 expected implicit body")
	}
	if ep0.Request.Body.TypeName != "UserCreateViewModel" {
		t.Errorf("ep0 body type: want UserCreateViewModel, got %s", ep0.Request.Body.TypeName)
	}
	if len(ep0.Request.Body.Fields) != 2 {
		t.Fatalf("ep0 body fields: want 2, got %d", len(ep0.Request.Body.Fields))
	}

	ep1 := result.Endpoints[1]
	if ep1.Request == nil || ep1.Request.Body == nil {
		t.Fatal("ep1 expected implicit body alongside path param")
	}
	if ep1.Request.Body.TypeName != "UserCreateViewModel" {
		t.Errorf("ep1 body type: want UserCreateViewModel, got %s", ep1.Request.Body.TypeName)
	}
	if len(ep1.Request.PathParams) != 1 || ep1.Request.PathParams[0].Name != "id" {
		t.Fatalf("ep1 expected path param id")
	}
}
