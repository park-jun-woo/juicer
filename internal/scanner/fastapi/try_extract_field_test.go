//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryExtractField 테스트
package fastapi

import "testing"

func TestTryExtractField(t *testing.T) {
	src := []byte("class M(BaseModel):\n    name: str\n    age: int = 25\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	block := findAllByType(root, "block")
	if len(block) == 0 {
		t.Fatal("no block")
	}
	count := 0
	for i := 0; i < int(block[0].ChildCount()); i++ {
		child := block[0].Child(i)
		f := tryExtractField(child, src)
		if f != nil {
			count++
		}
	}
	if count != 2 {
		t.Fatalf("expected 2 fields, got %d", count)
	}
}

func TestTryExtractField_DefaultNil(t *testing.T) {
	// a non field/assignment statement (pass) yields nil
	src := []byte("class M(BaseModel):\n    pass\n")
	root, _ := parsePython(src)
	for _, ps := range findAllByType(root, "pass_statement") {
		if f := tryExtractField(ps, src); f != nil {
			t.Fatalf("pass should yield nil, got %v", f)
		}
		return
	}
	t.Skip("no pass_statement")
}

func TestTryExtractField_Assignment(t *testing.T) {
	// raw assignment node
	src := []byte("class M(BaseModel):\n    age: int = 25\n")
	root, _ := parsePython(src)
	as := findAllByType(root, "assignment")
	if len(as) == 0 {
		t.Skip("no assignment node")
	}
	f := tryExtractField(as[0], src)
	if f == nil || f.name != "age" {
		t.Fatalf("assignment field: got %v", f)
	}
}
