//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAnnotatedDependsAlias_NoSubscript 테스트
package fastapi

import "testing"

func TestParseAnnotatedDependsAlias_NoSubscript(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x = 5\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}
