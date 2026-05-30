//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractKeywordArg — 이름으로 키워드 문자열 값 추출 분기를 검증
package django

import "testing"

func TestExtractKeywordArg_Found(t *testing.T) {
	src := []byte("x = path('a/', view, name='home')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	if got := extractKeywordArg(args, "name", src); got != "home" {
		t.Fatalf("got %q, want home", got)
	}
}

func TestExtractKeywordArg_WrongName(t *testing.T) {
	src := []byte("x = path('a/', view, name='home')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := extractKeywordArg(args, "missing", src); got != "" {
		t.Fatalf("expected empty for missing name, got %q", got)
	}
}

func TestExtractKeywordArg_NonStringValue(t *testing.T) {
	src := []byte("x = path('a/', view, count=5)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := extractKeywordArg(args, "count", src); got != "" {
		t.Fatalf("expected empty for non-string value, got %q", got)
	}
}
