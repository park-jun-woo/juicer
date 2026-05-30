//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestResolveType_Round5 테스트
package echo

import "testing"

func TestResolveType_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	name, fields := resolveType(typ)
	if name != "UserDto" {
		t.Fatalf("name: %q", name)
	}
	if len(fields) == 0 {
		t.Fatalf("expected fields, got %+v", fields)
	}

	for _, f := range fields {
		if f.Name == "Hidden" {
			t.Errorf("Hidden should be excluded")
		}
	}

	lname, _ := resolveType(namedType(t, info, "L"))
	if lname != "[]UserDto" {
		t.Fatalf("slice name: %q", lname)
	}
}
