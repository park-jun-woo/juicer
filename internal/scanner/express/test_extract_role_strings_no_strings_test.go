//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoleStrings_NoStrings 테스트
package express

import "testing"

func TestExtractRoleStrings_NoStrings(t *testing.T) {
	fi := mustParse(t, []byte(`requireRole(a, b);`))
	if got := extractRoleStrings(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %v", got)
	}
}
