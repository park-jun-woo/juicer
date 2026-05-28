//ff:func feature=scan type=test control=sequence topic=flask
//ff:what register_blueprint의 url_prefix 오버라이드를 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlueprintPrefix_Override(t *testing.T) {
	dir := t.TempDir()
	src := `from flask import Flask, Blueprint

app = Flask(__name__)
users_bp = Blueprint('users', __name__, url_prefix='/api/users')

@users_bp.route('/<int:id>')
def get_user(id):
    return {}

app.register_blueprint(users_bp, url_prefix='/v2/users')
`
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Path != "/v2/users/{id}" {
		t.Errorf("expected /v2/users/{id}, got %s", ep.Path)
	}
}
