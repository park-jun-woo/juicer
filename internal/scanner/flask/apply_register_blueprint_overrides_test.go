//ff:func feature=scan type=test control=sequence topic=flask
//ff:what applyRegisterBlueprintOverrides 테스트
package flask

import "testing"

func flaskFile(t *testing.T, src string) fileInfo {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: b, root: root}
}

func TestApplyRegisterBlueprintOverrides_Override(t *testing.T) {
	src := `from flask import Flask
from .auth import auth as auth_bp

app = Flask(__name__)
app.register_blueprint(auth_bp, url_prefix='/v2/auth')
`
	fi := flaskFile(t, src)
	prefixes := blueprintPrefix{}
	applyRegisterBlueprintOverrides(fi, prefixes)
	// alias auth_bp resolves to canonical "auth"
	if prefixes["auth"] != "/v2/auth" {
		t.Fatalf("expected override on canonical auth, got %v", prefixes)
	}
}

func TestApplyRegisterBlueprintOverrides_NoOverride(t *testing.T) {
	// register_blueprint without url_prefix -> no entry added
	src := `from flask import Flask
api = something()
app = Flask(__name__)
app.register_blueprint(api)
`
	fi := flaskFile(t, src)
	prefixes := blueprintPrefix{}
	applyRegisterBlueprintOverrides(fi, prefixes)
	if len(prefixes) != 0 {
		t.Fatalf("expected no overrides, got %v", prefixes)
	}
}

func TestApplyRegisterBlueprintOverrides_NoRegisterCall(t *testing.T) {
	src := `x = 1
y = foo(2)
`
	fi := flaskFile(t, src)
	prefixes := blueprintPrefix{}
	applyRegisterBlueprintOverrides(fi, prefixes)
	if len(prefixes) != 0 {
		t.Fatalf("expected no overrides, got %v", prefixes)
	}
}
