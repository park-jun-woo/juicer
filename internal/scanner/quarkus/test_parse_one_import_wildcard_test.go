//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestParseOneImport_Wildcard 테스트
package quarkus

import "testing"

func TestParseOneImport_Wildcard(t *testing.T) {
	root, src := parseQ(t, `import com.example.*;`)
	imps := findAllByType(root, "import_declaration")
	name, _ := parseOneImport(imps[0], src)
	if name != "" {
		t.Fatalf("wildcard should yield empty, got %q", name)
	}
}
