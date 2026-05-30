//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRequirePath_Valid 테스트
package express

import "testing"

func TestExtractRequirePath_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`const x = require('./routes');`))
	if got := extractRequirePath(firstCallExpr(t, fi), fi.Src); got != "./routes" {
		t.Fatalf("got %q", got)
	}
}
