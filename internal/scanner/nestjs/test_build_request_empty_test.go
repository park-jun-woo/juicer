//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildRequest_Empty 테스트
package nestjs

import "testing"

func TestBuildRequest_Empty(t *testing.T) {
	ep := endpointInfo{}
	req := buildRequest(ep)
	if req != nil {
		t.Fatal("expected nil for empty request")
	}
}
