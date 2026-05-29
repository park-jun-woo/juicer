//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 파라미터 없는 URL 경로 추출 테스트
package laravel

import "testing"

func TestExtractURLParams_None(t *testing.T) {
	params := extractURLParams("/users")
	if len(params) != 0 {
		t.Errorf("expected 0 params, got %d", len(params))
	}
}
