//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParsePython_Basic 테스트
package fastapi

import "testing"

func TestParsePython_Basic(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI
app = FastAPI()

@app.get("/")
async def root():
    return {"message": "hello"}
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if root == nil {
		t.Fatal("expected non-nil root")
	}
	if root.Type() != "module" {
		t.Fatalf("expected module, got %s", root.Type())
	}
}
