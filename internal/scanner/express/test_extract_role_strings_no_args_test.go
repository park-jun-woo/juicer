//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoleStrings_NoArgs 테스트
package express

import "testing"

func TestExtractRoleStrings_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("requireRole`x`;"))
	if got := extractRoleStrings(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %v", got)
	}
}
