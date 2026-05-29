//ff:func feature=scan type=test control=sequence topic=flask
//ff:what import alias로 등록된 Blueprint의 url_prefix가 라우트에 합성된다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlueprintAliasPrefix_CrossFile(t *testing.T) {
	dir := t.TempDir()
	authSrc := `from flask import Blueprint

auth = Blueprint('auth', __name__)

@auth.route('/login')
def login():
    return {}
`
	initSrc := `from flask import Flask
from .auth import auth as auth_blueprint

app = Flask(__name__)
app.register_blueprint(auth_blueprint, url_prefix='/auth')
`
	os.WriteFile(filepath.Join(dir, "auth.py"), []byte(authSrc), 0o644)
	os.WriteFile(filepath.Join(dir, "__init__.py"), []byte(initSrc), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	if result.Endpoints[0].Path != "/auth/login" {
		t.Errorf("expected /auth/login, got %s", result.Endpoints[0].Path)
	}
}
