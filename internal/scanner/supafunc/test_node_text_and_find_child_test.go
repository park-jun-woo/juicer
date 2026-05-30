//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestNodeTextAndFindChild 테스트
package supafunc

import "testing"

func TestNodeTextAndFindChild(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("arguments")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("nil expected")
	}
	id := findChildByType(calls[0], "identifier")
	if id != nil && nodeText(id, fi.Src) != "foo" {
		t.Fatalf("name %q", nodeText(id, fi.Src))
	}
}
