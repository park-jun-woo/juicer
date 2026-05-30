//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryJSONResponse_NotJSON 테스트
package laravel

import "testing"

func TestTryJSONResponse_NotJSON(t *testing.T) {
	if resp := tryJSONResponse(nil, nil, "return $x;"); resp != nil {
		t.Fatalf("expected nil, got %+v", resp)
	}
}
