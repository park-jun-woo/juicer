//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectDTORequests_QueryDTO 테스트
package nestjs

import "testing"

func TestCollectDTORequests_QueryDTO(t *testing.T) {
	ep := endpointInfo{queryDTOType: "ListUserReqDto"}
	imports := map[string]string{"ListUserReqDto": "@/common/dto/list-user-req.dto"}
	reqs := collectDTORequests(ep, imports, "/src/controller.ts", "/project", 0)
	if len(reqs) != 1 {
		t.Fatalf("expected 1, got %d", len(reqs))
	}
	if !reqs[0].isQuery {
		t.Fatal("expected isQuery=true")
	}
	if reqs[0].isBody {
		t.Fatal("expected isBody=false")
	}
	if reqs[0].typeName != "ListUserReqDto" {
		t.Fatalf("expected ListUserReqDto, got %q", reqs[0].typeName)
	}
}
