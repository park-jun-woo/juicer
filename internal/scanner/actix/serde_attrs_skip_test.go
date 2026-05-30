//ff:func feature=scan type=test control=sequence topic=actix
//ff:what serdeAttrsSkip — skip 어트리뷰트 존재 여부를 검증
package actix

import "testing"

func TestSerdeAttrsSkip(t *testing.T) {
	if !serdeAttrsSkip([]serdeAttr{{rename: "x"}, {skip: true}}) {
		t.Error("expected true when a skip attr is present")
	}
	if serdeAttrsSkip([]serdeAttr{{rename: "x"}, {hasDefault: true}}) {
		t.Error("expected false when no skip attr")
	}
	if serdeAttrsSkip(nil) {
		t.Error("expected false for empty attrs")
	}
}
