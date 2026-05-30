//ff:func feature=scan type=test control=sequence topic=django
//ff:what isSerializerSubclass — Serializer 부모 판별 분기를 검증
package django

import "testing"

func TestIsSerializerSubclass(t *testing.T) {
	if !isSerializerSubclass([]string{"ModelSerializer"}) {
		t.Error("expected true for ModelSerializer")
	}
	if isSerializerSubclass([]string{"object"}) {
		t.Error("expected false for non-serializer parent")
	}
	if isSerializerSubclass(nil) {
		t.Error("expected false for no parents")
	}
}
