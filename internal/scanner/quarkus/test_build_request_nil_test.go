//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildRequest_Nil 테스트
package quarkus

import "testing"

func TestBuildRequest_Nil(t *testing.T) {
	if r := buildRequest(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
