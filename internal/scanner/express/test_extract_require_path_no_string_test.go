//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRequirePath_NoString 테스트
package express

import "testing"

func TestExtractRequirePath_NoString(t *testing.T) {
	fi := mustParse(t, []byte(`require(modulePath);`))
	if got := extractRequirePath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
