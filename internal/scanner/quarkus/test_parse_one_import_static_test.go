//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestParseOneImport_Static 테스트
package quarkus

import "testing"

func TestParseOneImport_Static(t *testing.T) {
	root, src := parseQ(t, `import static com.example.Util.foo;`)
	imps := findAllByType(root, "import_declaration")
	name, _ := parseOneImport(imps[0], src)
	if name != "" {
		t.Fatalf("static import should yield empty, got %q", name)
	}
}
