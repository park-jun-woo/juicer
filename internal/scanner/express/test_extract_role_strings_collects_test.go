//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoleStrings_Collects 테스트
package express

import "testing"

func TestExtractRoleStrings_Collects(t *testing.T) {
	fi := mustParse(t, []byte(`requireRole('admin', x, 'editor');`))
	got := extractRoleStrings(firstCallExpr(t, fi), fi.Src)
	if len(got) != 2 || got[0] != "admin" || got[1] != "editor" {
		t.Fatalf("got %v", got)
	}
}
