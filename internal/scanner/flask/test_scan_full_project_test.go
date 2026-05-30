//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestScan_FullProject 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

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

	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}
}
