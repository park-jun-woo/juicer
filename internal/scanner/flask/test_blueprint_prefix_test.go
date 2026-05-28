//ff:func feature=scan type=test control=sequence topic=flask
//ff:what Blueprint url_prefix가 라우트 경로에 전파된다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlueprintPrefix_Propagation(t *testing.T) {
	dir := t.TempDir()
	src := `from flask import Flask, Blueprint

app = Flask(__name__)
api = Blueprint('api', __name__, url_prefix='/api')

@api.route('/users')
def list_users():
    return []

@api.route('/users/<int:user_id>')
def get_user(user_id):
    return {}

app.register_blueprint(api)
`
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" || ep0.Path != "/api/users" {
		t.Errorf("endpoint 0: expected GET /api/users, got %s %s", ep0.Method, ep0.Path)
	}

	ep1 := result.Endpoints[1]
	if ep1.Method != "GET" || ep1.Path != "/api/users/{user_id}" {
		t.Errorf("endpoint 1: expected GET /api/users/{user_id}, got %s %s", ep1.Method, ep1.Path)
	}
	if ep1.Request == nil || len(ep1.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param, got %v", ep1.Request)
	}
	if ep1.Request.PathParams[0].Name != "user_id" || ep1.Request.PathParams[0].Type != "integer" {
		t.Errorf("expected user_id (integer), got %s (%s)", ep1.Request.PathParams[0].Name, ep1.Request.PathParams[0].Type)
	}
}
