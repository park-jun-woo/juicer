//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAnnotatedDependsAlias_NoIdentifierLeft 테스트
package fastapi

import "testing"

func TestParseAnnotatedDependsAlias_NoIdentifierLeft(t *testing.T) {

	assign, src := firstAssignment(t, []byte("obj.attr = 5\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}
