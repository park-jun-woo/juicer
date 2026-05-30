//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryNoContentResponse 테스트
package laravel

import "testing"

func TestTryNoContentResponse(t *testing.T) {
	resp := tryNoContentResponse("return response()->noContent();")
	if resp == nil || resp.Status != "204" || resp.Kind != "empty" {
		t.Fatalf("got %+v", resp)
	}
}
