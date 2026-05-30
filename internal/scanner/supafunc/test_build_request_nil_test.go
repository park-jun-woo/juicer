//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestBuildRequest_Nil 테스트
package supafunc

import "testing"

func TestBuildRequest_Nil(t *testing.T) {
	if r := buildRequest(nil, nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
