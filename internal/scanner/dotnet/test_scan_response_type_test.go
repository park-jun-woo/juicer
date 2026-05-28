//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_ResponseType -- ActionResult<T> 응답 타입 추출 테스트
package dotnet

import "testing"

func TestScan_ResponseType(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "Controllers/UsersController.cs", basicControllerSource)
	writeFile(t, dir, "Models/UserDto.cs", basicUserDtoSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	ep0 := result.Endpoints[0]
	if len(ep0.Responses) != 1 {
		t.Fatalf("ep0 expected 1 response, got %d", len(ep0.Responses))
	}
	if ep0.Responses[0].TypeName != "UserDto" {
		t.Errorf("ep0 response type: want UserDto, got %s", ep0.Responses[0].TypeName)
	}
	if ep0.Responses[0].Body != "array" {
		t.Errorf("ep0 response body: want array, got %s", ep0.Responses[0].Body)
	}

	ep1 := result.Endpoints[1]
	if len(ep1.Responses) != 1 {
		t.Fatalf("ep1 expected 1 response, got %d", len(ep1.Responses))
	}
	if ep1.Responses[0].TypeName != "UserDto" {
		t.Errorf("ep1 response type: want UserDto, got %s", ep1.Responses[0].TypeName)
	}
}
