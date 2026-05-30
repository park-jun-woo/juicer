//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAnnotatedDependsAlias_NotAnnotated 테스트
package fastapi

import "testing"

func TestParseAnnotatedDependsAlias_NotAnnotated(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x = arr[0]\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}
