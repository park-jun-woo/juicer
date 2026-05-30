//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildResponse_Nil 테스트
package quarkus

import "testing"

func TestBuildResponse_Nil(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
