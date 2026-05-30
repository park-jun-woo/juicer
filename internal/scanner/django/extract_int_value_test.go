//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractIntValue — keyword 인자 정수 추출 분기를 검증
package django

import "testing"

func TestExtractIntValue(t *testing.T) {
	src := []byte("x = CharField(max_length=100)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	v := extractIntValue(kw[0], src)
	if v == nil || *v != 100 {
		t.Fatalf("expected 100, got %v", v)
	}
}

func TestExtractIntValue_Underscore(t *testing.T) {
	// Python integer literal with underscore separators exercises the
	// non-digit branch (underscores are skipped).
	src := []byte("x = CharField(max_length=1_000)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	v := extractIntValue(kw[0], src)
	if v == nil || *v != 1000 {
		t.Fatalf("expected 1000, got %v", v)
	}
}

func TestExtractIntValue_NoInteger(t *testing.T) {
	src := []byte("x = CharField(required=False)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if v := extractIntValue(kw[0], src); v != nil {
		t.Fatalf("expected nil for non-integer value, got %v", *v)
	}
}
