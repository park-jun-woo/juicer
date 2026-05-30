//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractRequirePath: 정상 / args없음 / string인자없음 분기
package express

import "testing"

func TestExtractRequirePath_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`const x = require('./routes');`))
	if got := extractRequirePath(firstCallExpr(t, fi), fi.Src); got != "./routes" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractRequirePath_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("require`x`;"))
	if got := extractRequirePath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractRequirePath_NoString(t *testing.T) {
	fi := mustParse(t, []byte(`require(modulePath);`))
	if got := extractRequirePath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
