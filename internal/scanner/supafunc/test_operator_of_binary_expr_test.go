//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestOperatorOfBinaryExpr 테스트
package supafunc

import "testing"

func TestOperatorOfBinaryExpr(t *testing.T) {
	fi := mustParse(t, []byte(`const x = a === b;`))
	bins := findAllByType(fi.Root, "binary_expression")
	if len(bins) == 0 {
		t.Fatal("no binary expr")
	}
	if got := operatorOfBinaryExpr(bins[0], fi.Src); got != "===" {
		t.Fatalf("got %q", got)
	}
}
