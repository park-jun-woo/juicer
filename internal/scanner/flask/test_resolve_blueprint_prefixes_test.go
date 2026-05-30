//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveBlueprintPrefixes 테스트
package flask

import "testing"

func TestResolveBlueprintPrefixes(t *testing.T) {
	def := flaskFile(t, `from flask import Blueprint
api = Blueprint("api", __name__, url_prefix="/api")
`)
	reg := flaskFile(t, `from flask import Flask
app = Flask(__name__)
app.register_blueprint(api, url_prefix="/override")
`)
	prefixes := resolveBlueprintPrefixes([]fileInfo{def, reg})

	if prefixes["api"] != "/override" {
		t.Fatalf("expected /override, got %v", prefixes)
	}
}
