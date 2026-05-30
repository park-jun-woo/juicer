//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveInterfaceEndpoints 테스트
package spring

import (
	"path/filepath"
	"testing"
)

func TestResolveInterfaceEndpoints(t *testing.T) {
	dir := t.TempDir()
	ifaceSrc := `
@RequestMapping("/api")
interface UserApi {
    @GetMapping("/users")
    String list();
}
`
	writeFile(t, dir, "UserApi.java", ifaceSrc)
	prefix, eps := resolveInterfaceEndpoints(filepath.Join(dir, "UserApi.java"), "UserApi", dir)
	if prefix != "/api" {
		t.Fatalf("prefix: %q", prefix)
	}
	if len(eps) != 1 {
		t.Fatalf("endpoints: %+v", eps)
	}
}
