//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplyActionKeywords 테스트
package django

import "testing"

func TestApplyActionKeywords(t *testing.T) {
	src := []byte(`x = action(detail=True, methods=['get', 'post'], url_path='custom')` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	ai := &actionInfo{}
	applyActionKeywords(ai, args, src)

	if !ai.detail {
		t.Error("expected detail true")
	}
	if len(ai.methods) != 2 {
		t.Errorf("methods = %v, want 2", ai.methods)
	}
	if ai.urlPath != "custom" {
		t.Errorf("urlPath = %q, want custom", ai.urlPath)
	}
}
