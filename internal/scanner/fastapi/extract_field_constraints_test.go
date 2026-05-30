//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractFieldConstraints: keyword 인자 추출 / 비keyword(positional) 스킵
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstArgList(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	al := findAllByType(root, "argument_list")
	if len(al) == 0 {
		t.Fatal("no argument_list")
	}
	return al[0], src
}

func TestExtractFieldConstraints(t *testing.T) {
	// includes a positional arg (...) which is not keyword_argument -> skipped
	args, src := firstArgList(t, []byte("x = Field(..., ge=1, le=9, min_length=2, max_length=8)\n"))
	f := &pydanticField{}
	extractFieldConstraints(args, src, f)
	if f.ge == nil || *f.ge != 1 || f.le == nil || *f.le != 9 {
		t.Fatalf("ge/le wrong: %+v", f)
	}
	if f.minLength == nil || *f.minLength != 2 || f.maxLength == nil || *f.maxLength != 8 {
		t.Fatalf("length wrong: %+v", f)
	}
	// the keyNode==nil branch is unreachable: a keyword_argument always has an
	// identifier name child.
}
