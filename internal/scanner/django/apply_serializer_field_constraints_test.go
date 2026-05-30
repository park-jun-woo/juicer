//ff:func feature=scan type=test control=sequence topic=django
//ff:what applySerializerFieldConstraints — serializer 필드 제약 추출 분기를 검증
package django

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
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
	// max_length constraint should be applied to the field.
	if f.MaxLength == nil || *f.MaxLength != 100 {
		t.Errorf("expected MaxLength 100, got %v", f.MaxLength)
	}
}

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
