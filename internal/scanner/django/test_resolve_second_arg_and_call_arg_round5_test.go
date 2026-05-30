//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestResolveSecondArg_And_CallArg_Round5 테스트
package django

import "testing"

func TestResolveSecondArg_And_CallArg_Round5(t *testing.T) {

	src := []byte("path('api/', include('app.urls'))\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	if len(pos) < 2 {
		t.Fatalf("expected 2 args, got %d", len(pos))
	}
	var e urlEntry
	resolveSecondArg(&e, pos[1], src)
	if !e.isInclude || e.includeModule != "app.urls" {
		t.Fatalf("include: %+v", e)
	}

	src2 := []byte("path('x/', MyView.as_view())\n")
	root2, _ := parsePython(src2)
	args2 := djFirst(t, root2, "argument_list")
	pos2 := positionalArgs(args2)
	var e2 urlEntry
	resolveSecondArg(&e2, pos2[1], src2)
	if e2.viewName != "MyView" {
		t.Fatalf("viewName: %q", e2.viewName)
	}
}
