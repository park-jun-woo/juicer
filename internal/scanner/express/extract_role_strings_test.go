//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractRoleStrings: 문자열인자 수집 / 비문자열 스킵 / args없음
package express

import "testing"

func TestExtractRoleStrings_Collects(t *testing.T) {
	fi := mustParse(t, []byte(`requireRole('admin', x, 'editor');`))
	got := extractRoleStrings(firstCallExpr(t, fi), fi.Src)
	if len(got) != 2 || got[0] != "admin" || got[1] != "editor" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractRoleStrings_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("requireRole`x`;"))
	if got := extractRoleStrings(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestExtractRoleStrings_NoStrings(t *testing.T) {
	fi := mustParse(t, []byte(`requireRole(a, b);`))
	if got := extractRoleStrings(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %v", got)
	}
}
