//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryInheritFileClasses 테스트
package fastapi

import "testing"

func TestTryInheritFileClasses(t *testing.T) {
	src := []byte("class Base(BaseModel):\n    id: int\n\nclass Child(Base):\n    name: str\n")
	root, _ := parsePython(src)
	fi := &fileInfo{src: src, root: root, models: map[string][]pydanticField{}}
	globalModels := map[string]*fileInfo{"Base": fi}

	if !tryInheritFileClasses(fi, globalModels) {
		t.Fatal("expected Child to be added")
	}
	// second run: nothing new -> false
	if tryInheritFileClasses(fi, globalModels) {
		t.Fatal("expected no new additions on second run")
	}

	// nil root -> false
	nilFi := &fileInfo{root: nil}
	if tryInheritFileClasses(nilFi, globalModels) {
		t.Fatal("nil root should return false")
	}
}
