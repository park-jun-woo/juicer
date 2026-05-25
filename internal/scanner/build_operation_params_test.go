//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildOperationParams_Nil 테스트
package scanner

import "testing"

func TestBuildOperationParams_Nil(t *testing.T) {
	result := buildOperationParams(nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
