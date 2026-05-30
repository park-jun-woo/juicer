//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractInterfaceMethodEndpoints_Round5 테스트
package spring

import "testing"

func TestExtractInterfaceMethodEndpoints_Round5(t *testing.T) {
	fi := sFileInfo(t, `
@RequestMapping("/api")
interface UserApi {
    @GetMapping("/users")
    java.util.List<UserDto> list();
}
`)
	iface := sFirst(t, fi.root, "interface_declaration")
	eps := extractInterfaceMethodEndpoints(iface, fi)
	if len(eps) == 0 {
		t.Fatalf("expected interface endpoints, got %d", len(eps))
	}
}
