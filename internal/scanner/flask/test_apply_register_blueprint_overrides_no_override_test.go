//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestApplyRegisterBlueprintOverrides_NoOverride 테스트
package flask

import "testing"

func TestApplyRegisterBlueprintOverrides_NoOverride(t *testing.T) {

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
