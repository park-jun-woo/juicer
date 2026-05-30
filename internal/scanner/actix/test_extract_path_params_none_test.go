//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractPathParams_None 테스트
package actix

import "testing"

func TestExtractPathParams_None(t *testing.T) {

	if p := extractPathParams("/users/list"); len(p) != 0 {
		t.Fatalf("expected no params, got %+v", p)
	}
}
