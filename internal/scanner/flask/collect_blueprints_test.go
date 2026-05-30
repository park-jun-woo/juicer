//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectBlueprints 테스트
package flask

import "testing"

func TestCollectBlueprints(t *testing.T) {
	src := []byte(`from flask import Blueprint

users_bp = Blueprint("users", __name__, url_prefix="/api/users")
plain = Blueprint("plain", __name__)
notbp = something()
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	bps := collectBlueprints(root, src)
	byVar := map[string]blueprintInfo{}
	for _, bp := range bps {
		byVar[bp.varName] = bp
	}
	if bp, ok := byVar["users_bp"]; !ok || bp.urlPrefix != "/api/users" {
		t.Fatalf("users_bp = %+v", byVar["users_bp"])
	}
	if _, ok := byVar["plain"]; !ok {
		t.Fatalf("plain blueprint not collected: %v", byVar)
	}
	if _, ok := byVar["notbp"]; ok {
		t.Fatalf("non-Blueprint assignment should not be collected")
	}
}

func TestCollectBlueprints_None(t *testing.T) {
	root, _ := parsePython([]byte("x = 1\n"))
	if bps := collectBlueprints(root, []byte("x = 1\n")); len(bps) != 0 {
		t.Fatalf("expected none, got %d", len(bps))
	}
}
