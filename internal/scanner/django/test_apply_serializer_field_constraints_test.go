//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplySerializerFieldConstraints 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplySerializerFieldConstraints(t *testing.T) {
	src := []byte(`x = CharField(max_length=100, required=False)` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	f := &scanner.Field{Name: "title", Type: "string"}
	applySerializerFieldConstraints(f, args, src, "CharField")

	if f.MaxLength == nil || *f.MaxLength != 100 {
		t.Errorf("expected MaxLength 100, got %v", f.MaxLength)
	}
}
