//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildFormRequestBody_NoFields 테스트
package laravel

import "testing"

func TestBuildFormRequestBody_NoFields(t *testing.T) {

	cm := &controllerMethod{formRequestRef: "NonexistentRequest"}
	if b := buildFormRequestBody(t.TempDir(), cm, map[string]*fileInfo{}); b != nil {
		t.Fatalf("expected nil, got %+v", b)
	}
}
