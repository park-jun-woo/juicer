//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindDecoratorNodes_AttributeNoCall 테스트
package fastapi

import "testing"

func TestFindDecoratorNodes_AttributeNoCall(t *testing.T) {

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
