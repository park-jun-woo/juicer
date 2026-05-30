//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestApplyRegisterBlueprintOverrides_Override 테스트
package flask

import "testing"

func TestApplyRegisterBlueprintOverrides_Override(t *testing.T) {
	src := `from flask import Flask
from .auth import auth as auth_bp

app = Flask(__name__)
app.register_blueprint(auth_bp, url_prefix='/v2/auth')
`
	fi := flaskFile(t, src)
	prefixes := blueprintPrefix{}
	applyRegisterBlueprintOverrides(fi, prefixes)

	if prefixes["auth"] != "/v2/auth" {
		t.Fatalf("expected override on canonical auth, got %v", prefixes)
	}
}
