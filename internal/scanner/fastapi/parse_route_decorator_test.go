//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseRouteDecorator 테스트
package fastapi

import "testing"

func TestParseRouteDecorator(t *testing.T) {
	src := []byte("@app.post('/items', status_code=201)\ndef create(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	method, path, routerVar, status, _, _ := parseRouteDecorator(decs[0], src)
	if method != "POST" {
		t.Fatalf("method: got %q", method)
	}
	if path != "/items" {
		t.Fatalf("path: got %q", path)
	}
	if routerVar != "app" {
		t.Fatalf("routerVar: got %q", routerVar)
	}
	if status != 201 {
		t.Fatalf("status: got %d", status)
	}
}

func TestParseRouteDecorator_NotHTTP(t *testing.T) {
	src := []byte("@app.middleware('http')\ndef h(): pass\n")
	root, _ := parsePython(src)
	decs := findAllByType(root, "decorator")
	method, _, _, _, _, _ := parseRouteDecorator(decs[0], src)
	if method != "" {
		t.Fatalf("expected empty method, got %q", method)
	}
}

func TestParseRouteDecorator_NoAttribute(t *testing.T) {
	// bare-name decorator -> no attribute node
	src := []byte("@staticmethod\ndef h(): pass\n")
	root, _ := parsePython(src)
	decs := findAllByType(root, "decorator")
	method, _, _, _, _, _ := parseRouteDecorator(decs[0], src)
	if method != "" {
		t.Fatalf("expected empty, got %q", method)
	}
}
