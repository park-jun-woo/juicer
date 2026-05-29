//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectImportAliases가 aliased_import 매핑을 수집한다
package flask

import "testing"

func TestCollectImportAliases(t *testing.T) {
	src := []byte("from .auth import auth as auth_blueprint\nfrom .api import api as api_blueprint, foo\nimport os\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	aliases := collectImportAliases(root, src)
	if aliases["auth_blueprint"] != "auth" {
		t.Errorf("expected auth, got %q", aliases["auth_blueprint"])
	}
	if aliases["api_blueprint"] != "api" {
		t.Errorf("expected api, got %q", aliases["api_blueprint"])
	}
	if _, ok := aliases["foo"]; ok {
		t.Errorf("non-aliased import foo must not be mapped")
	}
}
