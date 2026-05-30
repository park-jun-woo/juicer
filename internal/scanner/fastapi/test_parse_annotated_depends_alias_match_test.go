//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseAnnotatedDependsAlias_Match 테스트
package fastapi

import "testing"

func TestParseAnnotatedDependsAlias_Match(t *testing.T) {
	assign, src := firstAssignment(t, []byte("SessionDep = Annotated[Session, Depends(get_db)]\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "SessionDep" || fn != "get_db" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}
