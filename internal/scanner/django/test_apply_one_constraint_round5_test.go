//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplyOneConstraint_Round5 테스트
package django

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyOneConstraint_Round5(t *testing.T) {
	src := []byte("name = serializers.CharField(max_length=12)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	kw := djFirst(t, root, "keyword_argument")
	f := &scanner.Field{}
	applyOneConstraint(f, "max_length", kw, src)
	if f.MaxLength == nil || *f.MaxLength != 12 {
		t.Fatalf("max_length: %v", f.MaxLength)
	}
}
