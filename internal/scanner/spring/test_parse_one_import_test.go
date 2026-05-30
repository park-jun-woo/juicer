//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestParseOneImport 테스트
package spring

import "testing"

func TestParseOneImport(t *testing.T) {
	root, src := parseS(t, `import com.example.UserDto;`)
	imps := findAllByType(root, "import_declaration")
	name, fqcn := parseOneImport(imps[0], src)
	if name != "UserDto" || fqcn != "com.example.UserDto" {
		t.Fatalf("got %q %q", name, fqcn)
	}
}
