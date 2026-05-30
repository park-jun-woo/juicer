//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findDecoratorNodes 테스트
package fastapi

import "testing"

func TestFindDecoratorNodes(t *testing.T) {
	src := []byte("@router.get('/x')\ndef f(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorators")
	}
	callNode, attrNode := findDecoratorNodes(decs[0])
	if callNode == nil {
		t.Fatal("expected call node")
	}
	// attrNode might be nil if it's inside the call
	_ = attrNode
}

func TestFindDecoratorNodes_AttributeNoCall(t *testing.T) {
	// decorator that is a bare attribute, no call -> attrNode set, callNode nil
	src := []byte("@app.middleware\ndef f(): pass\n")
	root, _ := parsePython(src)
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorators")
	}
	callNode, attrNode := findDecoratorNodes(decs[0])
	if callNode != nil {
		t.Fatalf("expected nil call node, got %s", callNode.Type())
	}
	if attrNode == nil {
		t.Fatal("expected attribute node")
	}
}
