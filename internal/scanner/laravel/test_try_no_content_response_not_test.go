//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryNoContentResponse_Not 테스트
package laravel

import "testing"

func TestTryNoContentResponse_Not(t *testing.T) {
	if resp := tryNoContentResponse("return $x;"); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}
