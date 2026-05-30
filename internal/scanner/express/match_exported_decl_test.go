//ff:func feature=scan type=test control=sequence topic=express
//ff:what matchExportedDecl: export함수 / export변수 / 비export / 미매칭
package express

import "testing"

func TestMatchExportedDecl_Function(t *testing.T) {
	fi := mustParse(t, []byte(`export function h(req, res) { res.json({}); }`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}

func TestMatchExportedDecl_Variable(t *testing.T) {
	fi := mustParse(t, []byte(`export const h = (req, res) => { res.json({}); };`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}

func TestMatchExportedDecl_NotExport(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchExportedDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

func TestMatchExportedDecl_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`export function h() {}`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
