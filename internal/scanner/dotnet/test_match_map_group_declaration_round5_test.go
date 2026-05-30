//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMatchMapGroupDeclaration_Round5 테스트
package dotnet

import "testing"

func TestMatchMapGroupDeclaration_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var users = api.MapGroup("/users");
`)
	groups := map[string]string{"api": "/api"}
	stmt := findAllByType(fi.root, "local_declaration_statement")[0]
	matchMapGroupDeclaration(stmt, fi, groups)
	if groups["users"] != "/api/users" {
		t.Fatalf("expected nested prefix, got %v", groups)
	}
}
