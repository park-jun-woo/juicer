//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestAngleBracketDelta 테스트
package spring

import "testing"

func TestAngleBracketDelta(t *testing.T) {
	if angleBracketDelta('<') != 1 || angleBracketDelta('>') != -1 || angleBracketDelta('z') != 0 {
		t.Fatal("delta")
	}
}
