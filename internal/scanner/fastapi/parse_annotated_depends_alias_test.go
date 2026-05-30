//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseAnnotatedDependsAlias: Annotated Depends 별칭 / 비subscript / 비Annotated
package fastapi

import "testing"

func TestParseAnnotatedDependsAlias_Match(t *testing.T) {
	assign, src := firstAssignment(t, []byte("SessionDep = Annotated[Session, Depends(get_db)]\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "SessionDep" || fn != "get_db" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}

func TestParseAnnotatedDependsAlias_NoIdentifierLeft(t *testing.T) {
	// left side is a subscript/attribute, no plain identifier child
	assign, src := firstAssignment(t, []byte("obj.attr = 5\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}

func TestParseAnnotatedDependsAlias_NoSubscript(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x = 5\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}

func TestParseAnnotatedDependsAlias_NotAnnotated(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x = arr[0]\n"))
	alias, fn := parseAnnotatedDependsAlias(assign, src)
	if alias != "" || fn != "" {
		t.Fatalf("alias=%q fn=%q", alias, fn)
	}
}
