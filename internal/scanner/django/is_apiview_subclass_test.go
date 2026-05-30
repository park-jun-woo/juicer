//ff:func feature=scan type=test control=sequence topic=django
//ff:what isAPIViewSubclass — APIView 부모 판별 분기를 검증
package django

import "testing"

func TestIsAPIViewSubclass(t *testing.T) {
	if !isAPIViewSubclass([]string{"object", "APIView"}) {
		t.Error("expected true when APIView is a parent")
	}
	if isAPIViewSubclass([]string{"object", "Mixin"}) {
		t.Error("expected false without an APIView base")
	}
	if isAPIViewSubclass(nil) {
		t.Error("expected false for no parents")
	}
}
