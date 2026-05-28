//ff:func feature=scan type=test control=sequence topic=flask
//ff:what Flask E2E 스캔 테스트 — 라우트, Blueprint, URL 파라미터를 통합 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()
	src := `from flask import Flask, Blueprint

app = Flask(__name__)
api = Blueprint('api', __name__, url_prefix='/api')

@app.route('/health')
def health():
    return {'status': 'ok'}

@api.route('/users', methods=['GET'])
def list_users():
    return []

@api.route('/users', methods=['POST'])
def create_user():
    return {}, 201

@api.route('/users/<int:user_id>', methods=['GET'])
def get_user(user_id):
    return {}

@api.get('/users/<int:user_id>/posts')
def list_user_posts(user_id):
    return []

app.register_blueprint(api)
`
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(src), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Endpoints) != 5 {
		for i, ep := range result.Endpoints {
			t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
		}
		t.Fatalf("expected 5 endpoints, got %d", len(result.Endpoints))
	}

	// /health — GET
	ep := result.Endpoints[0]
	if ep.Method != "GET" || ep.Path != "/health" || ep.Handler != "health" {
		t.Errorf("endpoint 0: expected GET /health health, got %s %s %s", ep.Method, ep.Path, ep.Handler)
	}

	// /api/users — GET
	ep = result.Endpoints[1]
	if ep.Method != "GET" || ep.Path != "/api/users" || ep.Handler != "list_users" {
		t.Errorf("endpoint 1: expected GET /api/users list_users, got %s %s %s", ep.Method, ep.Path, ep.Handler)
	}

	// /api/users — POST
	ep = result.Endpoints[2]
	if ep.Method != "POST" || ep.Path != "/api/users" || ep.Handler != "create_user" {
		t.Errorf("endpoint 2: expected POST /api/users create_user, got %s %s %s", ep.Method, ep.Path, ep.Handler)
	}

	// /api/users/{user_id} — GET
	ep = result.Endpoints[3]
	if ep.Method != "GET" || ep.Path != "/api/users/{user_id}" {
		t.Errorf("endpoint 3: expected GET /api/users/{user_id}, got %s %s", ep.Method, ep.Path)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("endpoint 3: expected 1 path param")
	}
	if ep.Request.PathParams[0].Name != "user_id" || ep.Request.PathParams[0].Type != "integer" {
		t.Errorf("endpoint 3: expected user_id (integer), got %s (%s)",
			ep.Request.PathParams[0].Name, ep.Request.PathParams[0].Type)
	}

	// /api/users/{user_id}/posts — GET
	ep = result.Endpoints[4]
	if ep.Method != "GET" || ep.Path != "/api/users/{user_id}/posts" {
		t.Errorf("endpoint 4: expected GET /api/users/{user_id}/posts, got %s %s", ep.Method, ep.Path)
	}
}
