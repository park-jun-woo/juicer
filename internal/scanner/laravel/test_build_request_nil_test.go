//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildRequest_Nil 테스트
package laravel

import "testing"

func TestBuildRequest_Nil(t *testing.T) {

	if r := buildRequest(t.TempDir(), nil, nil, map[string]*fileInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
