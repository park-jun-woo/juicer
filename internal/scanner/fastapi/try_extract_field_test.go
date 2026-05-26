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
