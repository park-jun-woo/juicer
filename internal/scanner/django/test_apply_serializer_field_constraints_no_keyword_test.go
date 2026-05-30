//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplySerializerFieldConstraints_NoKeyword 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplySerializerFieldConstraints_NoKeyword(t *testing.T) {
	src := []byte(`x = CharField('a', 'b')` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	f := &scanner.Field{Name: "title", Type: "string"}
	applySerializerFieldConstraints(f, args, src, "CharField")
	if f.MaxLength != nil {
		t.Errorf("expected no constraints, got MaxLength=%v", f.MaxLength)
	}
}
