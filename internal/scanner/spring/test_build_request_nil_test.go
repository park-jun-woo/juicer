//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildRequest_Nil 테스트
package spring

import "testing"

func TestBuildRequest_Nil(t *testing.T) {
	if r := buildRequest(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
