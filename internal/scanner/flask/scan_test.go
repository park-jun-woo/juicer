//ff:func feature=scan type=test control=sequence topic=flask
//ff:what Scan 빈/엣지 케이스 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NoPyFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("hi"), 0o644)
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}

func TestScan_FullProject(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(`from flask import Flask
app = Flask(__name__)

@app.route('/health')
def health():
    return 'ok'

@app.route('/users/<int:user_id>', methods=['GET', 'POST'])
def user(user_id):
    return {}
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	// /health (GET) + /users (GET, POST) = 3 endpoints
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}
}
