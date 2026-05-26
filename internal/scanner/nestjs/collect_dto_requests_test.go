//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectDTORequests_BodyAndReturn 테스트
package nestjs

import "testing"

func TestCollectDTORequests_BodyAndReturn(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	imports := map[string]string{"CreateUserDto": "./dto/create-user.dto"}
	reqs := collectDTORequests(ep, imports, "/src/controller.ts", 0)
	if len(reqs) != 2 {
		t.Fatalf("expected 2, got %d", len(reqs))
	}
	if !reqs[0].isBody {
		t.Fatal("first req should be body")
	}
	if reqs[1].isBody {
		t.Fatal("second req should not be body")
	}
}
