//ff:func feature=scan type=test control=sequence topic=flask
//ff:what resolveBlueprintPrefixes 테스트
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
	// override from register_blueprint takes precedence
	if prefixes["api"] != "/override" {
		t.Fatalf("expected /override, got %v", prefixes)
	}
}

func TestResolveBlueprintPrefixes_NoOverride(t *testing.T) {
	def := flaskFile(t, `from flask import Blueprint
api = Blueprint("api", __name__, url_prefix="/api")
`)
	prefixes := resolveBlueprintPrefixes([]fileInfo{def})
	if prefixes["api"] != "/api" {
		t.Fatalf("expected /api, got %v", prefixes)
	}
}
