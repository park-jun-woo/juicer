//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestAngleBracketDelta 테스트
package quarkus

import "testing"

func TestAngleBracketDelta(t *testing.T) {
	if angleBracketDelta('<') != 1 || angleBracketDelta('>') != -1 || angleBracketDelta('x') != 0 {
		t.Fatal("delta")
	}
}
